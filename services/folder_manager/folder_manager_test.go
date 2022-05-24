package folder_manager_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"OverflowBackend/services/folder_manager"
	"context"
	"encoding/json"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func InitTestUseCase(ctrl *gomock.Controller) (*repository_proto.MockDatabaseRepositoryClient, *profile_proto.MockProfileClient, *folder_manager_proto.MockFolderManagerClient, *folder_manager.FolderManagerService) {
	current := time.Now()
	monkey.Patch(time.Now, func() time.Time { return current })
	log.SetLevel(log.FatalLevel)
	db := repository_proto.NewMockDatabaseRepositoryClient(ctrl)
	profile := profile_proto.NewMockProfileClient(ctrl)
	folder := folder_manager_proto.NewMockFolderManagerClient(ctrl)
	uc := folder_manager.FolderManagerService{}
	uc.Init(config.TestConfig(), db, profile)
	return db, profile, folder, &uc
}

func TestAddFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folder := models.Folder{
		Id:        0,
		Name:      "folder",
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderBytes, _ := json.Marshal(folder)

	folderEmptyBytes, _ := json.Marshal(models.Folder{})

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folder.Name,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderEmptyBytes,
	}, nil)

	mockDB.EXPECT().AddFolder(context.Background(), &repository_proto.AddFolderRequest{
		Name:   folder.Name,
		UserId: user.Id,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folder.Name,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil)

	resp, err := uc.AddFolder(context.Background(), &folder_manager_proto.AddFolderRequest{
		Data: &session,
		Name: folder.Name,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestAddMailToFolderById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folderName := "folder"
	var mailId int32 = 0
	move := true

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().AddMailToFolderById(context.Background(), &repository_proto.AddMailToFolderByIdRequest{
		UserId:     user.Id,
		FolderName: folderName,
		MailId:     mailId,
		Move:       move,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	//folder.EXPECT().FolderExist(context.Background(), user.Id, folderName).Return(true)

	fold := &models.Folder{
		Id:     0,
		Name:   "folder",
		UserId: 0,
	}
	foldBytes, _ := json.Marshal(fold)
	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		FolderName: folderName,
	}).Return(&repository_proto.ResponseFolder{
		Folder: foldBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil)

	resp, err := uc.AddMailToFolderById(context.Background(), &folder_manager_proto.AddMailToFolderByIdRequest{
		Data:       &session,
		FolderName: folderName,
		MailId:     mailId,
		Move:       move,
	})

	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestAddMailToFolderByObject(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	form := models.MailForm{
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
	}
	formBytes, _ := json.Marshal(form)

	folderName := "folder"

	mail := models.Mail{
		Id:        0,
		Sender:    session.Username,
		Addressee: form.Addressee,
		Theme:     form.Theme,
		Text:      form.Text,
		Files:     form.Files,
		Date:      time.Now(),
		Read:      false,
	}
	mailBytes, _ := json.Marshal(mail)

	folder := models.Folder{
		Id:        0,
		Name:      "folder",
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderBytes, _ := json.Marshal(folder)
	//folderEmptyBytes, _ := json.Marshal(models.Folder{})

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folder.Name,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().AddMailToFolderByObject(context.Background(), &repository_proto.AddMailToFolderByObjectRequest{
		UserId:     user.Id,
		FolderName: folderName,
		Mail:       mailBytes,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.AddMailToFolderByObject(context.Background(), &folder_manager_proto.AddMailToFolderByObjectRequest{
		Data:       &session,
		FolderName: folderName,
		Form:       formBytes,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestMoveFolderMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	var mailId int32 = 0

	folderNameSrc := "folder"
	folderNameDest := "folder2"

	folderSrc := models.Folder{
		Id:        0,
		Name:      folderNameSrc,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderSrcBytes, _ := json.Marshal(folderSrc)

	folderDest := models.Folder{
		Id:        0,
		Name:      folderNameDest,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderDestBytes, _ := json.Marshal(folderDest)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folderNameSrc,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderSrcBytes,
	}, nil)

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folderNameDest,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderDestBytes,
	}, nil)

	mockDB.EXPECT().MoveFolderMail(context.Background(), &repository_proto.MoveFolderMailRequest{
		UserId:         user.Id,
		FolderNameSrc:  folderNameSrc,
		FolderNameDest: folderNameDest,
		MailId:         mailId,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.MoveFolderMail(context.Background(), &folder_manager_proto.MoveFolderMailRequest{
		Data:           &session,
		FolderNameSrc:  folderNameSrc,
		FolderNameDest: folderNameDest,
		MailId:         mailId,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestChangeFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folderName := "folder"
	folderNameNew := "folder2"

	/*
		folder := models.Folder{
			Id: 0,
			Name: folderName,
			UserId: user.Id,
			CreatedAt: time.Now(),
		}
		folderBytes, _ := json.Marshal(folder)
	*/

	folderNew := models.Folder{}
	folderNewBytes, _ := json.Marshal(folderNew)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folderNameNew,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderNewBytes,
	}, nil)

	mockDB.EXPECT().ChangeFolderName(context.Background(), &repository_proto.ChangeFolderNameRequest{
		UserId:     user.Id,
		FolderName: folderName,
		NewName:    folderNameNew,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.ChangeFolder(context.Background(), &folder_manager_proto.ChangeFolderRequest{
		Data:          &session,
		FolderName:    folderName,
		FolderNewName: folderNameNew,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestDeleteFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folderName := "folder"

	folder := models.Folder{
		Id:        0,
		Name:      "folder",
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderBytes, _ := json.Marshal(folder)
	//folderEmptyBytes, _ := json.Marshal(models.Folder{})

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folderName,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().DeleteFolder(context.Background(), &repository_proto.DeleteFolderRequest{
		UserId:     user.Id,
		FolderName: folderName,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.DeleteFolder(context.Background(), &folder_manager_proto.DeleteFolderRequest{
		Data:       &session,
		FolderName: folderName,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestListFolders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folders := models.FolderList{
		Amount: 1,
		Folders: []models.Folder{
			{
				Id:        0,
				Name:      "folder",
				UserId:    user.Id,
				CreatedAt: time.Now(),
			},
		},
	}
	foldersBytes, _ := json.Marshal(folders)

	var limit int32 = 10
	var offset int32 = 0

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().GetFoldersByUser(context.Background(), &repository_proto.GetFoldersByUserRequest{
		UserId: user.Id,
		Limit:  limit,
		Offset: offset,
	}).Return(&repository_proto.ResponseFolders{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folders: foldersBytes,
	}, nil)

	resp, err := uc.ListFolders(context.Background(), &folder_manager_proto.ListFoldersRequest{
		Data:   &session,
		Limit:  limit,
		Offset: offset,
	})

	var response pkg.JsonResponse
	var respFolders models.FolderList
	json_err := json.Unmarshal(resp.Response.Response, &response)
	json_folders_err := json.Unmarshal(resp.Folders, &respFolders)
	if err != nil || json_err != nil || json_folders_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestListFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folderName := "folder"

	folder := models.Folder{
		Id:        0,
		Name:      "folder",
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderBytes, _ := json.Marshal(folder)
	//folderEmptyBytes, _ := json.Marshal(models.Folder{})

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folderName,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil)

	/*
		mailsAdd := models.MailAddList{
			Amount: 0,
			Mails: []models.MailAdditional{
			},
		}
		mailsAddBytes, _ := json.Marshal(mailsAdd)
	*/

	mails := models.MailList{
		Amount: 0,
		Mails:  []models.Mail{},
	}
	mailsBytes, _ := json.Marshal(mails)

	var limit int32 = 10
	var offset int32 = 0

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().GetFolderMail(context.Background(), &repository_proto.GetFolderMailRequest{
		UserId:     user.Id,
		FolderName: folderName,
		Limit:      limit,
		Offset:     offset,
	}).Return(&repository_proto.ResponseMails{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mails: mailsBytes,
	}, nil)

	resp, err := uc.ListFolder(context.Background(), &folder_manager_proto.ListFolderRequest{
		Data:       &session,
		FolderName: folderName,
		Limit:      limit,
		Offset:     offset,
	})

	var response pkg.JsonResponse
	var respMailsAdd models.MailAddList
	json_err := json.Unmarshal(resp.Response.Response, &response)
	json_mails_err := json.Unmarshal(resp.Mails, &respMailsAdd)
	if err != nil || json_err != nil || json_mails_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}

func TestDeleteFolderMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB, _, _, uc := InitTestUseCase(mockCtrl)

	user := models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Username:  "test",
		Password:  "test",
	}
	userBytes, _ := json.Marshal(user)

	session := utils_proto.Session{
		Username:      user.Username,
		Authenticated: true,
	}

	folderName := "folder"
	var mailId int32 = 0

	folder := models.Folder{
		Id:        0,
		Name:      "folder",
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	folderBytes, _ := json.Marshal(folder)
	//folderEmptyBytes, _ := json.Marshal(models.Folder{})

	mockDB.EXPECT().GetFolderByName(context.Background(), &repository_proto.GetFolderByNameRequest{
		UserId:     user.Id,
		FolderName: folderName,
	}).Return(&repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil)

	mockDB.EXPECT().GetUserInfoByUsername(context.Background(), &repository_proto.GetUserInfoByUsernameRequest{
		Username: user.Username,
	}).Return(&repository_proto.ResponseUser{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		User: userBytes,
	}, nil)

	mockDB.EXPECT().DeleteFolderMail(context.Background(), &repository_proto.DeleteFolderMailRequest{
		UserId: user.Id,
		MailId: mailId,
	}).Return(&utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil)

	resp, err := uc.DeleteFolderMail(context.Background(), &folder_manager_proto.DeleteFolderMailRequest{
		Data:       &session,
		FolderName: folderName,
		MailId:     mailId,
	})
	var response pkg.JsonResponse
	json_err := json.Unmarshal(resp.Response, &response)
	if err != nil || json_err != nil || response != pkg.NO_ERR {
		t.Errorf("Неверный ответ от UseCase.")
		return
	}
}
