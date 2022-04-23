package delivery_test

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/delivery"
	"OverflowBackend/internal/middlewares"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase"
	"OverflowBackend/internal/usecase/session"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var DefConf = config.TestConfig()

func init() {
	log.SetLevel(log.FatalLevel)
}

func InitTestRouter(uc usecase.UseCaseInterface, d *delivery.Delivery, urls []string, handles []func(http.ResponseWriter, *http.Request)) http.Handler {
	session.Init(DefConf)
	middlewares.Init(DefConf)
	
	d.Init(uc, DefConf)
	router := mux.NewRouter()
	for i := range urls {
		router.HandleFunc(urls[i], handles[i])
	}
	return middlewares.Middleware(router)
}

func SigninUser(client *http.Client, form models.SignInForm, srv_url string) error {
	dataJson, err := json.Marshal(form)
	if err != nil {
		return err
	}
	_, err, token := Get(client, fmt.Sprintf("%s/signin", srv_url), http.StatusMethodNotAllowed)
	if err != nil {
		return err
	}
	r, err := Post(client, dataJson, fmt.Sprintf("%s/signin", srv_url), http.StatusOK, token, "")
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Неверный статус от сервера: %v. Ожидается: %v.", r.StatusCode, http.StatusOK)
	}
	return nil
}

func Post(client *http.Client, data []byte, reqUrl string, expectedHttpStatus int, token string, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CSRF-Token", token)
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

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