package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/profile_proto"
	"bytes"
	"context"
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
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	resp, err := d.profile.GetInfo(context.Background(), &profile_proto.GetInfoRequest{
		Data: data,
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
	if (response != pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	w.Write(resp.Data)
}

// SetInfo godoc
// @Summary Изменение настроек пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное изменение настроек."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Accept json
// @Param SettingsForm body models.ProfileSettingsForm true "Форма настроек пользователя."
// @Produce json
// @Router /profile/set [post]
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) SetInfo(w http.ResponseWriter, r *http.Request) {
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

	var form models.ProfileSettingsForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}

	form.Firstname = xss.P.Sanitize(form.Firstname)
	form.Lastname = xss.P.Sanitize(form.Lastname)
	form.Password = xss.P.Sanitize(form.Password)
	formBytes, _ := json.Marshal(form)

	resp, err := d.profile.SetInfo(context.Background(), &profile_proto.SetInfoRequest{
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
	if (response != pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /profile/set [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SetInfo() {}

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
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	data, err := session.Manager.GetData(r)
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}

	var buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)
	avatar := models.Avatar{
		Name:     header.Filename,
		Username: data.Username,
		File:     buf.Bytes(),
	}
	avatarBytes, _ := json.Marshal(avatar)
	resp, err := d.profile.SetAvatar(context.Background(), &profile_proto.SetAvatarRequest{
		Data:   data,
		Avatar: avatarBytes,
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
	if (response != pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /profile/avatar/set [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SetAvatar() {}

// GetAvatar godoc
// @Summary Получение ссылки на аватарку пользователя
// @Description Получение ссылки на аватарку текущего пользователя или пользователя с конкретным логином (username).
// @Param username query string false "Имя пользователя, соответствующее аватарке."
// @Success 200 {object} pkg.JsonResponse "Ссылка на аватарку в формате /{static_dir}/{file}.{ext}."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, пользователь не найден или сессия не валидна."
// @Produce json
// @Router /profile/avatar [get]
func (d *Delivery) GetAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}

	username := r.URL.Query().Get("username")
	if len(username) == 0 {
		data, e := session.Manager.GetData(r)
		if e != nil {
			pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
			return
		}
		username = data.Username
	}

	resp, err := d.profile.GetAvatar(context.Background(), &profile_proto.GetAvatarRequest{
		Username: username,
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
	if (response != pkg.NO_ERR) {
		pkg.WriteJsonErrFull(w, &response)
		return
	}

	pkg.WriteJsonErr(w, pkg.STATUS_OK, resp.Url)
}
