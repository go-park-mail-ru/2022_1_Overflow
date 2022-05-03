package folder_manager

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type FolderManagerService struct {
	config *config.Config
	db repository_proto.DatabaseRepositoryClient
	profile profile_proto.ProfileClient
}

func (s *FolderManagerService) IsOwner(context context.Context, data *utils_proto.Session, folderUserId int32) (bool, pkg.JsonResponse, error){
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	})
	if err != nil {
		return false, pkg.INTERNAL_ERR, err
	}
	if (resp.Response.Status != utils_proto.DatabaseStatus_OK) {
		return false, pkg.DB_ERR, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return false, pkg.JSON_ERR, err
	}
	return user.Id == folderUserId, pkg.NO_ERR, nil
}

func (s *FolderManagerService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	s.config = config
	s.db = db
	s.profile = profile
}

func (s *FolderManagerService) AddFolder(context context.Context, request *folder_manager_proto.AddFolderRequest) (*folder_manager_proto.ResponseFolder, error) {
	username := request.Data.Username
	log.Debug("Добавление папки, name = ", request.Name, ", username = ", username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
		}, err
	}
	if (user == models.User{}) {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.NO_USER_EXIST.Bytes(),
			},
		}, nil
	}
	resp2, err := s.db.GetFolderByName(context, &repository_proto.GetFolderByNameRequest{
		UserId: user.Id,
		FolderName: request.Name,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
	}
	var folder models.Folder
	err = json.Unmarshal(resp2.Folder, &folder)
	if err != nil {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
		}, err
	}
	if (folder != models.Folder{}) {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "Такая папка уже существует.").Bytes(),
			},
		}, err
	}
	resp3, err := s.db.AddFolder(context, &repository_proto.AddFolderRequest{
		Name: request.Name,
		UserId: user.Id,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
	}
	resp4, err := s.db.GetFolderByName(context, &repository_proto.GetFolderByNameRequest{
		UserId: user.Id,
		FolderName: request.Name,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, err
	}
	if resp4.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
	}
	return &folder_manager_proto.ResponseFolder{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Folder: resp4.Folder,
	}, nil
}

