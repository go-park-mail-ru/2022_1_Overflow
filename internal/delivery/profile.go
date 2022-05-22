package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/security/xss"
	"OverflowBackend/internal/session"
	"OverflowBackend/internal/validation"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/profile_proto"
	"bytes"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
)

// GetInfo godoc
// @Summary Получение данных пользователя
// @Produce json
// @Success 200 {object} models.ProfileInfo "Информация о пользователе"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует, сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, пользователь не найден, неверные данные сессии."
// @Router /profile [get]
// @Tags profile
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
	if response != pkg.NO_ERR {
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
// @Tags profile
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
	form.Firstname = xss.EscapeInput(form.Firstname)
	form.Lastname = xss.EscapeInput(form.Lastname)
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
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
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /profile/set [get]
// @Tags profile
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func SetInfo() {}

// ChangePassword godoc
// @Summary Изменение пароля пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное изменение пароля."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Accept json
// @Param SettingsForm body models.ChangePasswordForm true "Форма изменение пароля."
// @Produce json
// @Router /profile/change_password [post]
// @Tags profile
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) ChangePassword(w http.ResponseWriter, r *http.Request) {
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
	var form models.ChangePasswordForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if !validation.SameFields(form.NewPassword, form.NewPasswordConf) {
		pkg.WriteJsonErrFull(w, pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, "Новый пароль и его подтверждение не совпадают."))
		return
	}
	if validation.SameFields(form.NewPassword, form.OldPassword) {
		pkg.WriteJsonErrFull(w, pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, "Старый и новый пароли совпадают."))
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErrFull(w, pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()))
		return
	}
	resp, err := d.profile.ChangePassword(context.Background(), &profile_proto.ChangePasswordRequest{
		Data:        data,
		PasswordOld: form.OldPassword,
		PasswordNew: form.NewPassword,
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

// @Router /profile/change_password [get]
// @Tags profile
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func ChangePassword() {}

const (
	MB = 1 << 20
)

// SetAvatar godoc
// @Summary Установка/смена аватарки пользователя
// @Success 200 {object} pkg.JsonResponse "Успешное установка аватарки."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Accept multipart/form-data
// @Param file formData file true "Файл аватарки."
// @Produce json
// @Router /profile/avatar/set [post]
// @Tags profile
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

	if err := r.ParseMultipartForm(10 * MB); err != nil {
		log.Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.BAD_FILETYPE)
		return
	}

	// Limit upload size
	r.Body = http.MaxBytesReader(w, r.Body, 10*MB) // 10 Mb

	file, multipartFileHeader, err := r.FormFile("file")
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.BAD_FILETYPE)
		return
	}
	defer file.Close()

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		log.Warning(err)
		return
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		log.Warning(err)
		return
	}

	//log.Printf("Name: %#v\n", multipartFileHeader.Filename)
	//log.Printf("Size: %#v\n", file.(Sizer).Size())
	log.Printf("MIME: %#v\n", http.DetectContentType(fileHeader))
	mime := http.DetectContentType(fileHeader)

	if mime != "image/png" && mime != "image/jpeg" && mime != "image/webp" {
		log.Warning("file is not jpeg or png")
		pkg.WriteJsonErrFull(w, &pkg.BAD_FILETYPE)
		return
	}

	//file, header, err := r.FormFile("file")
	//if err != nil {
	//	pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
	//	return
	//}
	//defer file.Close()

	//img, format, err := image.Decode(file)
	//if err != nil {
	//	log.Warning("can't decode file: ", err)
	//	pkg.WriteJsonErrFull(w, &pkg.BAD_FILETYPE)
	//	return
	//}

	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	avatar := models.Avatar{
		Name:     multipartFileHeader.Filename,
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
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /profile/avatar/set [get]
// @Tags profile
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
// @Tags profile
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
	if response != pkg.NO_ERR {
		pkg.WriteJsonErrFull(w, &response)
		return
	}
	pkg.WriteJsonErr(w, pkg.STATUS_OK, resp.Url)
}
