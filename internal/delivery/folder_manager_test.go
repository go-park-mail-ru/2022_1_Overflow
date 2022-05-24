package delivery_test

import (
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/add"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.AddFolder},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/add", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	form := models.AddFolderForm{
		FolderName: "folder",
	}
	formBytes, _ := json.Marshal(form)

	folderManagerUC.EXPECT().AddFolder(context.Background(), &folder_manager_proto.AddFolderRequest{
		Data: &utils_proto.Session{
			Username:      "test",
			Authenticated: true,
		},
		Name: form.FolderName,
	}).Return(&folder_manager_proto.ResponseFolder{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Folder: nil,
	}, nil)

	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAddMailToFolderById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/mail/add"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.AddMailToFolderById},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/mail/add", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	folderManagerUC.EXPECT().AddMailToFolderById(context.Background(), &folder_manager_proto.AddMailToFolderByIdRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderName: "folder",
		MailId:     1,
		Move:       true,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	form := models.AddMailToFolderByIdForm{
		FolderName: "folder",
		MailId:     1,
		Move:       true,
	}
	formBytes, _ := json.Marshal(form)

	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAddMailToFolderByObject(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/mail/add_form"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.AddMailToFolderByObject},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/mail/add_form", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	mailForm := models.MailForm{
		Addressee: "test",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
	}
	mailFormBytes, _ := json.Marshal(mailForm)

	form := models.AddMailToFolderByObjectForm{
		FolderName: "folder",
		Mail:       mailForm,
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	folderManagerUC.EXPECT().AddMailToFolderByObject(context.Background(), &folder_manager_proto.AddMailToFolderByObjectRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderName: "folder",
		Form:       mailFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMoveFolderMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/mail/move"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.MoveFolderMail},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/mail/move", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	folderNameSrc := "folder"
	folderNameDest := "folder1"

	form := models.MoveFolderMailForm{
		FolderNameSrc:  folderNameSrc,
		FolderNameDest: folderNameDest,
		MailId:         1,
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	folderManagerUC.EXPECT().MoveFolderMail(context.Background(), &folder_manager_proto.MoveFolderMailRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderNameSrc:  folderNameSrc,
		FolderNameDest: folderNameDest,
		MailId:         1,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestChangeFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/rename"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.ChangeFolder},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/rename", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	folderName := "folder"
	folderNameNew := "folder1"

	form := models.ChangeFolderForm{
		FolderName:    folderName,
		NewFolderName: folderNameNew,
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	folderManagerUC.EXPECT().ChangeFolder(context.Background(), &folder_manager_proto.ChangeFolderRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderName:    folderName,
		FolderNewName: folderNameNew,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/delete"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.DeleteFolder},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/delete", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	folderName := "folder"

	form := models.DeleteFolderForm{
		FolderName: folderName,
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	folderManagerUC.EXPECT().DeleteFolder(context.Background(), &folder_manager_proto.DeleteFolderRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderName: form.FolderName,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteFolderMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/mail/delete"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.DeleteFolderMail},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/mail/delete", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	folderName := "folder"

	form := models.DeleteFolderMailForm{
		FolderName: folderName,
		MailId:     1,
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	folderManagerUC.EXPECT().DeleteFolderMail(context.Background(), &folder_manager_proto.DeleteFolderMailRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderName: folderName,
		MailId:     1,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestListFolders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/list"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.ListFolders},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/folder/list", srv.URL)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	folderName := "folder"

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	folderManagerUC.EXPECT().ListFolders(context.Background(), &folder_manager_proto.ListFoldersRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		Limit:  100,
		Offset: 0,
	}).Return(&folder_manager_proto.ResponseFolders{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Folders: nil,
	}, nil)
	folderManagerUC.EXPECT().ListFolder(context.Background(), &folder_manager_proto.ListFolderRequest{
		Data: &utils_proto.Session{
			Username:      signInForm.Username,
			Authenticated: true,
		},
		FolderName: folderName,
		Limit:      100,
		Offset:     0,
	}).Return(&folder_manager_proto.ResponseMails{
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
		Mails: nil,
	}, nil)

	err := SigninUser(client, signInForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, _ = Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, _ = Get(client, fmt.Sprintf("%s/folder/list?folder_name=%s", srv.URL, folderName), http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
}
