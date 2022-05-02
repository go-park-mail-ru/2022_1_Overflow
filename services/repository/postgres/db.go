package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	//log "github.com/sirupsen/logrus"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
}

type Database struct {
	url  string
	Conn PgxIface
}

func (c *Database) Create(url string) (err error) {
	c.url = url
	c.Conn, err = pgxpool.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return nil
}

// Получить данные пользователя по его почте
func (c *Database) GetUserInfoByUsername(context context.Context, request *repository_proto.GetUserInfoByUsernameRequest) (*repository_proto.ResponseUser, error) {
	var user models.User
	userBytes, _ := json.Marshal(user)
	rows, err := c.Conn.Query(context, "Select * from overflow.users where username = $1;", request.Username)
	if err != nil {
		return &repository_proto.ResponseUser{
			User: userBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseUser{
				User: userBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		user.Id = values[0].(int32)
		user.Firstname = values[1].(string)
		user.Lastname = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	userBytes, _ = json.Marshal(user)
	return &repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Получить данные пользователя по его айди в бд
func (c *Database) GetUserInfoById(context context.Context, request *repository_proto.GetUserInfoByIdRequest) (*repository_proto.ResponseUser, error) {
	var user models.User
	userBytes, _ := json.Marshal(user)
	rows, err := c.Conn.Query(context, "Select * from overflow.users(id, first_name, last_name, password, username) where id = $1;", request.UserId)
	if err != nil {
		return &repository_proto.ResponseUser{
			User: userBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseUser{
				User: userBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		user.Id = values[0].(int32)
		user.Firstname = values[1].(string)
		user.Lastname = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	userBytes, _ = json.Marshal(user)
	return &repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

func (c *Database) UserConfig(context context.Context, user_id int32) error {
	rows, err := c.Conn.Query(context, "INSERT INTO overflow.folders(name, user_id) VALUES ($1, $2);", pkg.FOLDER_SPAM, user_id)
	if err != nil {
		return err
	}
	rows.Close()
	rows, err = c.Conn.Query(context, "INSERT INTO overflow.folders(name, user_id) VALUES ($1, $2);", pkg.FOLDER_DRAFTS, user_id)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

// Добавить пользователя
func (c *Database) AddUser(context context.Context, request *repository_proto.AddUserRequest) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res, err := c.Conn.Query(context, "INSERT INTO overflow.users(first_name, last_name, password, username) VALUES ($1, $2, $3, $4);", user.Firstname, user.Lastname, user.Password, user.Username)
	if err == nil {
		res.Close()
		resp, err := c.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
			Username: user.Username,
		})
		if err != nil {
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		err = json.Unmarshal(resp.User, &user)
		if err != nil{
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		c.UserConfig(context, user.Id) // конфигурация профиля пользователя
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Изменить пароль
func (c *Database) ChangeUserPassword(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.users SET password = $1 WHERE id = $2;", request.Data, user.Id)
	if err == nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Изменить имя
func (c *Database) ChangeUserFirstName(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.users SET first_name = $1 WHERE id = $2;", request.Data, user.Id)
	if err == nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Изменить фамилию
func (c *Database) ChangeUserLastName(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.users SET last_name = $1 WHERE id = $2;", request.Data, user.Id)
	if err == nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Добавить письмо
func (c *Database) AddMail(context context.Context, request *repository_proto.AddMailRequest) (*utils_proto.DatabaseResponse, error) {
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res, err := c.Conn.Query(context, "INSERT INTO overflow.mails(client_id, sender, addressee, theme, text, files, date) VALUES ($1, $2, $3, $4, $5, $6, $7);", mail.ClientId, mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date)
	if err == nil {
		res.Close()
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

//Удалить письмо
func (c *Database) DeleteMail(context context.Context, request *repository_proto.DeleteMailRequest) (*utils_proto.DatabaseResponse, error) {
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	username := request.Username
	res, err := c.Conn.Query(context, "UPDATE overflow.mails SET sender = 'null' WHERE id = $1 AND sender = $2;", mail.Id, username)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res.Close()
	res, err = c.Conn.Query(context, "UPDATE overflow.mails SET addressee = 'null' WHERE id = $1 AND addressee = $2;", mail.Id, username)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res.Close()
	res, err = c.Conn.Query(context, "DELETE FROM overflow.mails WHERE sender like 'null' and addressee like 'null';")
	if err == nil {
		res.Close()
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

//Прочитать письмо
func (c *Database) ReadMail(context context.Context, request *repository_proto.ReadMailRequest) (*utils_proto.DatabaseResponse, error) {
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res, err := c.Conn.Query(context, "UPDATE overflow.mails SET read = $1 WHERE id = $2;", true, mail.Id)
	if err == nil {
		res.Close()
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Получить письмо по ID
func (c *Database) GetMailInfoById(context context.Context, request *repository_proto.GetMailInfoByIdRequest) (*repository_proto.ResponseMail, error) {
	var mail models.Mail
	mailBytes, _ := json.Marshal(mail)
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.mails WHERE Id = $1;", request.MailId)
	if err != nil {
		return &repository_proto.ResponseMail{
			Mail: mailBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMail{
				Mail: mailBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Id = values[0].(int32)
		mail.ClientId = values[1].(int32)
		mail.Sender = values[2].(string)
		mail.Addressee = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Theme = values[5].(string)
		mail.Text = values[6].(string)
		mail.Files = values[7].(string)
		mail.Read = values[8].(bool)
	}
	mailBytes, _ = json.Marshal(mail)
	return &repository_proto.ResponseMail{
		Mail: mailBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Получить входящие сообщения пользователя
func (c *Database) GetIncomeMails(context context.Context, request *repository_proto.GetIncomeMailsRequest) (*repository_proto.ResponseMails, error) {
	var results []models.Mail
	resultsBytes, _ := json.Marshal(results)
	rows, err := c.Conn.Query(context, "SELECT * FROM getIncomeMails($1);", request.UserId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: resultsBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail models.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Mails: resultsBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Sender = values[0].(string)
		mail.Theme = values[1].(string)
		mail.Text = values[2].(string)
		mail.Files = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Read = values[5].(bool)
		mail.Id = values[6].(int32)
		results = append(results, mail)
	}
	resultsBytes, _ = json.Marshal(results)
	return &repository_proto.ResponseMails{
		Mails: resultsBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

//Получить отправленные пользователем сообщения
func (c *Database) GetOutcomeMails(context context.Context, request *repository_proto.GetOutcomeMailsRequest) (*repository_proto.ResponseMails, error) {
	var results []models.Mail
	resultsBytes, _ := json.Marshal(results)
	rows, err := c.Conn.Query(context, "SELECT * FROM getOutcomeMails($1);", request.UserId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: resultsBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail models.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Mails: resultsBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Addressee = values[0].(string)
		mail.Theme = values[1].(string)
		mail.Text = values[2].(string)
		mail.Files = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Id = values[5].(int32)
		results = append(results, mail)
	}
	resultsBytes, _ = json.Marshal(results)
	return &repository_proto.ResponseMails{
		Mails: resultsBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

func (c *Database) GetFolderById(context context.Context, request *repository_proto.GetFolderByIdRequest) (*repository_proto.ResponseFolder, error) {
	var folder models.Folder
	folderBytes, _ := json.Marshal(folder)
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.folders WHERE id = $1;", request.FolderId)
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
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.folders WHERE name = $1 AND user_id = $2;", request.FolderName, request.UserId)
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
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.folders WHERE user_id=$1;", request.UserId)
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
	rows, err := c.Conn.Query(context, "SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id=$1;", request.FolderId)
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
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: mailsBytes,
			}, err
		}
		mailId := values[0].(int32)
		resp, err := c.GetMailInfoById(context, &repository_proto.GetMailInfoByIdRequest{
			MailId: mailId,
		})
		if err != nil || resp.Response.Status != utils_proto.DatabaseStatus_OK {
			return &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: mailsBytes,
			}, err
		}
		var mail models.Mail
		err = json.Unmarshal(resp.Mail, &mail)
		if err != nil {
			return &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: mailsBytes,
			}, err
		}
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
	rows, err := c.Conn.Query(context, "DELETE FROM overflow.folders WHERE id=$1;", request.FolderId)
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
	rows, err := c.Conn.Query(context, "UPDATE overflow.folders SET name=$1 WHERE id=$2;", request.NewName, request.FolderId)
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
	rows, err := c.Conn.Query(context, "INSERT INTO overflow.folder_to_mail(folder_id, mail_id) VALUES ($1, $2);", request.FolderId, request.MailId)
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
		rowsRestore, err := c.Conn.Query(context, "UPDATE overflow.mails SET only_folder=$1 WHERE id=$2;", !request.Restore, request.FolderId)
		if err != nil {
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		defer rowsRestore.Close()
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
			Username: request.Username,
		})
	}
}