package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"encoding/json"
	"net/http"
	"strconv"
)

// Income godoc
// @Summary Получение входящих сообщений
// @Produce json
// @Success 200 {object} []models.Mail "Список входящих писем"
// @Failure 401 {object} pkg.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД."
// @Router /mail/income [get]
func (d *Delivery) Income(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	data, e := session.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}
	parsed, err := d.uc.Income(data)
	if err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	w.Write(parsed)
}

// Outcome godoc
// @Summary Получение исходящих сообщений
// @Produce json
// @Success 200 {object} []models.Mail "Список исходящих писем"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД."
// @Router /mail/outcome [get]
func (d *Delivery) Outcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	data, e := session.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}
	parsed, err := d.uc.Outcome(data)
	if err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	w.Write(parsed)
}

// DeleteMail godoc
// @Summary Удалить письмо по его id
// @Produce json
// @Param id query string true "ID запрашиваемого письма."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) DeleteMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}
	data, err := session.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.GET_ERR)
		return
	}
	if err := d.uc.DeleteMail(data, int32(id)); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, pkg.INTERNAL_ERR)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// ReadMail godoc
// @Summary Прочитать письмо по его id
// @Produce json
// @Param id query string true "ID запрашиваемого письма."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/read [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) ReadMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}
	data, err := session.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.GET_ERR)
		return
	}
	if err := d.uc.ReadMail(data, int32(id)); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

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
func (d *Delivery) SendMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}

	var form models.MailForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, pkg.JSON_ERR)
		return
	}

	form.Addressee = xss.P.Sanitize(form.Addressee)
	form.Files = xss.P.Sanitize(form.Files)
	form.Text = xss.P.Sanitize(form.Text)
	form.Theme = xss.P.Sanitize(form.Theme)

	if err := d.uc.SendMail(data, form); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// ForwardMail godoc
// @Summary Переслать уже существующее письмо
// @Produce json
// @Param mail_id query int true "ID запрашиваемого письма."
// @Param username query string true "Почта получателя."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/forward [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) ForwardMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}
	data, err := session.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}
	idStr := r.URL.Query().Get("mail_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.GET_ERR)
		return
	}
	username := r.URL.Query().Get("username")
	if len(username) == 0 {
		pkg.WriteJsonErrFull(w, pkg.GET_ERR)
		return
	}
	e := d.uc.ForwardMail(data, int32(id), username)
	if e != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, e)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}


// RespondMail godoc
// @Summary Ответить на письмо пользователя
// @Produce json
// @Param MailResponse body models.MailResponse true "Форма ответа на письмо"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse"Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры."
// @Router /mail/respond [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) RespondMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	_, err := session.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}

	var form models.MailResponse

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, pkg.JSON_ERR)
		return
	}

	//d.uc.RespondMail(form)
}
