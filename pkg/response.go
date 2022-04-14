package pkg

import (
	"encoding/json"
	"net/http"
)

const (
	STATUS_OK             = iota // 0
	STATUS_BAD_SESSION    = iota // 1
	STATUS_BAD_VALIDATION = iota // 2
	STATUS_INTERNAL       = iota // 3
	STATUS_ERR_JSON       = iota // 4
	STATUS_ERR_DB         = iota // 5
	STATUS_BAD_METHOD     = iota // 6
	STATUS_UNAUTHORIZED   = iota // 7
	STATUS_BAD_GET        = iota // 8
	STATUS_UNKNOWN        = iota // 9
	STATUS_USER_EXISTS    = iota // 10
	STATUS_NO_USER        = iota // 11
	STATUS_NOT_IMP 		  = iota // 12
	STATUS_WRONG_CREDS 	  = iota // 13
	STATUS_LOGGED_IN      = iota // 14
)

type JsonResponse struct {
	Status  int		`json:"status"`
	Message string	`json:"message"`
}

var (
	NO_ERR           = JsonResponse{STATUS_OK, "OK"}
	SESSION_ERR      = JsonResponse{STATUS_BAD_SESSION, "Ошибка получения сессии."}
	INTERNAL_ERR     = JsonResponse{STATUS_INTERNAL, "Внутренняя ошибка сервера."}
	JSON_ERR         = JsonResponse{STATUS_ERR_JSON, "Ошибка конвертации JSON."}
	DB_ERR           = JsonResponse{STATUS_ERR_DB, "Ошибка базы данных."}
	BAD_METHOD_ERR   = JsonResponse{STATUS_BAD_METHOD, "Запрещенный HTTP метод."}
	UNAUTHORIZED_ERR = JsonResponse{STATUS_UNAUTHORIZED, "Отказано в доступе."}
	GET_ERR          = JsonResponse{STATUS_BAD_GET, "Неверный GET запрос."}
	NOT_IMPLEMENTED_ERR = JsonResponse{STATUS_NOT_IMP, "Не имплементировано."}
	WRONG_CREDS_ERR = JsonResponse{STATUS_WRONG_CREDS, "Неверная пара логин/пароль."}
	LOGGED_IN_ERR = JsonResponse{STATUS_LOGGED_IN, "Пользователь уже выполнил вход."}
	NO_USER_EXIST = JsonResponse{STATUS_NO_USER, "Пользователя не существует."}
)

func WriteJsonErrFull(w http.ResponseWriter, err JsonResponse) {
	switch err.Status {
		case STATUS_OK: w.WriteHeader(http.StatusOK)
		case STATUS_UNAUTHORIZED: w.WriteHeader(http.StatusUnauthorized)
		case STATUS_BAD_METHOD: w.WriteHeader(http.StatusMethodNotAllowed)
		case STATUS_BAD_VALIDATION: w.WriteHeader(http.StatusBadRequest)
		case STATUS_WRONG_CREDS: w.WriteHeader(http.StatusBadRequest)
		case STATUS_NO_USER: w.WriteHeader(http.StatusNotFound)
		default: w.WriteHeader(http.StatusInternalServerError)
	}
	resp, _ := json.Marshal(err)
	w.Write(resp)
}

func WriteJsonErr(w http.ResponseWriter, status int, message string) {
	err := JsonResponse{
		status,
		message,
	}
	WriteJsonErrFull(w, err)
}

func CreateJsonErr(status int, message string) JsonResponse {
	err := JsonResponse{
		status,
		message,
	}
	return err
}
