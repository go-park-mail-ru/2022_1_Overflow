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

func TestSignin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin"}, []func(http.ResponseWriter, *http.Request){d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	form := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, formBytes, fmt.Sprintf("%s/signin", srv.URL), http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBadSignin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signin"}, []func(http.ResponseWriter, *http.Request){d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	data := models.SignInForm{
		Username: "test",
		Password: "pass",
	}
	formBytes, _ := json.Marshal(data)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.WRONG_CREDS_ERR.Bytes(),
	}, nil)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signin", srv.URL), http.StatusBadRequest, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignup(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	data := models.SignUpForm{
		Lastname:             "John",
		Firstname:            "Doe",
		Username:             "ededededed",
		Password:             "pass",
		PasswordConfirmation: "pass",
	}
	formBytes, _ := json.Marshal(data)

	authUC.EXPECT().SignUp(context.Background(), &auth_proto.SignUpRequest{
		Form: formBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBadPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	data := models.SignUpForm{
		Lastname:             "John",
		Firstname:            "Doe",
		Username:             "ededededed",
		Password:             "pass",
		PasswordConfirmation: "passd",
	}
	formBytes, _ := json.Marshal(data)

	authUC.EXPECT().SignUp(context.Background(), &auth_proto.SignUpRequest{
		Form: formBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, "").Bytes(),
	}, nil)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)

	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusBadRequest, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestEmptyForm(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/signup"}, []func(http.ResponseWriter, *http.Request){d.SignUp},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	data := models.SignUpForm{
		Lastname:             "",
		Firstname:            "",
		Username:             "ededededed",
		Password:             "pass",
		PasswordConfirmation: "passd",
	}
	formBytes, _ := json.Marshal(data)

	authUC.EXPECT().SignUp(context.Background(), &auth_proto.SignUpRequest{
		Form: formBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, "").Bytes(),
	}, nil)

	srv := httptest.NewServer(router)
	defer srv.Close()

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	dataJson, _ := json.Marshal(data)
	_, err, token := Get(client, fmt.Sprintf("%s/signup", srv.URL), http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, fmt.Sprintf("%s/signup", srv.URL), http.StatusBadRequest, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSignout(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
	attachUC := attach_proto.NewMockAttachClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/logout", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SignOut, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/logout", srv.URL)

	form := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	formBytes, _ := json.Marshal(form)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, form, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, nil, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}
