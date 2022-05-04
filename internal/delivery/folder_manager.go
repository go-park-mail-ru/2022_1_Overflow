package delivery

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/folder_manager_proto"
	"context"
	"encoding/json"
	"net/http"

	"gopkg.in/validator.v2"
)

// AddFolder godoc
// @Summary Добавить папку с письмами для пользователя
// @Produce json
// @Param AddFolderForm body models.AddFolderForm true "Форма запроса"
// @Success 200 {object} models.Folder "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/add [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags folder_manager
func (d *Delivery) AddFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	var form models.AddFolderForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.folderManager.AddFolder(context.Background(), &folder_manager_proto.AddFolderRequest{
		Data: data,
		Name: form.FolderName,
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
	w.Write(resp.Folder)
}

// @Router /folder/add [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags folder_manager
func AddFolder() {}

// AddMailToFolder godoc
// @Summary Добавить письмо в папку с письмами
// @Produce json
// @Param AddMailToFolderForm body models.AddMailToFolderForm true "Форма запроса"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/mail/add [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags folder_manager
func (d *Delivery) AddMailToFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	var form models.AddMailToFolderForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.folderManager.AddMailToFolder(context.Background(), &folder_manager_proto.AddMailToFolderRequest{
		Data:       data,
		FolderName: form.FolderName,
		MailId:     form.MailId,
		Move: form.Move,
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

// @Router /folder/mail/add [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags folder_manager
func AddMailToFolder() {}

// ChangeFolder godoc
// @Summary Переименовать папку с письмами
// @Produce json
// @Param ChangeFolderForm body models.ChangeFolderForm true "Форма запроса"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/rename [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags folder_manager
func (d *Delivery) ChangeFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	var form models.ChangeFolderForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.folderManager.ChangeFolder(context.Background(), &folder_manager_proto.ChangeFolderRequest{
		Data:          data,
		FolderName:    form.FolderName,
		FolderNewName: form.NewFolderName,
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

// @Router /folder/rename [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags folder_manager
func ChangeFolder() {}

// DeleteFolder godoc
// @Summary Удалить папку с письмами
// @Produce json
// @Param DeleteFolderForm body models.DeleteFolderForm true "Форма запроса"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags folder_manager
func (d *Delivery) DeleteFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	var form models.DeleteFolderForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.folderManager.DeleteFolder(context.Background(), &folder_manager_proto.DeleteFolderRequest{
		Data: data,
		FolderName: form.FolderName,
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

// @Router /folder/delete [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags folder_manager
func DeleteFolder() {}

// DeleteFolderMail godoc
// @Summary Удалить письмо из папки
// @Produce json
// @Param DeleteFolderMailForm body models.DeleteFolderMailForm true "Форма запроса"
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/mail/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
// @Tags folder_manager
func (d *Delivery) DeleteFolderMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		pkg.WriteJsonErrFull(w, &pkg.BAD_METHOD_ERR)
		return
	}
	data, e := session.Manager.GetData(r)
	if e != nil {
		pkg.WriteJsonErrFull(w, &pkg.SESSION_ERR)
		return
	}
	var form models.DeleteFolderMailForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		pkg.WriteJsonErrFull(w, &pkg.JSON_ERR)
		return
	}
	if err := validator.Validate(form); err != nil {
		pkg.WriteJsonErr(w, pkg.STATUS_BAD_VALIDATION, err.Error())
		return
	}
	resp, err := d.folderManager.DeleteFolderMail(context.Background(), &folder_manager_proto.DeleteFolderMailRequest{
		Data: data,
		FolderName: form.FolderName,
		MailId: form.MailId,
		Restore: form.Restore,
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

// @Router /folder/mail/delete [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
// @Tags folder_manager
func DeleteFolderMail() {}

// ListFolders godoc
// @Summary Получить список папок пользователя или список писем в определенной папке
// @Produce json
// @Param folder_name query string false "Имя папки с письмами"
// @Success 200 {object} []models.Folder "Список папок."
// @Success 200 {object} []models.MailAdditional "Список писем в папке."
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/list [get]
// @Tags folder_manager
func (d *Delivery) ListFolders(w http.ResponseWriter, r *http.Request) {
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
	folderName := r.URL.Query().Get("folder_name")
	if len(folderName) > 0 {
		resp, err := d.folderManager.ListFolder(context.Background(), &folder_manager_proto.ListFolderRequest{
			Data: data,
			FolderName: folderName,
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
		w.Write(resp.Mails)
	} else {
		resp, err := d.folderManager.ListFolders(context.Background(), &folder_manager_proto.ListFoldersRequest{
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
		w.Write(resp.Folders)
	}
}
