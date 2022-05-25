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
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/mailru/easyjson"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

func TestWSconnect(t *testing.T) {
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
	router := InitTestRouter(&d, []string{"/ws", "/signin"}, []func(http.ResponseWriter, *http.Request){d.WSConnect, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC, attachUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/ws", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := easyjson.Marshal(signinForm)

	//info, _ := easyjson.Marshal(models.User{
	//	Id:        0,
	//	Firstname: "test",
	//	Lastname:  "test",
	//	Password:  "test",
	//	Username:  "test",
	//})

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	_, err, _ := Get(client, url, http.StatusUnauthorized)
	if err != nil {
		t.Error(err)
		return
	}

	err = SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	_, err, _ = Get(client, url, http.StatusBadRequest)
	if err != nil {
		t.Error(err)
		return
	}

}
