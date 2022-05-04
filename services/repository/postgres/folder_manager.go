package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"time"
	//log "github.com/sirupsen/logrus"
)

func (c *Database) GetFolderById(context context.Context, request *repository_proto.GetFolderByIdRequest) (*repository_proto.ResponseFolder, error) {
	var folder models.Folder
	folderBytes, _ := json.Marshal(folder)
	rows, err := c.Conn.Query(context, "SELECT id, name, user_id, created_at FROM overflow.folders WHERE id = $1;", request.FolderId)
	if err != nil {
		return &repository_proto.ResponseFolder{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folder: folderBytes,
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseFolder{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Folder: folderBytes,
			}, err
		}
		folder.Id = values[0].(int32)
		folder.Name = values[1].(string)
		folder.UserId = values[2].(int32)
		folder.CreatedAt = values[3].(time.Time)
	}
	folderBytes, _ = json.Marshal(folder)
	return &repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil
}

func (c *Database) GetFolderByName(context context.Context, request *repository_proto.GetFolderByNameRequest) (*repository_proto.ResponseFolder, error) {
	var folder models.Folder
	folderBytes, _ := json.Marshal(folder)
	rows, err := c.Conn.Query(context, "SELECT id, name, user_id, created_at FROM overflow.folders WHERE name = $1 AND user_id = $2;", request.FolderName, request.UserId)
	if err != nil {
		return &repository_proto.ResponseFolder{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folder: folderBytes,
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseFolder{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Folder: folderBytes,
			}, err
		}
		folder.Id = values[0].(int32)
		folder.Name = values[1].(string)
		folder.UserId = values[2].(int32)
		folder.CreatedAt = values[3].(time.Time)
	}
	folderBytes, _ = json.Marshal(folder)
	return &repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: folderBytes,
	}, nil
}

func (c *Database) GetFoldersByUser(context context.Context, request *repository_proto.GetFoldersByUserRequest) (*repository_proto.ResponseFolders, error) {
	var folders []models.Folder
	foldersBytes, _ := json.Marshal(folders)
	rows, err := c.Conn.Query(context, "SELECT id, name, user_id, created_at FROM overflow.folders WHERE user_id=$1 ORDER BY created_at DESC;", request.UserId)
	if err != nil {
		return &repository_proto.ResponseFolders{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folders: foldersBytes,
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseFolders{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Folders: foldersBytes,
			}, err
		}
		var folder models.Folder
		folder.Id = values[0].(int32)
		folder.Name = values[1].(string)
		folder.UserId = values[2].(int32)
		folder.CreatedAt = values[3].(time.Time)
		folders = append(folders, folder)
	}
	foldersBytes, _ = json.Marshal(folders)
	return &repository_proto.ResponseFolders{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folders: foldersBytes,
	}, nil
}

func (c *Database) GetFolderMail(context context.Context, request *repository_proto.GetFolderMailRequest) (*repository_proto.ResponseMails, error) {
	var mails []models.Mail
	mailsBytes, _ := json.Marshal(mails)
	rows, err := c.Conn.Query(context, "SELECT addressee, sender, theme, text, files, date, read, id FROM overflow.mails WHERE id IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id in (SELECT id FROM overflow.folders WHERE user_id=$1 AND name=$2)) ORDER BY date DESC;", request.UserId, request.FolderName)
	if err != nil {
		return &repository_proto.ResponseMails{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Mails: mailsBytes,
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail models.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: mailsBytes,
			}, err
		}
		mail.Addressee = values[0].(string)
		mail.Sender = values[1].(string)
		mail.Theme = values[2].(string)
		mail.Text = values[3].(string)
		mail.Files = values[4].(string)
		mail.Date = values[5].(time.Time)
		mail.Read = values[6].(bool)
		mail.Id = values[7].(int32)
		mails = append(mails, mail)
	}
	mailsBytes, _ = json.Marshal(mails)
	return &repository_proto.ResponseMails{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mails: mailsBytes,
	}, nil
}

func (c *Database) DeleteFolder(context context.Context, request *repository_proto.DeleteFolderRequest) (*utils_proto.DatabaseResponse, error) {
	rows, err := c.Conn.Query(context, "DELETE FROM overflow.folders WHERE user_id=$1 AND name=$2;", request.UserId, request.FolderName)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	defer rows.Close()
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (c *Database) AddFolder(context context.Context, request *repository_proto.AddFolderRequest) (*utils_proto.DatabaseResponse, error) {
	rows, err := c.Conn.Query(context, "INSERT INTO overflow.folders(name, user_id) VALUES ($1, $2);", request.Name, request.UserId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	defer rows.Close()
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (c *Database) ChangeFolderName(context context.Context, request *repository_proto.ChangeFolderNameRequest) (*utils_proto.DatabaseResponse, error) {
	rows, err := c.Conn.Query(context, "UPDATE overflow.folders SET name=$1 WHERE user_id=$2 AND name=$3;", request.NewName, request.UserId, request.FolderName)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	defer rows.Close()
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (c *Database) AddMailToFolder(context context.Context, request *repository_proto.AddMailToFolderRequest) (*utils_proto.DatabaseResponse, error) {
	if request.Move {
		rowsUpdate, err := c.Conn.Query(context, "UPDATE overflow.mails SET only_folder=$1 WHERE id=$2;", request.Move, request.MailId)
		if err != nil {
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		defer rowsUpdate.Close()
	}
	rows, err := c.Conn.Query(context, "INSERT INTO overflow.folder_to_mail(folder_id, mail_id) SELECT id, $3 FROM overflow.folders WHERE user_id=$1 AND name=$2;", request.UserId, request.FolderName, request.MailId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	defer rows.Close()
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (c *Database) DeleteFolderMail(context context.Context, request *repository_proto.DeleteFolderMailRequest) (*utils_proto.DatabaseResponse, error) {
	if request.Restore {
		rowsRestore, err := c.Conn.Query(context, "UPDATE overflow.mails SET only_folder=$1 WHERE id=$2;", !request.Restore, request.MailId)
		if err != nil {
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		rowsRestore.Close()
		rows, err := c.Conn.Query(context, "DELETE FROM overflow.folder_to_mail WHERE folder_id IN (SELECT id FROM overflow.folders WHERE user_id=$1) AND mail_id=$2;", request.UserId, request.MailId)
		if err != nil {
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		rows.Close()
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, nil
	} else {
		resp, err := c.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
			MailId: request.MailId,
		})
		if err != nil {
			return resp.Response, err
		}
		return c.DeleteMail(context, &repository_proto.DeleteMailRequest{
			Mail: resp.Mail,
			UserId: request.UserId,
		})
	}
}