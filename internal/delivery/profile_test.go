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
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetInfo(t *testing.T) {
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
	router := InitTestRouter(&d, []string{"/profile", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetInfo, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	url := fmt.Sprintf("%s/profile", srv.URL)

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := json.Marshal(signinForm)

	info, _ := json.Marshal(models.User{
		Id:        0,
		Firstname: "test",
		Lastname:  "test",
		Password:  "test",
		Username:  "test",
	})

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	//&models.Session{Username: "test", Authenticated: true}
	getInfoData := &utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}
	profileUC.EXPECT().GetInfo(context.Background(), &profile_proto.GetInfoRequest{
		Data: getInfoData,
	}).Return(&profile_proto.GetInfoResponse{
		Data: info,
		Response: &utils_proto.JsonResponse{
			Response: pkg.NO_ERR.Bytes(),
		},
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

func TestSetInfo(t *testing.T) {
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
	router := InitTestRouter(&d, []string{"/profile/set", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SetInfo, d.SignIn},
		authUC, profileUC, mailboxUC, folderManagerUC)
	d.Init(DefConf, authUC, profileUC, mailboxUC, folderManagerUC)

	srv := httptest.NewServer(router)
	defer srv.Close()

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}
	signinFormBytes, _ := json.Marshal(signinForm)

	url := fmt.Sprintf("%s/profile/set", srv.URL)

	data := models.ProfileSettingsForm{
		Firstname: "changed",
		Lastname:  "changed",
	}
	dataBytes, _ := json.Marshal(data)

	authUC.EXPECT().SignIn(context.Background(), &auth_proto.SignInRequest{
		Form: signinFormBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	//&models.Session{Username: "test", Authenticated: true}, &data
	setInfoData := &utils_proto.Session{
		Username:      "test",
		Authenticated: wrapperspb.Bool(true),
	}
	profileUC.EXPECT().SetInfo(context.Background(), &profile_proto.SetInfoRequest{
		Data: setInfoData,
		Form: dataBytes,
	}).Return(&utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil)

	err := SigninUser(client, signinForm, srv.URL)

	if err != nil {
		t.Error(err)
		return
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, url, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Post(client, dataJson, url, http.StatusOK, token, "")
	if err != nil {
		t.Error(err)
		return
	}
}

/*
func TestGetAvatar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/profile/avatar", "/signin"}, []func(http.ResponseWriter, *http.Request){d.GetAvatar, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	url := fmt.Sprintf("%s/profile/avatar", srv.URL)
	expAvatarUrl := "/static/dummy.png"

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().GetAvatar("test").Return(expAvatarUrl, pkg.NO_ERR)

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
	resp := utils_proto.JsonResponse{}
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		t.Error(err)
		return
	}

	if resp.Status != pkg.STATUS_OK {
		t.Errorf("Неверный статус JSON ответа. Получено: %v, ожидалось: %v.", resp.Status, pkg.STATUS_OK)
		return
	}

	if resp.Message != expAvatarUrl {
		t.Errorf("Неверная ссылка на аватар пользователя. Получено: %v, ожидалось: %v.", resp.Message, expAvatarUrl)
		return
	}
}

func TestSetAvatar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUC := mocks.NewMockUseCaseInterface(mockCtrl)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	d := delivery.Delivery{}
	router := InitTestRouter(mockUC, &d, []string{"/profile/avatar/set", "/signin"}, []func(http.ResponseWriter, *http.Request){d.SetAvatar, d.SignIn})

	srv := httptest.NewServer(router)
	defer srv.Close()

	signinForm := models.SignInForm{
		Username: "test",
		Password: "test",
	}

	reqUrl := fmt.Sprintf("%s/profile/avatar/set", srv.URL)

	avatar := models.Avatar{
		Name:      "avatar",
		UserEmail: signinForm.Username,
		Content:   []byte{10, 10, 10, 10},
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err := writer.CreateFormFile("file", avatar.Name)

	if err != nil {
		t.Error(err)
		return
	}

	body.Write(avatar.Content)
	writer.Close()

	mockUC.EXPECT().SignIn(signinForm).Return(pkg.NO_ERR)
	mockUC.EXPECT().SetAvatar(&models.Session{Username: "test", Authenticated: true}, &avatar).Return(pkg.NO_ERR)

	err = SigninUser(client, signinForm, srv.URL)
	if err != nil {
		t.Error(err)
		return
	}

	_, err, token := Get(client, reqUrl, http.StatusMethodNotAllowed)
	if err != nil {
		t.Error(err)
		return
	}

	r, err := Post(client, body.Bytes(), reqUrl, http.StatusOK, token, writer.FormDataContentType())
	if err != nil {
		t.Error(err)
		return
	}

	resp := utils_proto.JsonResponse{}
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		t.Error(err)
		return
	}

	if resp.Status != pkg.STATUS_OK {
		t.Errorf("Неверный статус JSON ответа. Получено: %v, ожидалось: %v.", resp.Status, pkg.STATUS_OK)
		return
	}
}
*/
