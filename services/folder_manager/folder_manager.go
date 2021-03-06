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
	"time"

	"github.com/mailru/easyjson"

	log "github.com/sirupsen/logrus"
)

type FolderManagerService struct {
	config  *config.Config
	db      repository_proto.DatabaseRepositoryClient
	profile profile_proto.ProfileClient
}

func (s *FolderManagerService) FolderExists(context context.Context, userId int32, folderName string) bool {
	resp, err := s.db.GetFolderByName(context, &repository_proto.GetFolderByNameRequest{
		UserId:     userId,
		FolderName: folderName,
	})
	if err != nil || resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return false
	}
	var folder models.Folder
	err = easyjson.Unmarshal(resp.Folder, &folder)
	if err != nil {
		return false
	}
	return (folder != models.Folder{})
}

func (s *FolderManagerService) GetValidateUser(context context.Context, username string) (models.User, pkg.JsonResponse, error) {
	var user models.User
	resp, err := s.db.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
		Username: username,
	})
	if err != nil {
		log.Error(err)
		return user, pkg.DB_ERR, err
	}
	if resp.Response.Status != utils_proto.DatabaseStatus_OK {
		return user, pkg.DB_ERR, nil
	}
	err = easyjson.Unmarshal(resp.User, &user)
	if err != nil {
		return user, pkg.JSON_ERR, err
	}
	if (user == models.User{}) {
		return user, pkg.NO_USER_EXIST, nil
	}
	return user, pkg.NO_ERR, nil
}

func (s *FolderManagerService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient, profile profile_proto.ProfileClient) {
	s.config = config
	s.db = db
	s.profile = profile
}

