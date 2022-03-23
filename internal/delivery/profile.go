package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/pkg"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// GetInfo godoc
// @Summary Получение данных пользователя
// @Produce json
// @Success 200 {object} models.User "Информация о пользователе"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует, сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, пользователь не найден, неверные данные сессии."
// @Router /profile [get]
func (d *Delivery) GetInfo(w http.ResponseWriter, r *http.Request) {
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
	userJson, err := d.uc.GetInfo(data)
	if err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	w.Write(userJson)
}

// SetInfo godoc
// @Summary Изменение настроек пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное изменение настроек."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Accept json
// @Param SettingsForm body models.SettingsForm true "Форма настроек пользователя."
// @Produce json
// @Router /profile/set [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		pkg.WriteJsonErrFull(w, pkg.BAD_METHOD_ERR)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.SESSION_ERR)
		return
	}

	var form models.SettingsForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, pkg.JSON_ERR)
		return
	}

	if err := d.uc.SetInfo(data, &form); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// SetAvatar godoc
// @Summary Установка/смена аватарки пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное установка аватарки."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Accept multipart/form-data
// @Param file formData file true "Файл аватарки."
// @Produce json
// @Router /profile/avatar/set [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SetAvatar(w http.ResponseWriter, r *http.Request) {
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

	var buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		pkg.WriteJsonErrFull(w, pkg.INTERNAL_ERR)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)
	avatar := models.Avatar{
		Name:      header.Filename,
		UserEmail: data.Email,
		Content:   buf.Bytes(),
	}
	if err := d.uc.SetAvatar(data, &avatar); err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}
	pkg.WriteJsonErrFull(w, pkg.NO_ERR)
}

// GetAvatar godoc
// @Summary Получение ссылки на аватарку пользователя
// @Success 200 {object} pkg.JsonResponse "Ссылка на аватарку в формате /{static_dir}/{file}.{ext}."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, пользователь не найден или сессия не валидна."
// @Produce json
// @Router /profile/avatar [get]
func (d *Delivery) GetAvatar(w http.ResponseWriter, r *http.Request) {
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

	url, err := d.uc.GetAvatar(data)
	if err != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, err)
		return
	}

	pkg.WriteJsonErr(w, pkg.STATUS_OK, url)
}
