package postgres

import (
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
	var user utils_proto.User
	rows, err := c.Conn.Query(context, "Select * from overflow.users where username = $1;", request.Username)
	if err != nil {
		return &repository_proto.ResponseUser{
			User: &user,
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
				User: &user,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	return &repository_proto.ResponseUser{
		User: &user,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Получить данные пользователя по его айди в бд
func (c *Database) GetUserInfoById(context context.Context, request *repository_proto.GetUserInfoByIdRequest) (*repository_proto.ResponseUser, error) {
	var user utils_proto.User
	rows, err := c.Conn.Query(context, "Select * from overflow.users where id = $1;", request.UserId)
	if err != nil {
		return &repository_proto.ResponseUser{
			User: &user,
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
				User: &user,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	return &repository_proto.ResponseUser{
		User: &user,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Добавить пользователя
func (c *Database) AddUser(context context.Context, request *repository_proto.AddUserRequest) (*utils_proto.DatabaseResponse, error) {
	user := request.User
	res, err := c.Conn.Query(context, "insert into overflow.users(first_name, last_name, password, username) values ($1, $2, $3, $4);", user.FirstName, user.LastName, user.Password, user.Username)
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

// Изменить пароль
func (c *Database) ChangeUserPassword(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	_, err := c.Conn.Exec(context, "UPDATE overflow.users set password = $1 where id = $2;", request.Data, request.User.Id)
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
	_, err := c.Conn.Exec(context, "UPDATE overflow.users set first_name = $1 where id = $2;", request.Data, request.User.Id)
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
	_, err := c.Conn.Exec(context, "UPDATE overflow.users set last_name = $1 where id = $2;", request.Data, request.User.Id)
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
	mail := request.Mail
	res, err := c.Conn.Query(context, "insert into overflow.mails(client_id, sender, addressee, theme, text, files, date) values($1, $2, $3, $4, $5, $6, $7);", mail.ClientId, mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date.AsTime())
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
	mail := request.Mail
	username := request.Username
	res, err := c.Conn.Query(context, "UPDATE overflow.mails set sender = 'null' where id = $1 and sender = $2;", mail.Id, username)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res.Close()
	res, err = c.Conn.Query(context, "UPDATE overflow.mails set addressee = 'null' where id = $1 and addressee = $2;", mail.Id, username)
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
	mail := request.Mail
	res, err := c.Conn.Query(context, "UPDATE overflow.mails set read = $1 where id = $2;", true, mail.Id)
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
	var mail utils_proto.Mail
	rows, err := c.Conn.Query(context, "Select * from overflow.mails where Id = $1;", request.MailId)
	if err != nil {
		return &repository_proto.ResponseMail{
			Mail: &mail,
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
				Mail: &mail,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Id = values[0].(int32)
		mail.ClientId = values[1].(int32)
		mail.Sender = values[2].(string)
		mail.Addressee = values[3].(string)
		mail.Date = timestamppb.New(values[4].(time.Time))
		mail.Theme = values[5].(string)
		mail.Text = values[6].(string)
		mail.Files = values[7].(string)
		mail.Read = wrapperspb.Bool(values[8].(bool))
	}
	return &repository_proto.ResponseMail{
		Mail: &mail,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Получить входящие сообщения пользователя
func (c *Database) GetIncomeMails(context context.Context, request *repository_proto.GetIncomeMailsRequest) (*repository_proto.ResponseMails, error) {
	var results []*utils_proto.Mail
	rows, err := c.Conn.Query(context, "Select * from getIncomeMails($1);", request.UserId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: results,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail utils_proto.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Mails: results,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Sender = values[0].(string)
		mail.Theme = values[1].(string)
		mail.Text = values[2].(string)
		mail.Files = values[3].(string)
		mail.Date = timestamppb.New(values[4].(time.Time))
		mail.Read = wrapperspb.Bool(values[5].(bool))
		mail.Id = values[6].(int32)
		results = append(results, &mail)
	}
	return &repository_proto.ResponseMails{
		Mails: results,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

//Получить отправленные пользователем сообщения
func (c *Database) GetOutcomeMails(context context.Context, request *repository_proto.GetOutcomeMailsRequest) (*repository_proto.ResponseMails, error) {
	var results []*utils_proto.Mail
	rows, err := c.Conn.Query(context, "Select * from getOutcomeMails($1);", request.UserId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: results,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail utils_proto.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Mails: results,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Addressee = values[0].(string)
		mail.Theme = values[1].(string)
		mail.Text = values[2].(string)
		mail.Files = values[3].(string)
		mail.Date = timestamppb.New(values[4].(time.Time))
		mail.Id = values[5].(int32)
		results = append(results, &mail)
	}
	return &repository_proto.ResponseMails{
		Mails: results,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

func (c *Database) GetFolderById(context context.Context, request *repository_proto.GetFolderByIdRequest) (*repository_proto.ResponseFolder, error) {
	var folder utils_proto.Folder
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.folders WHERE id = $1;", request.FolderId)
	if err != nil {
		return &repository_proto.ResponseFolder{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folder: &folder,
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
				Folder: &folder,
			}, err
		}
		folder.Id = values[0].(int32)
		folder.Name = values[1].(string)
		folder.UserId = values[2].(int32)
		folder.CreatedAt = timestamppb.New(values[3].(time.Time))
	}
	return &repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: &folder,
	}, nil
}

func (c *Database) GetFolderByName(context context.Context, request *repository_proto.GetFolderByNameRequest) (*repository_proto.ResponseFolder, error) {
	var folder utils_proto.Folder
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.folders WHERE name = $1 AND user_id = $2;", request.FolderName, request.UserId)
	if err != nil {
		return &repository_proto.ResponseFolder{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folder: &folder,
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
				Folder: &folder,
			}, err
		}
		folder.Id = values[0].(int32)
		folder.Name = values[1].(string)
		folder.UserId = values[2].(int32)
		folder.CreatedAt = timestamppb.New(values[3].(time.Time))
	}
	return &repository_proto.ResponseFolder{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folder: &folder,
	}, nil
}

func (c *Database) GetFoldersByUser(context context.Context, request *repository_proto.GetFoldersByUserRequest) (*repository_proto.ResponseFolders, error) {
	var folders []*utils_proto.Folder
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.folders WHERE user_id=$1;", request.UserId)
	if err != nil {
		return &repository_proto.ResponseFolders{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Folders: folders,
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
				Folders: folders,
			}, err
		}
		var folder utils_proto.Folder
		folder.Id = values[0].(int32)
		folder.Name = values[1].(string)
		folder.UserId = values[2].(int32)
		folder.CreatedAt = timestamppb.New(values[3].(time.Time))
		folders = append(folders, &folder)
	}
	return &repository_proto.ResponseFolders{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Folders: folders,
	}, nil
}

func (c *Database) GetFolderMail(context context.Context, request *repository_proto.GetFolderMailRequest) (*repository_proto.ResponseMails, error) {
	var mails []*utils_proto.Mail
	rows, err := c.Conn.Query(context, "SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id=$1;", request.FolderId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			Mails: mails,
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
				Mails: mails,
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
				Mails: mails,
			}, err
		}
		mails = append(mails, resp.Mail)
	}
	return &repository_proto.ResponseMails{
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
		Mails: mails,
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