package sqlmocker

import (
	"OverflowBackend/internal/models"
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

type SQLMock struct {
	url string
	db *sql.DB
	mock sqlmock.Sqlmock
}

func (d *SQLMock) Create(url string) (err error) {
	d.url = url
	d.db, d.mock, err = sqlmock.New()
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return nil
}

func (d *SQLMock) GetUserInfoByUsername(username string) (models.User, error) {
	var user models.User
	rows, err := d.db.Query("Select * from overflow.users where username = $1", username)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password, &user.Username)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (d *SQLMock) GetUserInfoById(userId int32) (models.User, error) {
	var user models.User
	rows, err := d.db.Query("Select * from overflow.users where Id = $1", userId)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password, &user.Username)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (d *SQLMock) AddUser(user models.User) error {
	_, err := d.db.Query("insert into overflow.users(first_name, last_name, password, username) values ($1, $2, $3, $4);", user.FirstName, user.LastName, user.Password, user.Username)
	return err
}

func (d *SQLMock) ChangeUserPassword(user models.User, newPassword string) error {
	_, err := d.db.Query("UPDATE overflow.users set password = $1 where id = $2;", newPassword, user.Id)
	return err
}

func (d *SQLMock) AddMail(mail models.Mail) error {
	_, err := d.db.Query("insert into overflow.mails(client_id, sender, addressee, theme, text, files, date) values($1, $2, $3, $4, $5, $6, $7);", mail.Client_id, mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date)
	return err
}

func (d *SQLMock) DeleteMail(mail models.Mail, username string) error {
	_, err := d.db.Query("UPDATE overflow.mails set sender = 'null' where id = $1 and sender = $2;", mail.Id, username)
	if err != nil {
		return err
	}
	_, err = d.db.Query("UPDATE overflow.mails set addressee = 'null' where id = $1 and addressee = $2;", mail.Id, username)
	if err != nil {
		return err
	}
	_, err = d.db.Query("DELETE FROM overflow.mails WHERE sender like 'null' and addressee like 'null';")
	return err
}

func (d *SQLMock) ReadMail(mail models.Mail) error {
	_, err := d.db.Query("UPDATE overflow.mails set read = $1 where id = $2;", true, mail.Id)
	return err
}

func (d *SQLMock) GetMailInfoById(mailId int32) (models.Mail, error) {
	var mail models.Mail
	rows, err := d.db.Query("Select * from overflow.mails where Id = $1", mailId)
	if err != nil {
		return mail, err
	}
	for rows.Next() {
		err := rows.Scan(&mail.Id, &mail.Client_id, &mail.Sender, &mail.Addressee, &mail.Date, &mail.Theme, &mail.Text, &mail.Files, &mail.Read)
		if err != nil {
			return mail, err
		}
	}
	return mail, nil
}

func (d *SQLMock) GetIncomeMails(userId int32) ([]models.Mail, error) {
	var results []models.Mail
	rows, err := d.db.Query("Select * from getIncomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mail models.Mail
		err := rows.Scan(&mail.Sender, &mail.Theme, &mail.Text, &mail.Files, &mail.Date, &mail.Read, &mail.Id)
		if err != nil {
			return results, err
		}
		results = append(results, mail)
	}
	return results, nil
}

func (d *SQLMock) GetOutcomeMails(userId int32) ([]models.Mail, error) {
	var results []models.Mail
	rows, err := d.db.Query("Select * from getOutcomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mail models.Mail
		err := rows.Scan(&mail.Addressee, &mail.Theme, &mail.Text, &mail.Files, &mail.Date, &mail.Id)
		if err != nil {
			return results, err
		}
		results = append(results, mail)
	}
	return results, nil
}
