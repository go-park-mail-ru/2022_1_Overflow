package postgres

import (
	"OverflowBackend/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	url  string
	conn *pgxpool.Pool
}

func (c *Database) Create(url string) (err error) {
	c.url = url
	c.conn, err = pgxpool.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return nil
}

// Получить данные пользователя по его почте
func (c *Database) GetUserInfoByEmail(userEmail string) (models.User, error) {
	var user models.User
	rows, err := c.conn.Query(context.Background(), "Select * from overflow.users where email = $1", userEmail)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return user, err
		}
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Email = values[4].(string)
	}
	return user, nil
}

// Получить данные пользователя по его айди в бд
func (c *Database) GetUserInfoById(userId int32) (models.User, error) {
	var user models.User
	rows, err := c.conn.Query(context.Background(), "Select * from overflow.users where Id = $1", userId)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return user, err
		}
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Email = values[4].(string)
	}
	return user, nil
}

// Добавить пользователя
func (c *Database) AddUser(user models.User) error {
	_, err := c.conn.Query(context.Background(), "insert into overflow.users(first_name, last_name, password, email) values ($1, $2, $3, $4);", user.FirstName, user.LastName, user.Password, user.Email)
	return err
}

// Изменить пароль
func (c *Database) ChangeUserPassword(user models.User, newPassword string) error {
	_, err := c.conn.Query(context.Background(), "UPDATE overflow.users set password = $1 where id = $2;", newPassword, user.Id)
	return err
}

// Добавить письмо
func (c *Database) AddMail(email models.Mail) error {
	_, err := c.conn.Query(context.Background(), "insert into overflow.mails(client_id, sender, addressee,theme,  text, files, date) values($1, $2, $3, $4, $5, $6, $7);", email.Client_id, email.Sender, email.Addressee, email.Text, email.Files, email.Date)
	return err
}

//Удалить письмо
func (c *Database) DeleteMail(email models.Mail) error {
	_, err := c.conn.Query(context.Background(), "delete from overflow.mails where id = &1;", email.Id)
	return err
}

//Прочитать письмо
func (c *Database) ReadMail(email models.Mail) error {
	_, err := c.conn.Query(context.Background(), "UPDATE overflow.mails set read = $1 where id = $2;", true, email.Id)
	return err
}

// Получить письмо по ID
func (c *Database) GetMailInfoById(mailId int) (models.Mail, error) {
	var mail models.Mail
	rows, err := c.conn.Query(context.Background(), "Select * from overflow.mails where Id = $1", mailId)
	if err != nil {
		return mail, err
	}
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
	rows, err := c.conn.Query(context.Background(), "Select * from getIncomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mails models.Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.Sender = values[0].(string)
		mails.Files = values[3].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
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
	rows, err := c.conn.Query(context.Background(), "Select * from getOutcomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mails models.Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.Addressee = values[0].(string)
		mails.Files = values[3].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
		mails.Date = values[4].(time.Time)
		mails.Id = values[5].(int32)
		results = append(results, mails)
	}
	return results, nil
}
