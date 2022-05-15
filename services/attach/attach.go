package attach

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"OverflowBackend/proto/repository_proto"
	"bytes"
	"context"
	"encoding/json"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"io"
)

type AttachService struct {
	config *config.Config
	db     repository_proto.DatabaseRepositoryClient
	s3     *minio.Client
	attach_proto.UnimplementedAttachServer
}

func (s *AttachService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, s3 *minio.Client) {
	s.config = config
	s.db = db
	s.s3 = s3
}

func (s *AttachService) SaveAttach(ctx context.Context, request *attach_proto.SaveAttachRequest) (*attach_proto.Nothing, error) {
	var file models.Attach
	if err := json.Unmarshal(request.File, &file); err != nil {
		return &attach_proto.Nothing{Status: false}, err
	}

	fileName := pkg.RandSID(6) + "_" + file.Filename
	clearFile := bytes.NewReader(file.Payload)
	_, err := s.s3.PutObject(
		context.Background(),
		s.config.Minio.Bucket,
		fileName,
		clearFile,
		file.PayloadSize,
		minio.PutObjectOptions{},
	)
	if err != nil {
		log.Warning(err)
	}

	return &attach_proto.Nothing{Status: true}, nil
}

func (s *AttachService) GetAttach(ctx context.Context, request *attach_proto.GetAttachRequest) (*attach_proto.AttachResponse, error) {
	reader, err := s.s3.GetObject(
		context.Background(),
		s.config.Minio.Bucket,
		request.AttachID,
		minio.GetObjectOptions{},
	)
	defer reader.Close()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var file bytes.Buffer
	io.Copy(&file, reader)
	return &attach_proto.AttachResponse{
		File: file.Bytes(),
	}, nil

	return nil, nil
}
