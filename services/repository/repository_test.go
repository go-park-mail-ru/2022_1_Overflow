package repository_test

/*
import (
	"OverflowBackend/internal/models"
	"OverflowBackend/services/repository/postgres"
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/jackc/pgconn"
	"github.com/pashagolub/pgxmock"
)

func TestAddUser(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "passw",
		Username:  "john",
	}

	mock.ExpectQuery(
		regexp.QuoteMeta("insert into overflow.users(first_name, last_name, password, username) values ($1, $2, $3, $4);"),
	).WithArgs(
		user.Firstname,
		user.Lastname,
		user.Password,
		user.Username,
	).WillReturnRows(&pgxmock.Rows{})

	testDB := postgres.Database{
		Conn:   mock,
	}

	err = testDB.AddUser(user)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}

func TestGetUserInfoByUsername(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "passw",
		Username:  "john",
	}

	mock.ExpectQuery(
		regexp.QuoteMeta("Select * from overflow.users where username = $1;"),
	).WithArgs(
		user.Username,
	).WillReturnRows(pgxmock.NewRows([]string{"id", "first_name", "last_name", "password", "username"}))

	testDB := postgres.Database{
		Conn: mock,
	}

	_, err = testDB.GetUserInfoByUsername(user.Username)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}

func TestGetUserInfoById(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "passw",
		Username:  "john",
	}

	mock.ExpectQuery(
		regexp.QuoteMeta("Select * from overflow.users where id = $1;"),
	).WithArgs(
		user.Id,
	).WillReturnRows(pgxmock.NewRows([]string{"id", "first_name", "last_name", "password", "username"}))

	testDB := postgres.Database{
		Conn: mock,
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

func TestChangeUserPassword(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "passw",
		Username:  "john",
	}
	new_password := "pass"

	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE overflow.users set password = $1 where id = $2;"),
	).WithArgs(
		new_password,
		user.Id,
	).WillReturnResult(pgconn.CommandTag{})

	testDB := postgres.Database{
		Conn:   mock,
	}

	err = testDB.ChangeUserPassword(user, new_password)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}

func TestChangeUserFirstName(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "passw",
		Username:  "john",
	}
	new_firstname := "Doe"

	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE overflow.users set first_name = $1 where id = $2;"),
	).WithArgs(
		new_firstname,
		user.Id,
	).WillReturnResult(pgconn.CommandTag{})

	testDB := postgres.Database{
		Conn: mock,
	}

	err = testDB.ChangeUserFirstName(user, new_firstname)
	if err != nil {
		t.Error(err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
		return
	}
}

func TestChangeUserLastName(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	user := models.User{
		Id:        0,
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "passw",
		Username:  "john",
	}
	new_lastname := "John"

	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE overflow.users set last_name = $1 where id = $2;"),
	).WithArgs(
		new_lastname,
		user.Id,
	).WillReturnResult(pgconn.CommandTag{})

	testDB := postgres.Database{
		Conn: mock,
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

func TestAddMail(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	defer mock.Close(context.Background())

	mail := models.Mail{
		Id:        0,
		Sender:    "test",
		Addressee: "test2",
		Theme:     "test",
		Text:      "test",
		Files:     "files",
		Date:      time.Now(),
		Read:      false,
	}

	mock.ExpectQuery(
		regexp.QuoteMeta("insert into overflow.mails(client_id, sender, addressee, theme, text, files, date) values($1, $2, $3, $4, $5, $6, $7);"),
	).WithArgs(
		mail.Sender,
		mail.Addressee,
		mail.Theme,
		mail.Text,
		mail.Files,
		mail.Date,
	).WillReturnRows(&pgxmock.Rows{})

	testDB := postgres.Database{
		Conn: mock,
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
*/