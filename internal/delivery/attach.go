package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"bytes"
	"context"
	"encoding/json"
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	"io"
	"net/http"
	"strconv"
)

// UploadAttach godoc
// @Summary Добавление вложения в письмо
// @Success 200 {object} pkg.JsonResponse "Успешное добавление вложения."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Accept multipart/form-data
// @Param attach formData file true "Файл вложения."
// @Produce json
// @Router /mail/attach/add [post]
// @Tags mailbox
// @Param mailID formData string true "MailID"
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
	mailID, err := strconv.Atoi(r.FormValue("mailID"))
	if err != nil {
		log.Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	var buf bytes.Buffer
	file, header, err := r.FormFile("attach")
	if err != nil {
		log.Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	defer file.Close()

	if header.Size > (20 << 20) {
		log.Warning("Размер файла превышает 20мб.")
		pkg.WriteJsonErrFull(w, &pkg.BAD_FILETYPE)
		return
	}

	if _, err := io.Copy(&buf, file); err != nil {
		log.Warning(err)
	}
	attach := models.Attach{
		Filename:    header.Filename,
		PayloadSize: header.Size,
		Payload:     buf.Bytes(),
	}
	attachBytes, _ := easyjson.Marshal(attach)
	_, err = d.attach.SaveAttach(context.Background(), &attach_proto.SaveAttachRequest{
		Username: data.Username,
		MailID:   int32(mailID),
		File:     attachBytes,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"Username": data.Username,
			"MailID":   mailID,
		}).Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}
	pkg.WriteJsonErrFull(w, &pkg.NO_ERR)
}

// UploadAttach
// @Router /mail/attach/add [get]
// @Tags mailbox
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func UploadAttach() {}

// GetAttach godoc
// @Summary Получение вложения по filename и mailID
// @Success 200 {object} pkg.JsonResponse "Успешная оттдача файла."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
// @Produce json
// @Router /mail/attach/get [post]
// @Tags mailbox
// @Accept json
// @Param GetAttachForm body models.GetAttachForm true "Форма получения вложения"
// @Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) GetAttach(w http.ResponseWriter, r *http.Request) {
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
		MailID:   form.MailID,
		Filename: form.AttachID,
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

// GetAttach
// @Router /mail/attach/get [get]
// @Tags mailbox
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func GetAttach() {}

// ListAttach godoc
//@Summary Получение списка вложений письма
//@Success 200 {object} pkg.JsonResponse "Успешное установка аватарки."
//@Failure 405 {object} pkg.JsonResponse
//@Failure 500 {object} pkg.JsonResponse "Ошибка валидации формы, БД или сессия не валидна."
//@Accept json
//@Param GetListAttachForm body models.GetListAttachForm true "Форма получения списка вложений."
//@Produce json
//@Router /mail/attach/list [post]
//@Tags mailbox
//@Param X-CSRF-Token header string true "CSRF токен"
func (d *Delivery) ListAttach(w http.ResponseWriter, r *http.Request) {
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

	var form models.GetListAttachForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}

	grpcResp, err := d.attach.ListAttach(context.Background(), &attach_proto.GetAttachRequest{
		Username: data.Username,
		MailID:   form.MailID,
		Filename: "",
	})
	if err != nil {
		log.WithFields(log.Fields{
			"Username": data.Username,
			"MailID":   form.MailID,
		}).Warning(err)
		pkg.WriteJsonErrFull(w, &pkg.INTERNAL_ERR)
		return
	}

	var attaches models.AttachList
	if err := easyjson.Unmarshal(grpcResp.Filenames, &attaches); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}

	if _, err := w.Write(grpcResp.Filenames); err != nil {
		log.Warning()
	}
}

// ListAttach
// @Router /mail/attach/list [get]
// @Tags mailbox
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func ListAttach() {}
