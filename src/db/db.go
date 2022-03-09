package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Структура для юзера
type UserT struct {
	Id        int32
	FirstName string
	LastName  string
	Password  string
	Email     string
}

// Структура письма
type Mail struct {
	Client_id 	int32		`json:"id"`
	Sender		string		`json:"sender"`
	Addressee	string		`json:"addressee"`
	Theme		string		`json:"theme"`
	Text		string		`json:"text"`
	Files 		string		`json:"files"`
	Date		time.Time	`json:"date"`
}

type DatabaseConnection struct {
	url string
	conn *pgxpool.Pool
}

func (c *DatabaseConnection) Create(url string) (err error) {
	c.url = url
	c.conn, err = pgxpool.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	return
}

//Получить данные пользователя по его почте
func (c *DatabaseConnection) GetUserInfoByEmail(userEmail string) (UserT, error) {
	var user UserT
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

//Получить данные пользователя по его айди в бд
func (c *DatabaseConnection) GetUserInfoById(userId int) (UserT, error) {
	var user UserT
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

// Добавить юзера
func (c *DatabaseConnection) AddUser(user UserT) error {
	_, err := c.conn.Query(context.Background(), "insert into overflow.users(first_name, last_name, password, email) values ($1, $2, $3, $4);", user.FirstName, user.LastName, user.Password, user.Email)
	return err
}

// Добавить почту
func (c *DatabaseConnection) AddMail(email Mail) error {
	_, err := c.conn.Query(context.Background(), "insert into overflow.mails(client_id, sender, addressee,theme,  text, files, date) values($1, $2, $3, $4, $5, $6, $7);", email.Client_id, email.Sender, email.Addressee, email.Text, email.Files, email.Date)
	return err
}

// Получить входящие сообщения пользователя
func (c *DatabaseConnection) GetIncomeMails(userId int32) ([]Mail, error) {
	var results []Mail
	rows, err := c.conn.Query(context.Background(), "Select * from getIncomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mails Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.Sender = values[0].(string)
		mails.Files = values[3].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
		mails.Date = values[4].(time.Time)
		results = append(results, mails)
	}
	return results, nil
}

//Получить отправленные пользователем сообщения
func (c *DatabaseConnection) GetOutcomeMails(userId int) ([]Mail, error) {
	var results []Mail
	rows, err := c.conn.Query(context.Background(), "Select * from getOutcomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mails Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.Addressee = values[0].(string)
		mails.Files = values[3].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
		mails.Date = values[4].(time.Time)
		results = append(results, mails)

	}
	return results, nil
}

func main() {
	urlExample := "postgres://postgres:postgres@localhost:5432/postgres"
	var conn DatabaseConnection
	err := conn.Create(urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to c.connect to database: %v\n", err)
		os.Exit(1)
	}
	var user UserT
	user.FirstName = "Mikhail"
	user.LastName = "Rabinovich"
	user.Email = "animelov123123er69@overflow.ru"
	user.Password = "12312213123312"
	//AddUser(user, c.conn)
	//user = GetUserInfoById(2)
	//fmt.Print(user)
	//user = GetUserInfoByEmail("animelov123123er69@overflow.ru", c.conn)
	//fmt.Print(user, c.conn)
	results, err := conn.GetOutcomeMails(1)
	fmt.Print(results)
	results, err = conn.GetIncomeMails(1)
	fmt.Print(results)
}
