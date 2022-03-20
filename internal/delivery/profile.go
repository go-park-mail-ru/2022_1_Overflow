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
// @Summary Получение данных профиля пользователя
// @Produce json
// @Success 200 {object} models.SettingsForm "Информация о пользователе"
// @Failure 401 "Сессия отсутствует, сессия не валидна."
// @Failure 500 "Ошибка БД, пользователь не найден, неверные данные сессии."
// @Router /profile [get]
func (d *Delivery) GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w, http.MethodGet)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJson, err := d.uc.GetInfo(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

// SetInfo godoc
// @Summary Изменение настроек пользователя
// @Success 200 "Успешное изменение настроек."
// @Failure 500 "Ошибка валидации формы, БД или сессия не валидна."
// @Accept json
// @Param Avatar body models.SettingsForm true "Форма настроек пользователя."
// @Produce plain
// @Router /set_profile [post]
func (d *Delivery) SetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

	var form models.SettingsForm

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data.Email != form.User.Email {
		http.Error(w, "Запрещено изменять почту пользователя.", http.StatusInternalServerError)
		return
	}
	
	if err := d.uc.SetInfo(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// SetAvatar godoc
// @Summary Установка/смена аватарки пользователя
// @Success 200 "Успешное установка аватарки."
// @Failure 500 "Ошибка валидации формы, БД или сессия не валидна."
// @Accept multipart/form-data
// @Param file formData file true "Файл аватарки."
// @Produce plain
// @Router /set_profile/avatar [post]
func (d *Delivery) SetAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w, http.MethodPost)
		return
	}

	data, err := session.GetData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//r.ParseMultipartForm(32 << 20) // проверить макс. длину
	var buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)
	avatar := models.Avatar{
		Name: header.Filename,
		UserEmail: data.Email,
		Content: buf.Bytes(),
	}
	if err := d.uc.SetAvatar(data, &avatar); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}