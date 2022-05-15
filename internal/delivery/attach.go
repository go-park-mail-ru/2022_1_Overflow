package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"bytes"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	"io"
	"net/http"
	"strconv"
)

// UploadAttach godoc
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
func (d *Delivery) UploadAttach(w http.ResponseWriter, r *http.Request) {
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
	file, header, err := r.FormFile("attach")
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)
	attach := models.Attach{
		Filename:    header.Filename,
		PayloadSize: header.Size,
		Payload:     buf.Bytes(),
	}
	attachBytes, _ := json.Marshal(attach)
	_, err = d.attach.SaveAttach(context.Background(), &attach_proto.SaveAttachRequest{
		Username: data.Username,
		MessID:   "0",
		File:     attachBytes,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /profile/avatar/set [get]
// @Tags profile
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func UploadAttach() {}

// GetAttach godoc
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
func (d *Delivery) GetAttach(w http.ResponseWriter, r *http.Request) {
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

	var form models.GetAttachForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}

	grpcResp, err := d.attach.GetAttach(context.Background(), &attach_proto.GetAttachRequest{
		Username: data.Username,
		MessID:   form.MailID,
		AttachID: form.AttachID,
	})
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}

	file := bytes.NewReader(grpcResp.File)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+form.AttachID+"")
	w.Header().Set("Content-Length", strconv.Itoa(file.Len()))
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Can`t download photo!", http.StatusInternalServerError)
		log.Error(err)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// @Router /profile/avatar/set [get]
// @Tags profile
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func GetAttach() {}
