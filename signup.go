package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserT struct {
	LastName  string
	FirstName string
	Email     string
	Password  string
}

type SignupHandler struct {
	validKeys []string
}

// Инициализация обработчика регистрации. Обязательна к вызову.
func (handler *SignupHandler) Init() {
	handler.validKeys = []string {"first_name", "last_name", "email", "password", "password_confirmation"}
}

func (handler *SignupHandler) Handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handler.UserSignup)
	return mux
}

// Основная функция-обработчик запроса регистрации.
func (handler *SignupHandler) UserSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is available.", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
    if err != nil {
        w.Write(handler.createJsonResponse(1, err.Error(), nil))
		return
    }
	if err := handler.validateData(r); err != nil {
		w.Write(handler.createJsonResponse(2, err.Error(), nil))
		return
	}
	user, err := handler.convertToUser(r)
	if err != nil {
		w.Write(handler.createJsonResponse(3, err.Error(), nil))
		return
	}
	if err := writeToDatabase(user); err != nil {
		w.Write(handler.createJsonResponse(4, err.Error(), nil))
		return
	}
	w.Write(handler.createJsonResponse(0, "OK", nil))
}

/*
Создать ответ на запрос в формате JSON, где:
	status - статус ответа;
	message - доп. сообщение ответа;
	content - передаваемое содержимое ответа;

	Статусы:
		0 - запрос прошел успешно,
		1 - невозможно обработать данные формы в запросе,
		2 - данные формы не прошли валидацию,
		3 - ошибка при преобразовании данных формы в структуру БД,
		4 - ошибка при записи в БД.
*/
func (handler *SignupHandler) createJsonResponse(status int, message string, content interface{}) []byte {
	resp, _ := json.Marshal(
		map[string]interface{}{
			"status":  status,
			"message": message,
			"content": content,
		},
	)
	return resp
}

func (handler *SignupHandler) validateData(r *http.Request) (err error) {
	for _, key := range handler.validKeys {
		val := r.FormValue(key)
		if len(strings.TrimSpace(val)) == 0 {
			return fmt.Errorf("Поле %v не может быть пустым.", key)
		}
	}
	if r.FormValue("password") != r.FormValue("password_confirmation") {
		return fmt.Errorf("Пароли не совпадают.")
	}
	return
}

func (handler *SignupHandler) convertToUser(r *http.Request) (user UserT, err error) {
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	user.Password = hashPassword(r.FormValue("password"))
	return
}

func writeToDatabase(user UserT) (err error) {
	//log.Println("Writing to database..")
	return
}

func hashPassword(passw string) string {
	//log.Println("Hashing password..")
	return passw
}