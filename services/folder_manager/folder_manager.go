package folder_manager

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type FolderManagerService struct {
	config *config.Config
	db repository_proto.DatabaseRepositoryClient
	profile profile_proto.ProfileClient
}

func (s *FolderManagerService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	s.config = config
	s.db = db
	s.profile = profile
}

func (s *FolderManagerService) AddFolder(context context.Context, request *folder_manager_proto.AddFolderRequest) (*utils_proto.JsonResponse, error) {
	username := request.Data.Username
	log.Debug("Добавление папки, name = ", request.Name, ", username = ", username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	if proto.Equal(resp.User, &utils_proto.User{}) {
		return &pkg.NO_USER_EXIST, nil
	}
	resp2, err := s.db.AddFolder(context, &repository_proto.AddFolderRequest{
		Name: request.Name,
		UserId: resp.User.Id,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}

func (s *FolderManagerService) AddMailToFolder(context context.Context, request *folder_manager_proto.AddMailToFolderRequest) (*utils_proto.JsonResponse, error) {
	username := request.Data.Username
	log.Debug("Добавление письма в папку, folderId = ", request.FolderId, ", username = ", username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	if proto.Equal(resp.User, &utils_proto.User{}) {
		return &pkg.NO_USER_EXIST, nil
	}
	resp2, err := s.db.AddMailToFolder(context, &repository_proto.AddMailToFolderRequest{
		FolderId: request.FolderId,
		MailId: request.MailId,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}

func (s *FolderManagerService) ChangeFolder(context context.Context, request *folder_manager_proto.ChangeFolderRequest) (*utils_proto.JsonResponse, error) {
	username := request.Data.Username
	log.Debug("Изменение имени папки, username = ", username, ", folderid = ", request.FolderId, ", folderNewName", request.FolderNewName)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	if proto.Equal(resp.User, &utils_proto.User{}) {
		return &pkg.NO_USER_EXIST, nil
	}
	resp2, err := s.db.ChangeFolderName(context, &repository_proto.ChangeFolderNameRequest{
		FolderId: request.FolderId,
		NewName: request.FolderNewName,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}

func (s *FolderManagerService) DeleteFolder(context context.Context, request *folder_manager_proto.DeleteFolderRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Удаление папки, name = ", request.Name, ", username = ", request.Data.Username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	if proto.Equal(resp.User, &utils_proto.User{}) {
		return &pkg.NO_USER_EXIST, nil
	}
	resp2, err := s.db.GetFolderByName(context, &repository_proto.GetFolderByNameRequest{
		UserId: resp.User.Id,
		FolderName: request.Name,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	resp3, err := s.db.DeleteFolder(context, &repository_proto.DeleteFolderRequest{
		FolderId: resp2.Folder.Id,
	})
	if err != nil {
		log.Error(err)
		return &pkg.DB_ERR, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &pkg.DB_ERR, nil
	}
	return &pkg.NO_ERR, nil
}

func (s *FolderManagerService) ListFolders(context context.Context, request *folder_manager_proto.ListFoldersRequest) (*folder_manager_proto.ResponseFolders, error) {
	log.Debug("Получение списка папок, username = ", request.Data.Username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolders{
			Response: &pkg.DB_ERR,
			Folders: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolders{
			Response: &pkg.DB_ERR,
			Folders: nil,
		}, nil
	}
	if proto.Equal(resp.User, &utils_proto.User{}) {
		return &folder_manager_proto.ResponseFolders{
			Response: &pkg.NO_USER_EXIST,
			Folders: nil,
		}, nil
	}
	resp2, err := s.db.GetFoldersByUser(context, &repository_proto.GetFoldersByUserRequest{
		UserId: resp.User.Id,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolders{
			Response: &pkg.DB_ERR,
			Folders: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolders{
			Response: &pkg.DB_ERR,
			Folders: nil,
		}, nil
	}
	foldersBytes, err := json.Marshal(resp2.Folders)
	if err != nil {
		return &folder_manager_proto.ResponseFolders{
			Response: &pkg.JSON_ERR,
			Folders: nil,
		}, nil
	}
	return &folder_manager_proto.ResponseFolders{
		Response: &pkg.NO_ERR,
		Folders: foldersBytes,
	}, nil
}

func (s *FolderManagerService) ListFolder(context context.Context, request *folder_manager_proto.ListFolderRequest) (*folder_manager_proto.ResponseMails, error) {
	log.Debug("Получение списка писем из папки, username = ", request.Data.Username, ", folderId = ", request.FolderId)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response: &pkg.DB_ERR,
			Mails: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseMails{
			Response: &pkg.DB_ERR,
			Mails: nil,
		}, nil
	}
	if proto.Equal(resp.User, &utils_proto.User{}) {
		return &folder_manager_proto.ResponseMails{
			Response: &pkg.NO_USER_EXIST,
			Mails: nil,
		}, nil
	}
	resp2, err := s.db.GetFolderMail(context, &repository_proto.GetFolderMailRequest{
		FolderId: request.FolderId,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response: &pkg.DB_ERR,
			Mails: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseMails{
			Response: &pkg.DB_ERR,
			Mails: nil,
		}, nil
	}
	var mails_add []*utils_proto.MailAdditional
	for _, mail := range resp2.Mails {
		mail_add := utils_proto.MailAdditional{}
		mail_add.Mail = mail
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: mail.Sender},
		)
		if err != nil {
			return &folder_manager_proto.ResponseMails{
				Response: &pkg.DB_ERR,
				Mails: nil,
			}, err
		}
		if !proto.Equal(resp.Response, &pkg.NO_ERR) {
			return &folder_manager_proto.ResponseMails{
				Response: &pkg.DB_ERR,
				Mails: nil,
			}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add = append(mails_add, &mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response: &pkg.DB_ERR,
			Mails: nil,
		}, err
	}
	return &folder_manager_proto.ResponseMails{
		Response: &pkg.NO_ERR,
		Mails: parsed,
	}, err
}