package postgres

import (
	"OverflowBackend/internal/models"
	"context"
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
func (c *Database) GetUserInfoByUsername(username string) (models.User, error) {
	var user models.User
	rows, err := c.Conn.Query(context.Background(), "Select * from overflow.users where username = $1;", username)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return user, err
		}
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	return user, nil
}

// Получить данные пользователя по его айди в бд
func (c *Database) GetUserInfoById(userId int32) (models.User, error) {
	var user models.User
	rows, err := c.Conn.Query(context.Background(), "Select * from overflow.users where id = $1;", userId)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return user, err
		}
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	return user, nil
}

// Добавить пользователя
func (c *Database) AddUser(user models.User) error {
	res, err := c.Conn.Query(context.Background(), "insert into overflow.users(first_name, last_name, password, username) values ($1, $2, $3, $4);", user.FirstName, user.LastName, user.Password, user.Username)
	if err == nil {
		res.Close()
	}
	return err
}

// Изменить пароль
func (c *Database) ChangeUserPassword(user models.User, newPassword string) error {
	_, err := c.Conn.Exec(context.Background(), "UPDATE overflow.users set password = $1 where id = $2;", newPassword, user.Id)
	return err
}

// Изменить имя
func (c *Database) ChangeUserFirstName(user models.User, newFirstName string) error {
	_, err := c.Conn.Exec(context.Background(), "UPDATE overflow.users set first_name = $1 where id = $2;", newFirstName, user.Id)
	return err
}

// Изменить фамилию
func (c *Database) ChangeUserLastName(user models.User, newLastName string) error {
	_, err := c.Conn.Exec(context.Background(), "UPDATE overflow.users set last_name = $1 where id = $2;", newLastName, user.Id)
	return err
}

// Добавить письмо
func (c *Database) AddMail(mail models.Mail) error {
	res, err := c.Conn.Query(context.Background(), "insert into overflow.mails(client_id, sender, addressee, theme, text, files, date) values($1, $2, $3, $4, $5, $6, $7);", mail.Client_id, mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date)
	if err == nil {
		res.Close()
	}
	return err
}

//Удалить письмо
func (c *Database) DeleteMail(mail models.Mail, username string) error {
	res, err := c.Conn.Query(context.Background(), "UPDATE overflow.mails set sender = 'null' where id = $1 and sender = $2;", mail.Id, username)
	if err != nil {
		return err
	}
	res.Close()
	res, err = c.Conn.Query(context.Background(), "UPDATE overflow.mails set addressee = 'null' where id = $1 and addressee = $2;", mail.Id, username)
	if err != nil {
		return err
	}
	res.Close()
	res, err = c.Conn.Query(context.Background(), "DELETE FROM overflow.mails WHERE sender like 'null' and addressee like 'null';")
	if err == nil {
		res.Close()
	}
	return err
}

//Прочитать письмо
func (c *Database) ReadMail(mail models.Mail) error {
	res, err := c.Conn.Query(context.Background(), "UPDATE overflow.mails set read = $1 where id = $2;", true, mail.Id)
	if err == nil {
		res.Close()
	}
	return err
}

// Получить письмо по ID
func (c *Database) GetMailInfoById(mailId int32) (models.Mail, error) {
	var mail models.Mail
	rows, err := c.Conn.Query(context.Background(), "Select * from overflow.mails where Id = $1", mailId)
	if err != nil {
		return mail, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return mail, err
		}
		mail.Id = values[0].(int32)
		mail.Client_id = values[1].(int32)
		mail.Sender = values[2].(string)
		mail.Addressee = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Theme = values[5].(string)
		mail.Text = values[6].(string)
		mail.Files = values[7].(string)
		mail.Read = values[8].(bool)
	}
	return mail, nil
}

// Получить входящие сообщения пользователя
func (c *Database) GetIncomeMails(userId int32) ([]models.Mail, error) {
	var results []models.Mail
	rows, err := c.Conn.Query(context.Background(), "Select * from getIncomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		var mails models.Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.Sender = values[0].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
		mails.Files = values[3].(string)
		mails.Date = values[4].(time.Time)
		mails.Read = values[5].(bool)
		mails.Id = values[6].(int32)
		results = append(results, mails)
	}
	return results, nil
}

//Получить отправленные пользователем сообщения
func (c *Database) GetOutcomeMails(userId int32) ([]models.Mail, error) {
	var results []models.Mail
	rows, err := c.Conn.Query(context.Background(), "Select * from getOutcomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		var mails models.Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.Addressee = values[0].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
		mails.Files = values[3].(string)
		mails.Date = values[4].(time.Time)
		mails.Id = values[5].(int32)
		results = append(results, mails)
	}
	return results, nil
}
