package delivery_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/internal/usecase/session"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var DefConf = config.TestConfig()

func init() {
	log.SetLevel(log.FatalLevel)
}

func InitTestRouter(db repository.DatabaseRepository, d *delivery.Delivery, urls []string, handles []func(http.ResponseWriter, *http.Request)) http.Handler {
	session.Init(DefConf)
	uc := usecase.UseCase{}
	uc.Init(db, DefConf)
	d.Init(&uc, DefConf)
	router := mux.NewRouter()
	for i := range urls {
		router.HandleFunc(urls[i], handles[i])
	}
	return middlewares.Middleware(router)
}

func createTestUsers(db repository.DatabaseRepository) {
	db.AddUser(models.User{
		Id:        0,
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	})
	db.AddUser(models.User{
		Id:        1,
		FirstName: "test2",
		LastName:  "test2",
		Username:  "test2",
		Password:  "test2",
	})
}

func PrepareMails(db repository.DatabaseRepository, num int) {
	for i := 0; i < num; i++ {
		mail := models.Mail{
			Id: int32(i),
			Client_id: 0,
			Sender: "test",
			Addressee: "test2",
			Theme: "test",
			Files: "test",
			Date: time.Now(),
			Read: false,
		}
		db.AddMail(mail)
	}
}

func SigninUser(client *http.Client, form models.SignInForm, srv_url string) error {
	dataJson, err := json.Marshal(form)
	if err != nil {
		return err
	}
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv_url), http.StatusOK)
	if err != nil {
		return err
	}
	r, err := Post(client, dataJson, fmt.Sprintf("%s/signin", srv_url), http.StatusOK, token)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusOK)
	}
	return nil
}

func Post(client *http.Client, data []byte, reqUrl string, expectedHttpStatus int, token string) (*http.Response, error) {
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CSRF-Token", token)

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != expectedHttpStatus {
		return nil, fmt.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, expectedHttpStatus)
	}

	return r, nil
}

func Get(client *http.Client, reqUrl string, expectedHttpStatus int) (r *http.Response, err error, token string) {
	r, err = client.Get(reqUrl)
	if err != nil {
		return
	}

	if r.StatusCode != expectedHttpStatus {
		err = fmt.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, expectedHttpStatus)
		return
	}

	token = r.Header.Get("X-CSRF-Token")
	return
}