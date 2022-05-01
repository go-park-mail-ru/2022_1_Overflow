package delivery

import (
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

// Income godoc
// @Summary Получение входящих сообщений
// @Produce json
// @Success 200 {object} []utils_proto.MailAdditional "Список входящих писем"
// @Failure 401 {object} utils_proto.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} utils_proto.JsonResponse
// @Failure 500 {object} utils_proto.JsonResponse "Ошибка БД."
// @Router /mail/income [get]
func (d *Delivery) Income(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	resp, err := d.mailbox.Income(context.Background(), data)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	if !proto.Equal(resp.Response, &pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, resp.Response)
		return
	}
	w.Write(resp.Mails)
}

// Outcome godoc
// @Summary Получение исходящих сообщений
// @Produce json
// @Success 200 {object} []utils_proto.MailAdditional "Список исходящих писем"
// @Failure 401 {object} utils_proto.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} utils_proto.JsonResponse
// @Failure 500 {object} utils_proto.JsonResponse "Ошибка БД."
// @Router /mail/outcome [get]
func (d *Delivery) Outcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	resp, err := d.mailbox.Outcome(context.Background(), data)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	if !proto.Equal(resp.Response, &pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, resp.Response)
		return
	}
	w.Write(resp.Mails)
}

// GetMail godoc
// @Summary Получение сообщения по его id
// @Produce json
// @Param id query int true "ID запрашиваемого письма."
// @Success 200 {object} utils_proto.Mail "Объект письма."
// @Failure 401 {object} utils_proto.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} utils_proto.JsonResponse
// @Failure 500 {object} utils_proto.JsonResponse "Ошибка БД."
// @Router /mail/get [get]
func (d *Delivery) GetMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}

	mail_id, e := strconv.Atoi(r.URL.Query().Get("id"))
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.mailbox.GetMail(context.Background(), &mailbox_proto.GetMailRequest{
		Data: data,
		Id:   int32(mail_id),
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	if !proto.Equal(resp.Response, &pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, resp.Response)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Mail)
}

// DeleteMail godoc
// @Summary Удалить письмо по его id
// @Produce json
// @Param id query int true "ID запрашиваемого письма."
// @Success 200 {object} utils_proto.JsonResponse "OK"
// @Failure 401 {object} utils_proto.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} utils_proto.JsonResponse
// @Failure 500 {object} utils_proto.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) DeleteMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	data, err := session.Manager.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.mailbox.DeleteMail(context.Background(), &mailbox_proto.DeleteMailRequest{
		Data: data,
		Id:   int32(id),
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	if !proto.Equal(resp, &pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, resp)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /mail/delete [get]
// @Response 200 {object} utils_proto.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func DeleteMail() {}

// ReadMail godoc
// @Summary Прочитать письмо по его id
// @Produce json
// @Param id query string true "ID запрашиваемого письма."
// @Success 200 {object} utils_proto.JsonResponse "OK"
// @Failure 401 {object} utils_proto.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} utils_proto.JsonResponse
// @Failure 500 {object} utils_proto.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/read [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) ReadMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	data, err := session.Manager.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.mailbox.ReadMail(context.Background(), &mailbox_proto.ReadMailRequest{
		Data: data,
		Id:   int32(id),
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	if !proto.Equal(resp, &pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, resp)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /mail/read [get]
// @Response 200 {object} utils_proto.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func ReadMail() {}

// SendMail godoc
// @Summary Выполняет отправку письма получателю
// @Success 200 {object} utils_proto.JsonResponse "Успешная отправка письма."
// @Failure 401 {object} utils_proto.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 500 {object} utils_proto.JsonResponse "Получатель не существует, ошибка БД."
// @Accept json
// @Param MailForm body utils_proto.MailForm true "Форма письма"
// @Produce json
// @Router /mail/send [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SendMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	data, err := session.Manager.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}

	var form utils_proto.MailForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}

	form.Addressee = xss.P.Sanitize(form.Addressee)
	form.Files = xss.P.Sanitize(form.Files)
	form.Text = xss.P.Sanitize(form.Text)
	form.Theme = xss.P.Sanitize(form.Theme)

	resp, err := d.mailbox.SendMail(context.Background(), &mailbox_proto.SendMailRequest{
		Data: data,
		Form: &form,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	if !proto.Equal(resp, &pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, resp)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /mail/send [get]
// @Response 200 {object} utils_proto.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SendMail() {}
