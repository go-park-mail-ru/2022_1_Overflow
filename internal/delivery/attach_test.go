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
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

func TestListAttach(t *testing.T) {
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
	router := InitTestRouter(&d, []string{"/mail/attach/list", "/signin"}, []func(http.ResponseWriter, *http.Request){d.ListAttach, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/attach/list", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := easyjson.Marshal(signinForm)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	//sessionData := &utils_proto.Session{
	//	Username:      "test",
	//	Authenticated: true,
	//}

	filenames := models.AttachList{
		Attaches: []models.AttachShort{{
			Filename: "testfile",
			Url:      "/api/v1/testfile",
		},
		},
	}
	filenamesBytes, err := easyjson.Marshal(filenames)
	if err != nil {
		log.Warning(err)
	}

	attachUC.EXPECT().ListAttach(context.Background(), &attach_proto.GetAttachRequest{
		Username: signinForm.Username,
		MailID:   1,
		Filename: "",
	}).Return(&attach_proto.AttachListResponse{
		Filenames: filenamesBytes,
	}, nil)

	attachUC.EXPECT().ListAttach(context.Background(), &attach_proto.GetAttachRequest{
		Username: signinForm.Username,
		MailID:   2,
		Filename: "",
	}).Return(nil, errors.New("no such mailID"))

	//&models.Session{Username: "test", Authenticated: true}, int32(0)
	err = SigninUser(client, signinForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	formBytes := []byte(`{"mail_id":1}`)

	r, err := Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}

	var resp models.AttachList

	err = json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		t.Error(err)
		return
	}

	formBytesFail := []byte(`{"mail_id":2}`)

	_, err = Post(client, formBytesFail, url, http.StatusInternalServerError, token, "")
	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetAttach(t *testing.T) {
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
	router := InitTestRouter(&d, []string{"/mail/attach/get", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetAttach, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/mail/attach/get", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := easyjson.Marshal(signinForm)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	attachUC.EXPECT().GetAttach(context.Background(), &attach_proto.GetAttachRequest{
		Username: signinForm.Username,
		MailID:   1,
		Filename: "dummy.png",
	}).Return(&attach_proto.AttachResponse{
		File: []byte{10, 10, 10},
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

	formBytes := []byte(`
		{
			"mail_id":1,
			"attach_id":"dummy.png"
		}`)

	_, err = Post(client, formBytes, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}

	formBytes = []byte(`
		{
			mail_id:1,
			"attach_id":"dummy.png"
		}`)

	_, err = Post(client, formBytes, url, http.StatusInternalServerError, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

//func TestUploadAttach(t *testing.T) {
//	mockCtrl := gomock.NewController(t)
//	defer mockCtrl.Finish()
//
//	authUC := auth_proto.NewMockAuthClient(mockCtrl)
//	folderManagerUC := folder_manager_proto.NewMockFolderManagerClient(mockCtrl)
//	mailboxUC := mailbox_proto.NewMockMailboxClient(mockCtrl)
//	profileUC := profile_proto.NewMockProfileClient(mockCtrl)
//	attachUC := attach_proto.NewMockAttachClient(mockCtrl)
//
//	jar, _ := cookiejar.New(nil)
//
//	client := &http.Client{
//		Jar: jar,
//	}
//
//	d := delivery.Delivery{}
//	router := InitTestRouter(&d, []string{"/mail/attach/add", "/signin"}, []func(http.ResponseWriter, *http.Request){d.UploadAttach, d.SignIn},
//		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
//	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC, attachUC)
//
//	srv := httptest.NewServer(router)
//	defer srv.Close()
//
//	url := fmt.Sprintf("%s/mail/attach/add", srv.URL)
//
//	signinForm := models.SignInForm{
//		Username: "test",
//		Password: "test",
//	}
//	signinFormBytes, _ := easyjson.Marshal(signinForm)
//
//	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
//		Form: signinFormBytes,
//	}).Return(&utils_proto.JsonResponse{
//		Response: pkg.NO_ERR.Bytes(),
//	}, nil)
//
//	//file := []byte{10, 10, 10}
//	file, err := os.ReadFile("../../static/dummy.png")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	attachUC.EXPECT().SaveAttach(context.Background(), &attach_proto.SaveAttachRequest{
//		Username: signinForm.Username,
//		MailID:   1,
//		File:     file,
//	}).Return(&attach_proto.Nothing{
//		Status: true,
//	}, nil)
//
//	//&models.Session{Username: "test", Authenticated: true}, int32(0)
//	err = SigninUser(client, signinForm, srv.URL)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	_, err = writer.CreateFormFile("attach", "dummy.png")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	body.Write(file)
//	writer.Close()
//
//	//formBytes := []byte(`{"mail_id":1}`)
//
//	_, err = Post(client, body.Bytes(), url, http.StatusOK, token, "")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	//
//	//var resp models.AttachList
//	//
//	//err = json.NewDecoder(r.Body).Decode(&resp)
//	//
//	//if err != nil {
//	//	t.Error(err)
//	//	return
//	//}
//
//	//body := &bytes.Buffer{}
//	//writer := multipart.NewWriter(body)
//
//	//mediaFiles := []string{"dummy.png"}
//	metadata := `{"title": "New title", "description": "New description"}`
//	// Metadata part.
//	metadataHeader := textproto.MIMEHeader{}
//	metadataHeader.Set("Content-Type", "application/json")
//	metadataHeader.Set("Content-ID", "metadata")
//	part, _ := writer.CreatePart(metadataHeader)
//	part.Write([]byte(metadata))
//
//	mediaHeader := textproto.MIMEHeader{}
//	mediaHeader.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\".", "dummy.png"))
//	mediaHeader.Set("Content-ID", "media")
//	mediaHeader.Set("Content-Filename", "dummy.png")
//
//	mediaPart, _ := writer.CreatePart(mediaHeader)
//	io.Copy(mediaPart, bytes.NewReader(file))
//
//	// Close multipart writer.
//	writer.Close()
//
//	// Request Content-Type with boundary parameter.
//	//contentType := fmt.Sprintf("multipart/related; boundary=%s", writer.Boundary())
//
//}
