package repository_test

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/repository/postgres"
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock"
)

func TestUserGet(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		FirstName: "John",
		LastName:  "Doe",
		Password:  "passw",
		Username:  "john",
	}

	mock.ExpectBegin()

	mock.ExpectQuery(
		regexp.QuoteMeta("insert into overflow.users(first_name, last_name, password, username) values ($1, $2, $3, $4);"),
	).WithArgs(
		user.FirstName,
		user.LastName,
		user.Password,
		user.Username,
	).WillReturnRows(&pgxmock.Rows{})

	mock.ExpectQuery(
		regexp.QuoteMeta("Select * from overflow.users where username = $1;"),
	).WithArgs(
		user.Username,
	).WillReturnRows(pgxmock.NewRows([]string{"id", "first_name", "last_name", "password", "username"}))

	mock.ExpectQuery(
		regexp.QuoteMeta("Select * from overflow.users where id = $1;"),
	).WithArgs(
		user.Id,
	).WillReturnRows(pgxmock.NewRows([]string{"id", "first_name", "last_name", "password", "username"}))

	conn, err := mock.Begin(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	testDB := postgres.Database{
		Db:   mock,
		Conn: conn,
	}

	err = testDB.AddUser(user)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testDB.GetUserInfoByUsername(user.Username)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testDB.GetUserInfoById(user.Id)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}

func TestUserChange(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		FirstName: "John",
		LastName:  "Doe",
		Password:  "passw",
		Username:  "john",
	}
	new_password := "pass"
	new_firstname := "Doe"
	new_lastname := "John"

	mock.ExpectBegin()

	mock.ExpectQuery(
		regexp.QuoteMeta("insert into overflow.users(first_name, last_name, password, username) values ($1, $2, $3, $4);"),
	).WithArgs(
		user.FirstName,
		user.LastName,
		user.Password,
		user.Username,
	).WillReturnRows(&pgxmock.Rows{})

	mock.ExpectQuery(
		regexp.QuoteMeta("UPDATE overflow.users set password = $1 where id = $2;"),
	).WithArgs(
		new_password,
		user.Id,
	).WillReturnRows(&pgxmock.Rows{})

	mock.ExpectQuery(
		regexp.QuoteMeta("UPDATE overflow.users set first_name = $1 where id = $2;"),
	).WithArgs(
		new_firstname,
		user.Id,
	).WillReturnRows(&pgxmock.Rows{})

	mock.ExpectQuery(
		regexp.QuoteMeta("UPDATE overflow.users set last_name = $1 where id = $2;"),
	).WithArgs(
		new_lastname,
		user.Id,
	).WillReturnRows(&pgxmock.Rows{})

	conn, err := mock.Begin(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	testDB := postgres.Database{
		Db:   mock,
		Conn: conn,
	}

	err = testDB.AddUser(user)
	if err != nil {
		t.Error(err)
		return
	}

	err = testDB.ChangeUserPassword(user, new_password)
	if err != nil {
		t.Error(err)
		return
	}

	err = testDB.ChangeUserFirstName(user, new_firstname)
	if err != nil {
		t.Error(err)
		return
	}

	err = testDB.ChangeUserLastName(user, new_lastname)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}

func TestMailBoxChange(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	mail := models.Mail{
		Id:        0,
		Client_id: 0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	mock.ExpectBegin()

	mock.ExpectQuery(
		regexp.QuoteMeta("insert into overflow.mails(client_id, sender, addressee, theme, text, files, date) values($1, $2, $3, $4, $5, $6, $7);"),
	).WithArgs(
		mail.Client_id,
		mail.Sender,
		mail.Addressee,
		mail.Theme,
		mail.Text,
		mail.Files,
		mail.Date,
	).WillReturnRows(&pgxmock.Rows{})

	conn, err := mock.Begin(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	testDB := postgres.Database{
		Db:   mock,
		Conn: conn,
	}

	err = testDB.AddMail(mail)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}
