package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

// Находится ли письмо в какой либо папке
func (c *Database) IsMailInAnyFolder(context context.Context, mailId int32) bool {
	var counter int
	err := c.Conn.QueryRow(context, "SELECT COUNT(*) FROM overflow.folder_to_mail WHERE mail_id=$1", mailId).Scan(&counter)
	return err == nil && counter > 0
}

// Является ли письмо перемещенным в какую либо папку
func (c *Database) IsMailMoved(context context.Context, mailId int32) bool {
	var counter int
	err := c.Conn.QueryRow(context, "SELECT COUNT(*) FROM overflow.folder_to_mail WHERE mail_id=$1 AND only_folder=true", mailId).Scan(&counter)
	return err == nil && counter > 0
}

func (c *Database) GetFolderById(context context.Context, request *repository_proto.GetFolderByIdRequest) (response *repository_proto.ResponseFolder, err error) {
	var folder models.Folder
	folderBytes, _ := json.Marshal(folder)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseFolder{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Folder: folderBytes,
			}
		}
	}()
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

func (c *Database) GetFolderByName(context context.Context, request *repository_proto.GetFolderByNameRequest) (response *repository_proto.ResponseFolder, err error) {
	var folder models.Folder
	folderBytes, _ := json.Marshal(folder)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseFolder{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Folder: folderBytes,
			}
		}
	}()
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

func (c *Database) GetFoldersByUser(context context.Context, request *repository_proto.GetFoldersByUserRequest) (response *repository_proto.ResponseFolders, err error) {
	var folders models.FolderList
	foldersBytes, _ := json.Marshal(folders)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseFolders{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Folders: foldersBytes,
			}
		}
	}()
	var count int
	err = c.Conn.QueryRow(context, "SELECT COUNT(*) FROM overflow.folders WHERE user_id=$1 AND name NOT IN ($2, $3);", request.UserId, pkg.FOLDER_SPAM, pkg.FOLDER_DRAFTS).Scan(&count)
	if err != nil {
		return &repository_proto.ResponseFolders{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folders: foldersBytes,
		}, err
	}
	folders.Amount = count
	rows, err := c.Conn.Query(context, "SELECT id, name, user_id, created_at FROM overflow.folders WHERE user_id=$1 AND name NOT IN ($2, $3) ORDER BY created_at DESC OFFSET $5 LIMIT $4;", request.UserId, pkg.FOLDER_SPAM, pkg.FOLDER_DRAFTS, request.Limit, request.Offset)
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
		folders.Folders = append(folders.Folders, folder)
	}
	foldersBytes, _ = json.Marshal(folders)
	return &repository_proto.ResponseFolders{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folders: foldersBytes,
	}, nil
}

func (c *Database) GetFolderMail(context context.Context, request *repository_proto.GetFolderMailRequest) (response *repository_proto.ResponseMails, err error) {
	var mails models.MailList
	mailsBytes, _ := json.Marshal(mails)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: mailsBytes,
			}
		}
	}()
	var count int
	err = c.Conn.QueryRow(context, "SELECT COUNT(*) FROM overflow.mails WHERE id IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id in (SELECT id FROM overflow.folders WHERE user_id=$1 AND name=$2))", request.UserId, request.FolderName).Scan(&count)
	if err != nil {
		return &repository_proto.ResponseMails{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Mails: mailsBytes,
		}, err
	}
	mails.Amount = count
	rows, err := c.Conn.Query(context, "SELECT addressee, sender, theme, text, files, date, read, id FROM overflow.mails WHERE id IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id in (SELECT id FROM overflow.folders WHERE user_id=$1 AND name=$2)) ORDER BY date DESC OFFSET $4 LIMIT $3;", request.UserId, request.FolderName, request.Limit, request.Offset)
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
		mails.Mails = append(mails.Mails, mail)
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

// Добавить письмо в папку
func (c *Database) AddMailToFolderById(context context.Context, request *repository_proto.AddMailToFolderByIdRequest) (*utils_proto.DatabaseResponse, error) {
	if c.IsMailMoved(context, request.MailId) {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, nil
	}
	rows, err := c.Conn.Query(context, "INSERT INTO overflow.folder_to_mail(folder_id, mail_id, only_folder) SELECT id, $3, $4 FROM overflow.folders WHERE user_id=$1 AND name=$2;", request.UserId, request.FolderName, request.MailId, request.Move)
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

func (c *Database) AddMailToFolderByObject(context context.Context, request *repository_proto.AddMailToFolderByObjectRequest) (*utils_proto.DatabaseResponse, error) {
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	var mailId int32 
	err = c.Conn.QueryRow(context, "INSERT INTO overflow.mails(addressee, date, files, sender, text, theme) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;", mail.Addressee, mail.Date, mail.Files, mail.Sender, mail.Text, mail.Theme).Scan(&mailId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "INSERT INTO overflow.folder_to_mail(folder_id, mail_id, only_folder) SELECT id, $3, true FROM overflow.folders WHERE user_id=$1 AND name=$2;", request.UserId, request.FolderName, mailId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

// Удалить письмо из папки
func (c *Database) DeleteFolderMail(context context.Context, request *repository_proto.DeleteFolderMailRequest) (*utils_proto.DatabaseResponse, error) {
	_, err := c.Conn.Exec(context, "DELETE FROM overflow.folder_to_mail WHERE folder_id IN (SELECT id FROM overflow.folders WHERE user_id=$1) AND mail_id=$2;", request.UserId, request.MailId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}

func (c *Database) MoveFolderMail(context context.Context, request *repository_proto.MoveFolderMailRequest) (*utils_proto.DatabaseResponse, error) {
	var folderIdSrc int32
	var folderIdDest int32
	err := c.Conn.QueryRow(context, "SELECT id FROM overflow.folders WHERE user_id=$1 AND name=$2;", request.UserId, request.FolderNameSrc).Scan(&folderIdSrc)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	err = c.Conn.QueryRow(context, "SELECT id FROM overflow.folders WHERE user_id=$1 AND name=$2;", request.UserId, request.FolderNameDest).Scan(&folderIdDest)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.folder_to_mail SET folder_id=$2 WHERE folder_id=$1 AND mail_id=$3", folderIdSrc, folderIdDest, request.MailId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	return &utils_proto.DatabaseResponse{
		Status: utils_proto.DatabaseStatus_OK,
	}, nil
}