func (s *FolderManagerService) AddMailToFolder(context context.Context, request *folder_manager_proto.AddMailToFolderRequest) (*utils_proto.JsonResponse, error) {
	username := request.Data.Username
	log.Debug("Добавление письма в папку, folderId = ", request.FolderId, ", username = ", username, "move = ", request.Move)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (user == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_USER_EXIST.Bytes(),
		}, nil
	}
	isOwner, response, err := s.IsOwner(context, request.Data, request.FolderId)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.INTERNAL_ERR.Bytes(),
		}, err
	}
	if response != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: response.Bytes(),
		}, nil
	}
	if !isOwner {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp2, err := s.db.AddMailToFolder(context, &repository_proto.AddMailToFolderRequest{
		FolderId: request.FolderId,
		MailId: request.MailId,
		Move: request.Move,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *FolderManagerService) ChangeFolder(context context.Context, request *folder_manager_proto.ChangeFolderRequest) (*utils_proto.JsonResponse, error) {
	username := request.Data.Username
	log.Debug("Изменение имени папки, username = ", username, ", folderid = ", request.FolderId, ", folderNewName", request.FolderNewName)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (user == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_USER_EXIST.Bytes(),
		}, nil
	}
	isOwner, response, err := s.IsOwner(context, request.Data, request.FolderId)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.INTERNAL_ERR.Bytes(),
		}, err
	}
	if response != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: response.Bytes(),
		}, nil
	}
	if !isOwner {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp2, err := s.db.GetFolderById(context, &repository_proto.GetFolderByIdRequest{
		FolderId: request.FolderId,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var folder models.Folder
	err = json.Unmarshal(resp2.Folder, &folder)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if pkg.IsFolderReserved(folder.Name) {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp3, err := s.db.ChangeFolderName(context, &repository_proto.ChangeFolderNameRequest{
		FolderId: request.FolderId,
		NewName: request.FolderNewName,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *FolderManagerService) DeleteFolder(context context.Context, request *folder_manager_proto.DeleteFolderRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Удаление папки, name = ", request.Name, ", username = ", request.Data.Username)
	if pkg.IsFolderReserved(request.Name) {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (user == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_USER_EXIST.Bytes(),
		}, nil
	}
	resp2, err := s.db.GetFolderByName(context, &repository_proto.GetFolderByNameRequest{
		UserId: user.Id,
		FolderName: request.Name,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var folder models.Folder
	err = json.Unmarshal(resp2.Folder, &folder)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	isOwner, response, err := s.IsOwner(context, request.Data, folder.Id)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.INTERNAL_ERR.Bytes(),
		}, err
	}
	if response != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: response.Bytes(),
		}, nil
	}
	if !isOwner {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp3, err := s.db.DeleteFolder(context, &repository_proto.DeleteFolderRequest{
		FolderId: folder.Id,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *FolderManagerService) ListFolders(context context.Context, request *folder_manager_proto.ListFoldersRequest) (*folder_manager_proto.ResponseFolders, error) {
	log.Debug("Получение списка папок, username = ", request.Data.Username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Folders: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Folders: nil,
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Folders: nil,
		}, err
	}
	if (user == models.User{}) {
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.NO_USER_EXIST.Bytes(),
			},
			Folders: nil,
		}, nil
	}
	resp2, err := s.db.GetFoldersByUser(context, &repository_proto.GetFoldersByUserRequest{
		UserId: user.Id,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Folders: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Folders: nil,
		}, nil
	}
	if err != nil {
		return &folder_manager_proto.ResponseFolders{
			Response:&utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Folders: nil,
		}, err
	}
	return &folder_manager_proto.ResponseFolders{
		Response:&utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Folders: resp2.Folders,
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
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	if (user == models.User{}) {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.NO_USER_EXIST.Bytes(),
			},
			Mails: nil,
		}, nil
	}
	isOwner, response, err := s.IsOwner(context, request.Data, request.FolderId)
	if err != nil {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.INTERNAL_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	if response != pkg.NO_ERR {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: response.Bytes(),
			},
			Mails: nil,
		}, nil
	}
	if !isOwner {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.UNAUTHORIZED_ERR.Bytes(),
			},
			Mails: nil,
		}, nil
	}
	resp2, err := s.db.GetFolderMail(context, &repository_proto.GetFolderMailRequest{
		FolderId: request.FolderId,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, nil
	}
	var mails []models.Mail
	err = json.Unmarshal(resp2.Mails, &mails)
	if err != nil {
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	var mails_add []models.MailAdditional
	for _, mail := range mails {
		mail_add := models.MailAdditional{}
		mail_add.Mail = mail
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: mail.Sender},
		)
		if err != nil {
			return &folder_manager_proto.ResponseMails{
				Response:&utils_proto.JsonResponse{
					Response: pkg.DB_ERR.Bytes(),
				},
				Mails: nil,
			}, err
		}
		var response pkg.JsonResponse
		err = json.Unmarshal(resp.Response.Response, &response)
		if err != nil {
			return &folder_manager_proto.ResponseMails{
				Response:&utils_proto.JsonResponse{
					Response: pkg.JSON_ERR.Bytes(),
				},
				Mails: nil,
			}, nil
		}
		if response != pkg.NO_ERR {
			return &folder_manager_proto.ResponseMails{
				Response:&utils_proto.JsonResponse{
					Response: pkg.DB_ERR.Bytes(),
				},
				Mails: nil,
			}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add = append(mails_add, mail_add)
	}
	parsed, err := json.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response:&utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	return &folder_manager_proto.ResponseMails{
		Response:&utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Mails: parsed,
	}, nil
}

func (s *FolderManagerService) DeleteFolderMail(context context.Context, request *folder_manager_proto.DeleteFolderMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("Удаление письма из папки, folderId = ", request.FolderId, ", mailId = ", request.MailId, ", username = ", request.Data.Username)
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: request.Data.Username,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var user models.User
	err = json.Unmarshal(resp.User, &user)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (user == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.NO_USER_EXIST.Bytes(),
		}, nil
	}
	isOwner, response, err := s.IsOwner(context, request.Data, request.FolderId)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.INTERNAL_ERR.Bytes(),
		}, err
	}
	if response != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: response.Bytes(),
		}, nil
	}
	if !isOwner {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp2, err := s.db.DeleteFolderMail(context, &repository_proto.DeleteFolderMailRequest{
		Username: request.Data.Username,
		FolderId: request.FolderId,
		MailId: request.MailId,
		Restore: request.Restore,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}