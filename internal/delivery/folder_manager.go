package delivery

import (
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/folder_manager_proto"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// AddFolder godoc
// @Summary Добавить папку с письмами для пользователя
// @Produce json
// @Param folder_name query string true "Имя папки."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/add [post]
// @Param X-CSRF-Token header string true "CSRF токен"
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
	folderName := r.URL.Query().Get("folder_name")
	if len(folderName) == 0 {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.folderManager.AddFolder(context.Background(), &folder_manager_proto.AddFolderRequest{
		Data: data,
		Name: folderName,
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

// @Router /folder/add [get]
// @Response 200 {object} pkg.JsonResponse
// @Header 200 {string} X-CSRF-Token "CSRF токен"
func AddFolder() {}

// AddMailToFolder godoc
// @Summary Добавить письмо в папку с письмами
// @Produce json
// @Param folder_id query int true "ID папки."
// @Param mail_id query int true "ID добавляемого письма."
// @Param move query bool true "Следует ли переместить письмо в эту папку (с последующим удалением из источника)."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/mail/add [post]
// @Param X-CSRF-Token header string true "CSRF токен"
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
	folderId, err := strconv.Atoi(r.URL.Query().Get("folder_id"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	mailId, err := strconv.Atoi(r.URL.Query().Get("mail_id"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	move, err  := strconv.ParseBool(r.URL.Query().Get("move"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.folderManager.AddMailToFolder(context.Background(), &folder_manager_proto.AddMailToFolderRequest{
		Data:       data,
		FolderId: int32(folderId),
		MailId:     int32(mailId),
		Move: move,
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
func AddMailToFolder() {}

// ChangeFolder godoc
// @Summary Переименовать папку с письмами
// @Produce json
// @Param folder_id query int true "ID изменяемой папки."
// @Param new_folder_name query string true "Новое имя папки."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/rename [post]
// @Param X-CSRF-Token header string true "CSRF токен"
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
	folderId, err := strconv.Atoi(r.URL.Query().Get("folder_id"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	folderNewName := r.URL.Query().Get("new_folder_name")
	if len(folderNewName) == 0 {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.folderManager.ChangeFolder(context.Background(), &folder_manager_proto.ChangeFolderRequest{
		Data:          data,
		FolderId:    int32(folderId),
		FolderNewName: folderNewName,
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
func ChangeFolder() {}

// DeleteFolder godoc
// @Summary Удалить папку с письмами
// @Produce json
// @Param folder_name query string true "Имя папки."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
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
	folderName := r.URL.Query().Get("folder_name")
	if len(folderName) == 0 {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.folderManager.DeleteFolder(context.Background(), &folder_manager_proto.DeleteFolderRequest{
		Data: data,
		Name: folderName,
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
func DeleteFolder() {}

// DeleteFolderMail godoc
// @Summary Удалить письмо из папки
// @Produce json
// @Param folder_id query int true "ID папки"
// @Param mail_id query int true "ID удаляемого письма"
// @Param restore query bool true "Восстановить письмо (добавить обратно во входящие)."
// @Success 200 {object} pkg.JsonResponse "OK"
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/mail/delete [post]
// @Param X-CSRF-Token header string true "CSRF токен"
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
	folderId, err := strconv.Atoi(r.URL.Query().Get("folder_id"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	mailId, err := strconv.Atoi(r.URL.Query().Get("mail_id"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	restore, err  := strconv.ParseBool(r.URL.Query().Get("restore"))
	if err != nil {
		pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
		return
	}
	resp, err := d.folderManager.DeleteFolderMail(context.Background(), &folder_manager_proto.DeleteFolderMailRequest{
		Data: data,
		FolderId: int32(folderId),
		MailId: int32(mailId),
		Restore: restore,
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
func DeleteFolderMail() {}

// ListFolders godoc
// @Summary Получить список папок пользователя или список писем в определенной папке
// @Produce json
// @Param folder_id query int false "ID папки с письмами"
// @Success 200 {object} []models.Folder "Список папок."
// @Success 200 {object} []models.MailAdditional "Список писем в папке."
// @Failure 401 {object} pkg.JsonResponse "Сессия отсутствует или сессия не валидна."
// @Failure 405 {object} pkg.JsonResponse
// @Failure 500 {object} pkg.JsonResponse "Ошибка БД, неверные GET параметры."
// @Router /folder/list [get]
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
	folderIdStr := r.URL.Query().Get("folder_id")
	if len(folderIdStr) > 0 {
		folderId, err := strconv.Atoi(folderIdStr)
		if err != nil {
			pkg.WriteJsonErrFull(w, &pkg.GET_ERR)
			return
		}
		resp, err := d.folderManager.ListFolder(context.Background(), &folder_manager_proto.ListFolderRequest{
			Data: data,
			FolderId: int32(folderId),
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
