package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/session"
	ws "OverflowBackend/internal/websocket"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/mailbox_proto"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/validator.v2"
)

// Income godoc
// @Summary Получение входящих сообщений
// @Produce json
// @Param limit query int false "Ограничение на количество писем"
// @Param offset query int false "Смещение в списке писем"
// @Success 200 {object} models.MailAddList "Список входящих писем"
// @Failure 401 {object} pkg.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД."
// @Router /mail/income [get]
// @Tags mailbox
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
	limit, e := strconv.Atoi(r.URL.Query().Get("limit"))
	if e != nil || limit > 100 {
		limit = 100
	}
	offset, e := strconv.Atoi(r.URL.Query().Get("offset"))
	if e != nil {
		offset = 0
	}
	resp, err := d.mailbox.Income(context.Background(), &mailbox_proto.IncomeRequest{
		Data:   data,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	w.Write(resp.Mails)
}

// Outcome godoc
// @Summary Получение исходящих сообщений
// @Produce json
// @Param limit query int false "Ограничение на количество писем"
// @Param offset query int false "Смещение в списке писем"
// @Success 200 {object} models.MailAddList "Список исходящих писем"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД."
// @Router /mail/outcome [get]
// @Tags mailbox
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
	limit, e := strconv.Atoi(r.URL.Query().Get("limit"))
	if e != nil || limit > 100 {
		limit = 100
	}
	offset, e := strconv.Atoi(r.URL.Query().Get("offset"))
	if e != nil {
		offset = 0
	}
	resp, err := d.mailbox.Outcome(context.Background(), &mailbox_proto.OutcomeRequest{
		Data:   data,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	w.Write(resp.Mails)
}

// GetMail godoc
// @Summary Получение сообщения по его id
// @Produce json
// @Param id query int true "ID запрашиваемого письма."
// @Success 200 {object} models.Mail "Объект письма."
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД."
// @Router /mail/get [get]
// @Tags mailbox
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
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Mail)
}

// DeleteMail godoc
// @Summary Удалить письмо по его id
// @Produce json
// @Param DeleteMailForm body models.DeleteMailForm true "Форма запроса"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags mailbox
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
	var form models.DeleteMailForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.mailbox.DeleteMail(context.Background(), &mailbox_proto.DeleteMailRequest{
		Data: data,
		Id:   form.Id,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /mail/delete [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags mailbox
func DeleteMail() {}

// ReadMail godoc
// @Summary Отметить число прочитанным/непрочитанным по его id. При отсутствии параметра isread запрос отмечает письмо с заданным id прочитанным.
// @Produce json
// @Param ReadMailForm body models.ReadMailForm true "Форма запроса"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/read [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags mailbox
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
	var form models.ReadMailForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.mailbox.ReadMail(context.Background(), &mailbox_proto.ReadMailRequest{
		Data: data,
		Id:   form.Id,
		Read: form.IsRead,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /mail/read [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags mailbox
func ReadMail() {}

// SendMail godoc
// @Summary Выполняет отправку письма получателю
// @Success 200 {object} pkg.JsonResponse "Успешная отправка письма."
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 500 {object} pkg.JsonResponse "Получатель не существует, ошибка БД."
// @Accept json
// @Param MailForm body models.MailForm true "Форма письма"
// @Produce json
// @Router /mail/send [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags mailbox
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
	var form models.MailForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	form.Addressee = xss.EscapeInput(form.Addressee)
	form.Files = xss.EscapeInput(form.Files)
	form.Text = xss.EscapeInput(form.Text)
	form.Theme = xss.EscapeInput(form.Theme)
	formBytes, _ := json.Marshal(form)
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.mailbox.SendMail(context.Background(), &mailbox_proto.SendMailRequest{
		Data: data,
		Form: formBytes,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response, &response)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}

	d.ws <- ws.WSMessage{
		Type:          ws.TYPE_ALERT,
		Username:      form.Addressee,
		Message:       "Новое письмо!",
		MessageStatus: ws.STATUS_INFO,
	}

	response.Message = resp.Param
	pkg.WriteJsonErrFull(w, &response)
}

// @Router /mail/send [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags mailbox
func SendMail() {}
