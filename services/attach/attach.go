package attach

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"OverflowBackend/proto/repository_proto"
	"bytes"
	"context"
	"errors"
	"github.com/mailru/easyjson"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
)

type AttachService struct {
	config *config.Config
	db     repository_proto.DatabaseRepositoryClient
	s3     *minio.Client
	attach_proto.UnimplementedAttachServer
}

const (
	BASE_ATTACH_URL    = "/minio-storage/attach/"
	ATTACH_PREFIX_SIZE = 6
)

var ErrAccess = errors.New("Пользователь не имеет доступа к данному письму.")
var ErrJson = errors.New("Ошибка упаковки/распаковки JSON.")
var ErrAccessAttach = errors.New("Пользователь не имеет доступа к данному вложению.")

func (s *AttachService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, s3 *minio.Client) {
	s.config = config
	s.db = db
	s.s3 = s3
}

func (s *AttachService) SaveAttach(ctx context.Context, request *attach_proto.SaveAttachRequest) (*attach_proto.Nothing, error) {
	respMail, err := s.db.GetMailInfoById(context.Background(), &repository_proto.GetMailInfoByIdRequest{
		MailId: request.MailID,
	})
	if err != nil {
		return &attach_proto.Nothing{Status: false}, err
	}

	var mail models.Mail
	err = easyjson.Unmarshal(respMail.Mail, &mail)
	if mail.Sender != request.Username {
		log.Warning(ErrAccess)
		return &attach_proto.Nothing{Status: false}, ErrAccess
	}

	var file models.Attach
	if err := easyjson.Unmarshal(request.File, &file); err != nil {
		return &attach_proto.Nothing{Status: false}, err
	}

	fileName := pkg.RandSID(ATTACH_PREFIX_SIZE) + "_" + file.Filename
	clearFile := bytes.NewReader(file.Payload)
	_, err = s.s3.PutObject(
		context.Background(),
		s.config.Minio.Bucket,
		fileName,
		clearFile,
		file.PayloadSize,
		minio.PutObjectOptions{},
	)
	if err != nil {
		log.Warning(err)
		return &attach_proto.Nothing{Status: false}, err
	}

	_, err = s.db.AddAttachLink(context.Background(), &repository_proto.AddAttachLinkRequest{
		MailID:   request.MailID,
		Filename: fileName,
	})
	if err != nil {
		return &attach_proto.Nothing{Status: false}, err
	}

	return &attach_proto.Nothing{Status: true}, nil
}

func (s *AttachService) GetAttach(ctx context.Context, request *attach_proto.GetAttachRequest) (*attach_proto.AttachResponse, error) {
	respMail, err := s.db.GetMailInfoById(context.Background(), &repository_proto.GetMailInfoByIdRequest{
		MailId: request.MailID,
	})
	if err != nil {
		return &attach_proto.AttachResponse{}, err
	}

	var mail models.Mail
	err = easyjson.Unmarshal(respMail.Mail, &mail)
	if mail.Sender != request.Username {
		log.Warning(ErrAccess)
		return &attach_proto.AttachResponse{}, ErrAccess
	}

	respAttach, err := s.db.CheckAttachLink(context.Background(), &repository_proto.GetAttachRequest{
		MailID:   request.MailID,
		Filename: request.Filename,
	})
	if err != nil {
		log.Error(err)
		return &attach_proto.AttachResponse{}, err
	}
	if !respAttach.Status {
		log.Warning(ErrAccessAttach)
		return &attach_proto.AttachResponse{}, ErrAccessAttach
	}

	reader, err := s.s3.GetObject(
		context.Background(),
		s.config.Minio.Bucket,
		request.Filename,
		minio.GetObjectOptions{},
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer reader.Close()

	var file bytes.Buffer
	io.Copy(&file, reader)
	return &attach_proto.AttachResponse{
		File: file.Bytes(),
	}, nil

}

func (s *AttachService) ListAttach(ctx context.Context, request *attach_proto.GetAttachRequest) (*attach_proto.AttachListResponse, error) {
	respMail, err := s.db.GetMailInfoById(context.Background(), &repository_proto.GetMailInfoByIdRequest{
		MailId: request.MailID,
	})
	if err != nil {
		return &attach_proto.AttachListResponse{}, err
	}

	var mail models.Mail
	err = easyjson.Unmarshal(respMail.Mail, &mail)
	if err != nil {
		return &attach_proto.AttachListResponse{}, ErrJson
	}
	if mail.Sender != request.Username && mail.Addressee != request.Username {
		log.Warning(ErrAccess)
		return &attach_proto.AttachListResponse{}, ErrAccess
	}

	resp, err := s.db.ListAttaches(context.Background(), &repository_proto.GetAttachRequest{
		MailID:   request.MailID,
		Filename: "",
	})
	if err != nil {
		log.Warning(err)
		return &attach_proto.AttachListResponse{
			Filenames: nil,
		}, err
	}

	var attaches models.AttachList
	if err := easyjson.Unmarshal(resp.Filenames, &attaches); err != nil {
		log.Warning(err)
		return &attach_proto.AttachListResponse{
			Filenames: nil,
		}, err
	}

	for i := 0; i < len(attaches.Attaches); i++ {
		attaches.Attaches[i].Url = BASE_ATTACH_URL + attaches.Attaches[i].Filename
		attaches.Attaches[i].Filename = attaches.Attaches[i].Filename[ATTACH_PREFIX_SIZE+1:]
	}

	attachesBytes, err := easyjson.Marshal(attaches)
	if err != nil {
		log.Warning(err)
		return &attach_proto.AttachListResponse{
			Filenames: nil,
		}, err
	}

	return &attach_proto.AttachListResponse{
		Filenames: attachesBytes,
	}, nil
}

func (s *AttachService) CheckAttachPermission(ctx context.Context, request *attach_proto.AttachPermissionRequest) (*attach_proto.AttachPermissionResponse, error) {
	clearFileName := strings.TrimPrefix(request.FileUrl, BASE_ATTACH_URL)
	//log.Info(request.Username)
	//log.Info(clearFileName)

	resp, err := s.db.CheckAttachPermission(context.Background(), &repository_proto.AttachPermissionRequest{
		Username: request.Username,
		Filename: clearFileName,
	})
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	return &attach_proto.AttachPermissionResponse{
		Access: resp.Access,
	}, nil
}
