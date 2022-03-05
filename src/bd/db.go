package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Структура для юзера
type UserT struct {
	id        int32
	firstName string
	lastName  string
	password  string
	email     string
}

// Структура письма
type Mail struct {
	client_id                             int32
	sender, addressee, theme, text, files string
	date                                  time.Time
}

//Получить данные пользователя по его почте
func GetUserInfoByEmail(userEmail string, conn *pgxpool.Pool) (UserT, error) {
	var user UserT
	rows, err := conn.Query(context.Background(), "Select * from overflow.users where email = $1", userEmail)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return user, err
		}
		user.id = values[0].(int32)
		user.firstName = values[1].(string)
		user.lastName = values[2].(string)
		user.password = values[3].(string)
		user.email = values[4].(string)
	}
	return user, nil
}

//Получить данные пользователя по его айди в бд
func GetUserInfoById(userId int, conn *pgxpool.Pool) (UserT, error) {
	var user UserT
	rows, err := conn.Query(context.Background(), "Select * from overflow.users where id = $1", userId)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return user, err
		}
		user.id = values[0].(int32)
		user.firstName = values[1].(string)
		user.lastName = values[2].(string)
		user.password = values[3].(string)
		user.email = values[4].(string)
	}
	return user, nil
}

// Добавить юзера
func AddUser(user UserT, conn *pgxpool.Pool) {
	conn.QueryRow(context.Background(), "insert into overflow.users(first_name, last_name, password, email) values ($1, $2, $3, $4);", user.firstName, user.lastName, user.password, user.email)
}

// Добавить почту
func AddMail(email Mail, conn *pgxpool.Pool) {
	conn.QueryRow(context.Background(), "insert into overflow.mails(client_id, sender, addressee,theme,  text, files, date) values($1, $2, $3, $4, $5, $6, $7);", email.client_id, email.sender, email.addressee, email.text, email.files, email.date)

}

// Получить входящие сообщения пользователя
func GetIncomeMails(userId int, conn *pgxpool.Pool) ([]Mail, error) {
	var results []Mail
	rows, err := conn.Query(context.Background(), "Select * from getIncomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mails Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.sender = values[0].(string)
		mails.files = values[3].(string)
		mails.theme = values[1].(string)
		mails.text = values[2].(string)
		mails.date = values[4].(time.Time)
		results = append(results, mails)
	}
	return results, nil
}

//Получить отправленные пользователем сообщения
func GetOutcomeMails(userId int, conn *pgxpool.Pool) ([]Mail, error) {
	var results []Mail
	rows, err := conn.Query(context.Background(), "Select * from getOutcomeMails($1)", userId)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var mails Mail
		values, err := rows.Values()
		if err != nil {
			return results, err
		}
		mails.addressee = values[0].(string)
		mails.files = values[3].(string)
		mails.theme = values[1].(string)
		mails.text = values[2].(string)
		mails.date = values[4].(time.Time)
		results = append(results, mails)

	}
	return results, nil
}

func main() {
	urlExample := "postgres://postgres:postgres@localhost:5432/postgres"
	conn, err := pgxpool.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	var user UserT
	user.firstName = "Mikhail"
	user.lastName = "Rabinovich"
	user.email = "animelov123123er69@overflow.ru"
	user.password = "12312213123312"
	//AddUser(user, conn)
	//user = GetUserInfoById(2)
	//fmt.Print(user)
	//user = GetUserInfoByEmail("animelov123123er69@overflow.ru", conn)
	//fmt.Print(user, conn)
	results, err := GetOutcomeMails(1, conn)
	fmt.Print(results)
	results, err = GetIncomeMails(1, conn)
	fmt.Print(results)
}
