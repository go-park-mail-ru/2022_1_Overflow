package delivery_test

//
//import (
//	"OverflowBackend/internal/delivery"
//	"OverflowBackend/internal/models"
//	"OverflowBackend/pkg"
//	"OverflowBackend/proto/attach_proto"
//	"OverflowBackend/proto/auth_proto"
//	"OverflowBackend/proto/folder_manager_proto"
//	"OverflowBackend/proto/mailbox_proto"
//	"OverflowBackend/proto/profile_proto"
//	"OverflowBackend/proto/utils_proto"
//	"bytes"
//	"context"
//	"github.com/golang/mock/gomock"
//	"github.com/mailru/easyjson"
//	"mime/multipart"
//	"net/http"
//	"net/http/cookiejar"
//	"net/http/httptest"
//	"testing"
//)
//
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
//	err := SigninUser(client, signinForm, srv.URL)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	attach := models.Attach{
//		Filename:    "testfile",
//		PayloadSize: 3,
//		Payload:     []byte{10, 10, 10},
//	}
//	attachBytes, _ := easyjson.Marshal(attach)
//
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	_, err = writer.CreateFormFile("file", "testfile")
//	_, err = writer.CreateFormField("MailID")
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	body.Write(attachBytes)
//	writer.Close()
//
//	attachUC.EXPECT().SaveAttach(context.Background(), &attach_proto.SaveAttachRequest{
//		Username: "test",
//		MailID:   1,
//		File:     []byte{10, 10, 10},
//	}).Return(nil, nil)
//	//
//	////dataJson, _ := easyjson.Marshal(data)
//	//_, err, token := Get(client, fmt.Sprintf("%s/attach/add", srv.URL), http.StatusMethodNotAllowed)
//	//if err != nil {
//	//	t.Error(err)
//	//	return
//	//}
//	//
//	//_, err = Post(client, body.Bytes(), fmt.Sprintf("%s/attach/add", srv.URL), http.StatusOK, token, "")
//	//if err != nil {
//	//	t.Error(err)
//	//	return
//	//}
//}