func (s *FolderManagerService) AddFolder(context context.Context, request *folder_manager_proto.AddFolderRequest) (*folder_manager_proto.ResponseFolder, error) {
	username := request.Data.Username
	log.Debug("???????????????????? ??????????, name = ", request.Name, ", username = ", username)
	user, resp, err := s.GetValidateUser(context, username)
	if err != nil || resp != pkg.NO_ERR {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: resp.Bytes(),
			},
		}, err
	}
	if s.FolderExists(context, user.Id, request.Name) {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ?????? ????????????????????.").Bytes(),
			},
		}, nil
	}
	name, err := pkg.ValidateFormatFolderName(request.Name)
	if err != nil {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()).Bytes(),
			},
		}, nil
	}
	request.Name = name
	resp3, err := s.db.AddFolder(context, &repository_proto.AddFolderRequest{
		Name:   request.Name,
		UserId: user.Id,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
	}
	if resp3.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
	}
	resp4, err := s.db.GetFolderByName(context, &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: request.Name,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolder{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
		}, nil
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

func (s *FolderManagerService) AddMailToFolderById(context context.Context, request *folder_manager_proto.AddMailToFolderByIdRequest) (*utils_proto.JsonResponse, error) {
	username := request.Data.Username
	log.Debug("???????????????????? ???????????? ?? ??????????, folderName = ", request.FolderName, ", username = ", username, ", mailId = ", request.MailId, ", move = ", request.Move)
	user, resp, err := s.GetValidateUser(context, username)
	if err != nil || resp != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if !s.FolderExists(context, user.Id, request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	resp2, err := s.db.AddMailToFolderById(context, &repository_proto.AddMailToFolderByIdRequest{
		UserId:     user.Id,
		FolderName: request.FolderName,
		MailId:     request.MailId,
		Move:       request.Move,
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

func (s *FolderManagerService) AddMailToFolderByObject(context context.Context, request *folder_manager_proto.AddMailToFolderByObjectRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("???????????????????? ???????????? ?? ??????????, folderName = ", request.FolderName, ", username = ", request.Data.Username)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if !s.FolderExists(context, user.Id, request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	var form models.MailForm
	err = easyjson.Unmarshal(request.Form, &form)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	mail := models.Mail{
		Sender:    request.Data.Username,
		Addressee: form.Addressee,
		Theme:     form.Theme,
		Text:      form.Text,
		Files:     form.Files,
		Date:      time.Now(),
	}
	mailBytes, _ := easyjson.Marshal(mail)
	resp3, err := s.db.AddMailToFolderByObject(context, &repository_proto.AddMailToFolderByObjectRequest{
		UserId:     user.Id,
		FolderName: request.FolderName,
		Mail:       mailBytes,
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

func (s *FolderManagerService) MoveFolderMail(context context.Context, request *folder_manager_proto.MoveFolderMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("?????????????????????? ???????????? ???? ?????????? ?? ??????????, username = ", request.Data.Username, ", folderNameSrc = ", request.FolderNameSrc, ", folderNameDest = ", request.FolderNameDest)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if request.FolderNameSrc == pkg.FOLDER_DRAFTS {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNAUTHORIZED, "???????????? ???????????????????? ???????????? ???? ?????????? ?? ??????????????????????.").Bytes(),
		}, nil
	}
	if !s.FolderExists(context, user.Id, request.FolderNameSrc) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	if !s.FolderExists(context, user.Id, request.FolderNameDest) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ???????????????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	resp2, err := s.db.MoveFolderMail(context, &repository_proto.MoveFolderMailRequest{
		UserId:         user.Id,
		FolderNameSrc:  request.FolderNameSrc,
		FolderNameDest: request.FolderNameDest,
		MailId:         request.MailId,
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
	log.Debug("?????????????????? ?????????? ??????????, username = ", username, ", folderName = ", request.FolderName, ", folderNewName", request.FolderNewName)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if pkg.IsFolderReserved(request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	if s.FolderExists(context, user.Id, request.FolderNewName) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ?????? ????????????????????.").Bytes(),
		}, nil
	}
	name, err := pkg.ValidateFormatFolderName(request.FolderNewName)
	if err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()).Bytes(),
		}, nil
	}
	request.FolderNewName = name
	resp3, err := s.db.ChangeFolderName(context, &repository_proto.ChangeFolderNameRequest{
		UserId:     user.Id,
		FolderName: request.FolderName,
		NewName:    request.FolderNewName,
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
	log.Debug("???????????????? ??????????, folderName = ", request.FolderName, ", username = ", request.Data.Username)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if !s.FolderExists(context, user.Id, request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	if pkg.IsFolderReserved(request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.UNAUTHORIZED_ERR.Bytes(),
		}, nil
	}
	resp3, err := s.db.DeleteFolder(context, &repository_proto.DeleteFolderRequest{
		UserId:     user.Id,
		FolderName: request.FolderName,
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
	log.Debug("?????????????????? ???????????? ??????????, username = ", request.Data.Username)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &folder_manager_proto.ResponseFolders{
			Response: &utils_proto.JsonResponse{
				Response: resp.Bytes(),
			},
			Folders: nil,
		}, err
	}
	resp2, err := s.db.GetFoldersByUser(context, &repository_proto.GetFoldersByUserRequest{
		UserId:       user.Id,
		Limit:        request.Limit,
		Offset:       request.Offset,
		ShowReserved: request.ShowReserved,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseFolders{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Folders: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseFolders{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Folders: nil,
		}, nil
	}
	if err != nil {
		return &folder_manager_proto.ResponseFolders{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Folders: nil,
		}, err
	}
	return &folder_manager_proto.ResponseFolders{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Folders: resp2.Folders,
	}, nil
}

func (s *FolderManagerService) ListFolder(context context.Context, request *folder_manager_proto.ListFolderRequest) (*folder_manager_proto.ResponseMails, error) {
	log.Debug("?????????????????? ???????????? ?????????? ???? ??????????, username = ", request.Data.Username, ", folderName = ", request.FolderName)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &folder_manager_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: resp.Bytes(),
			},
			Mails: nil,
		}, err
	}
	if !s.FolderExists(context, user.Id, request.FolderName) {
		return &folder_manager_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ???? ????????????????????.").Bytes(),
			},
			Mails: nil,
		}, nil
	}
	resp2, err := s.db.GetFolderMail(context, &repository_proto.GetFolderMailRequest{
		UserId:     user.Id,
		FolderName: request.FolderName,
		Limit:      request.Limit,
		Offset:     request.Offset,
	})
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	if resp2.Response.Status != utils_proto.DatabaseStatus_OK {
		return &folder_manager_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, nil
	}
	var mails models.MailList
	err = easyjson.Unmarshal(resp2.Mails, &mails)
	if err != nil {
		return &folder_manager_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.JSON_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	var mails_add models.MailAddList
	mails_add.Amount = mails.Amount
	for _, mail := range mails.Mails {
		mail_add := models.MailAdditional{}
		mail_add.Mail = mail
		var username string
		if mail.Sender == user.Username {
			username = mail.Addressee
		} else {
			username = mail.Sender
		}
		resp, err := s.profile.GetAvatar(
			context,
			&profile_proto.GetAvatarRequest{Username: username, DummyName: request.DummyName},
		)
		if err != nil {
			return &folder_manager_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.DB_ERR.Bytes(),
				},
				Mails: nil,
			}, err
		}
		var response pkg.JsonResponse
		err = easyjson.Unmarshal(resp.Response.Response, &response)
		if err != nil {
			return &folder_manager_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.JSON_ERR.Bytes(),
				},
				Mails: nil,
			}, nil
		}
		if response != pkg.NO_ERR {
			return &folder_manager_proto.ResponseMails{
				Response: &utils_proto.JsonResponse{
					Response: pkg.DB_ERR.Bytes(),
				},
				Mails: nil,
			}, nil
		}
		mail_add.AvatarUrl = resp.Url
		mails_add.Mails = append(mails_add.Mails, mail_add)
	}
	parsed, err := easyjson.Marshal(mails_add)
	if err != nil {
		log.Error(err)
		return &folder_manager_proto.ResponseMails{
			Response: &utils_proto.JsonResponse{
				Response: pkg.DB_ERR.Bytes(),
			},
			Mails: nil,
		}, err
	}
	return &folder_manager_proto.ResponseMails{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Mails: parsed,
	}, nil
}

func (s *FolderManagerService) UpdateFolderMail(context context.Context, request *folder_manager_proto.UpdateFolderMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("???????????????????? ???????????? ?? ??????????, folderName = ", request.FolderName, ", mailId = ", request.MailId, ", username = ", request.Data.Username)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		if err != nil {
			log.Error(err)
		}
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if !s.FolderExists(context, user.Id, request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	resp2, err := s.db.IsMailMoved(context, &repository_proto.IsMailMovedRequest{
		UserId: user.Id,
		MailId: request.MailId,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if !resp2.Moved {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNAUTHORIZED, "???????????????? ?? ??????????????: ???????????? ???? ???????????????? ???????????????????????? ?? ??????????.").Bytes(),
		}, nil
	}
	resp3, err := s.db.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
		MailId: request.MailId,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp3.Response.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	var mail models.Mail
	err = easyjson.Unmarshal(resp3.Mail, &mail)
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if mail.Sender != user.Username {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_UNAUTHORIZED, "???????????????? ?? ??????????????: ?????????????????? ???????????? ???????????? ?????????? ???????????? ??????????????????????.").Bytes(),
		}, nil
	}
	mailFormBytes := request.MailForm
	var mailForm models.MailForm
	err = easyjson.Unmarshal(mailFormBytes, &mailForm)
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	mail.Date = time.Now()
	mail.Addressee = mailForm.Addressee
	mail.Files = mailForm.Files
	mail.Text = mailForm.Text
	mail.Theme = mailForm.Theme
	mailBytes, err := easyjson.Marshal(mail)
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	resp4, err := s.db.UpdateMail(context, &repository_proto.UpdateMailRequest{
		UserId: user.Id,
		MailId: request.MailId,
		Mail:   mailBytes,
	})
	if err != nil {
		log.Error(err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp4.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *FolderManagerService) DeleteFolderMail(context context.Context, request *folder_manager_proto.DeleteFolderMailRequest) (*utils_proto.JsonResponse, error) {
	log.Debug("???????????????? ???????????? ???? ??????????, folderName = ", request.FolderName, ", mailId = ", request.MailId, ", username = ", request.Data.Username)
	user, resp, err := s.GetValidateUser(context, request.Data.Username)
	if err != nil || resp != pkg.NO_ERR {
		return &utils_proto.JsonResponse{
			Response: resp.Bytes(),
		}, err
	}
	if !s.FolderExists(context, user.Id, request.FolderName) {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_OBJECT_EXISTS, "?????????? ?????????? ???? ????????????????????.").Bytes(),
		}, nil
	}
	resp2, err := s.db.DeleteFolderMail(context, &repository_proto.DeleteFolderMailRequest{
		UserId: user.Id,
		MailId: request.MailId,
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
