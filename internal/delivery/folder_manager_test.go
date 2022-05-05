package delivery_test

import (
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestAddFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/add"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.AddFolder},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

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
			Username: "test",
			Authenticated: wrapperspb.Bool(true),
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

func TestAddMailToFolder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin", "/folder/mail/add"}, []func(http.ResponseWriter, *http.Request){d.SignIn, d.AddMailToFolder},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

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

	folderManagerUC.EXPECT().AddMailToFolder(context.Background(), &folder_manager_proto.AddMailToFolderRequest{
		Data: &utils_proto.Session{
			Username: signInForm.Username,
			Authenticated: wrapperspb.Bool(true),
		},
		FolderName: "folder",
		MailId: 1,
		Move: true,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	form := models.AddMailToFolderForm{
		FolderName: "folder",
		MailId: 1,
		Move: true,
	}
	formBytes, _ := json.Marshal(form)

	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}