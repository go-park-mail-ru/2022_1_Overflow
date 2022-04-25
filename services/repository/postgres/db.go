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
	res, err := c.Conn.Query(context, "insert into overflow.mails(client_id, sender, addressee, theme, text, files, date) values($1, $2, $3, $4, $5, $6, $7);", mail.ClientId, mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date)
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
	rows, err := c.Conn.Query(context, "Select * from overflow.mails where Id = $1", request.MailId)
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
	rows, err := c.Conn.Query(context, "Select * from getIncomeMails($1)", request.UserId)
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
	rows, err := c.Conn.Query(context, "Select * from getOutcomeMails($1)", request.UserId)
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
