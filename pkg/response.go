package pkg

import (
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
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
	STATUS_NOT_IMP        = iota // 12
	STATUS_WRONG_CREDS    = iota // 13
	STATUS_LOGGED_IN      = iota // 14
	STATUS_OBJECT_EXISTS  = iota // 15
	STATUS_BAD_FILETYPE   = iota // 16
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (s *JsonResponse) Bytes() []byte {
	bytes, _ := easyjson.Marshal(s)
	return bytes
}

var (
	NO_ERR              = JsonResponse{Status: STATUS_OK, Message: "OK"}
	SESSION_ERR         = JsonResponse{Status: STATUS_BAD_SESSION, Message: "Ошибка получения сессии."}
	INTERNAL_ERR        = JsonResponse{Status: STATUS_INTERNAL, Message: "Внутренняя ошибка сервера."}
	JSON_ERR            = JsonResponse{Status: STATUS_ERR_JSON, Message: "Ошибка конвертации JSON."}
	DB_ERR              = JsonResponse{Status: STATUS_ERR_DB, Message: "Ошибка базы данных."}
	BAD_METHOD_ERR      = JsonResponse{Status: STATUS_BAD_METHOD, Message: "Запрещенный HTTP метод."}
	UNAUTHORIZED_ERR    = JsonResponse{Status: STATUS_UNAUTHORIZED, Message: "Отказано в доступе."}
	GET_ERR             = JsonResponse{Status: STATUS_BAD_GET, Message: "Неверный GET запрос."}
	NOT_IMPLEMENTED_ERR = JsonResponse{Status: STATUS_NOT_IMP, Message: "Не имплементировано."}
	WRONG_CREDS_ERR     = JsonResponse{Status: STATUS_WRONG_CREDS, Message: "Неверная пара логин/пароль."}
	LOGGED_IN_ERR       = JsonResponse{Status: STATUS_LOGGED_IN, Message: "Пользователь уже выполнил вход."}
	NO_USER_EXIST       = JsonResponse{Status: STATUS_NO_USER, Message: "Пользователя не существует."}
	BAD_FILETYPE        = JsonResponse{Status: STATUS_BAD_FILETYPE, Message: "Неподдерживаемый тип файла."}
)

func WriteJsonErrFull(w http.ResponseWriter, err *JsonResponse) {
	switch err.Status {
	case STATUS_OK:
		w.WriteHeader(http.StatusOK)
	case STATUS_UNAUTHORIZED:
		w.WriteHeader(http.StatusUnauthorized)
	case STATUS_BAD_METHOD:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case STATUS_BAD_VALIDATION:
		w.WriteHeader(http.StatusBadRequest)
	case STATUS_WRONG_CREDS:
		w.WriteHeader(http.StatusBadRequest)
	case STATUS_NO_USER:
		w.WriteHeader(http.StatusNotFound)
	case STATUS_BAD_FILETYPE:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	resp, _ := easyjson.Marshal(err)
	_, er := w.Write(resp)
	if er != nil {
		log.Warning("Resp: ", er)
	}
}

func WriteJsonErr(w http.ResponseWriter, status int, message string) {
	err := JsonResponse{
		Status:  status,
		Message: message,
	}
	WriteJsonErrFull(w, &err)
}

func CreateJsonErr(status int, message string) *JsonResponse {
	err := JsonResponse{
		Status:  status,
		Message: message,
	}
	return &err
}
