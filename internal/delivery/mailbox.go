package delivery

import (
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"net/http"
	"strconv"
)

// Income godoc
// @Summary Получение входящих сообщений
// @Produce json
// @Success 200 {object} []models.Mail "Список входящих писем"
// @Failure 401 "Сессия отсутствует или сессия не валидна."
// @Failure 405
// @Failure 500 "Ошибка БД."
// @Router /mail/income [get]
func (d *Delivery) Income(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parsed, err := d.uc.Income(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}

// Outcome godoc
// @Summary Получение исходящих сообщений
// @Produce json
// @Success 200 {object} []models.Mail "Список исходящих писем"
// @Failure 401 "Сессия отсутствует или сессия не валидна."
// @Failure 405
// @Failure 500 "Ошибка БД."
// @Router /mail/outcome [get]
func (d *Delivery) Outcome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parsed, err := d.uc.Outcome(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}

// DeleteMail godoc
// @Summary Удалить письмо по его id
// @Produce json
// @Param id query string true "ID запрашиваемого письма."
// @Success 200 "OK"
// @Failure 401 "Письмо не принадлежит пользователю, сессия отсутствует или сессия не валидна."
// @Failure 405
// @Failure 500 "Ошибка БД, неверные GET параметры."
// @Router /mail/delete [get]
func (d *Delivery) DeleteMail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}
	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = d.uc.DeleteMail(data, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ReadMail godoc
// @Summary Прочитать письмо по его id
// @Produce json
// @Param id query string true "ID запрашиваемого письма."
// @Success 200 "OK"
// @Failure 401 "Письмо не принадлежит пользователю, сессия отсутствует или сессия не валидна."
// @Failure 405
// @Failure 500 "Ошибка БД, неверные GET параметры."
// @Router /mail/read [get]
func (d *Delivery) ReadMail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}
	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = d.uc.ReadMail(data, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
