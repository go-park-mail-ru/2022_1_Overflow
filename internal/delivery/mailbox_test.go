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
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/mail/send", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SendMail, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	sendUrl := fmt.Sprintf("%s/mail/send", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	mailData := models.MailForm{
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "test",
	}
	formBytesSignIn, _ := json.Marshal(signinForm)
	formBytesMailData, _ := json.Marshal(mailData)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: formBytesSignIn,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	data := &utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}
	mailboxUC.EXPECT().SendMail(context.Background(), &mailbox_proto.SendMailRequest{
		Data: data,
		Form: formBytesMailData,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)
	//&models.Session{Username: "test", Authenticated: true}, data)

	dataJson, _ := json.Marshal(mailData)

	_, err := Post(client, dataJson, sendUrl, http.StatusForbidden, "", "")
	if err != nil {
		t.Error(err)
		return
	}

	// ==============================================

	err = SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, sendUrl, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, sendUrl, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIncome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/mail/income", "/signin"}, []func(http.ResponseWriter, *http.Request){d.Income, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/income", srv.URL)

	signInForm := models.SignInForm{
		Username: "test2",
		Password: "test2",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	mails, _ := json.Marshal([]models.MailAdditional{})

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	incomeData := &utils_proto.Session{
		Username:      "test2",
		Authenticated: wrapperspb.Bool(true),
	}
	mailboxUC.EXPECT().Income(context.Background(), &mailbox_proto.IncomeRequest{
		Data:  incomeData,
		Limit: 100,
	}).Return(&mailbox_proto.ResponseMails{
		Response: &utils_proto.JsonResponse{Response: pkg.NO_ERR.Bytes()},
		Mails:    mails,
	}, nil)
	//Return(mails, pkg.NO_ERR)
	//&models.Session{Username:"test2", Authenticated: true}
	_, err, _ := Get(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	err = SigninUser(client, signInForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = Post(client, nil, url, http.StatusForbidden, "", "")
	if err != nil {
		t.Error(err)
		return
	}

	_, err, _ = Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestOutcome(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/mail/outcome", "/signin"}, []func(http.ResponseWriter, *http.Request){d.Outcome, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/outcome", srv.URL)

	signInForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signInFormBytes, _ := json.Marshal(signInForm)

	mails, _ := json.Marshal([]models.MailAdditional{})

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signInFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	outcomeData := &utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}
	mailboxUC.EXPECT().Outcome(context.Background(), &mailbox_proto.OutcomeRequest{
		Data:  outcomeData,
		Limit: 100,
	}).Return(&mailbox_proto.ResponseMails{
		Response: &utils_proto.JsonResponse{Response: pkg.NO_ERR.Bytes()},
		Mails:    mails,
	}, nil)
	//&models.Session{Username:"test", Authenticated: true}

	_, err, _ := Get(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	err = SigninUser(client, signInForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = Post(client, nil, url, http.StatusForbidden, "", "")
	if err != nil {
		t.Error(err)
		return
	}

	_, err, _ = Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRead(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/mail/read", "/signin"}, []func(http.ResponseWriter, *http.Request){d.ReadMail, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	signinForm := models.SignInForm{
		Username: "test2",
		Password: "test2",
	}
	signinFormBytes, _ := json.Marshal(signinForm)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	readMailData := &utils_proto.Session{
		Username:      "test2",
		Authenticated: wrapperspb.Bool(true),
	}
	//readMailDataBytes, _ := json.Marshal(readMailData)
	mailboxUC.EXPECT().ReadMail(context.Background(), &mailbox_proto.ReadMailRequest{
		Data: readMailData,
		Id:   1,
		Read: true,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)
	//&models.Session{Username: "test2", Authenticated: true}, int32(0)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	url := fmt.Sprintf("%s/mail/read?id=0", srv.URL)

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	form := models.ReadMailForm{
		Id:     1,
		IsRead: true,
	}
	formBytes, _ := json.Marshal(form)

	r, err := Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}

	var resp utils_proto.JsonResponse

	err = json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		t.Error(err)
		return
	}
}

func TestDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/mail/delete", "/signin"}, []func(http.ResponseWriter, *http.Request){d.DeleteMail, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/delete?id=0", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := json.Marshal(signinForm)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	deleteData := &utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}
	mailboxUC.EXPECT().DeleteMail(context.Background(), &mailbox_proto.DeleteMailRequest{
		Data: deleteData,
		Id:   1,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)
	//&models.Session{Username: "test", Authenticated: true}, int32(0)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	form := &models.DeleteMailForm{
		Id: 1,
	}
	formBytes, _ := json.Marshal(form)
	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authUC := auth_proto.NewMockAuthClient(mockCtrl)
	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
	profileUC := profile_proto.NewMockProfileClient(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(&d, []string{"/mail/get", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetMail, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/get?id=0", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := json.Marshal(signinForm)

	mail := models.Mail{
		Id:        0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}
	mailBytes, _ := json.Marshal(mail)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	getMailData := &utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}
	mailboxUC.EXPECT().GetMail(context.Background(), &mailbox_proto.GetMailRequest{
		Data: getMailData,
	}).Return(&mailbox_proto.ResponseMail{
		Mail: mailBytes,
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
	}, nil)

	//&models.Session{Username: "test", Authenticated: true}, int32(0)
	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	r, err, _ := Get(client, url, http.StatusOK)
	if err != nil {
		t.Error(err)
		return
	}

	var resp utils_proto.JsonResponse

	err = json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		t.Error(err)
		return
	}
}
