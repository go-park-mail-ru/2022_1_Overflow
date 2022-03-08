package main

import (
	"context"
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
	Client_id                             int32
	Sender, Addressee, Theme, Text, Files string
	Date                                  time.Time
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
		user.Id = values[0].(int32)
		user.FirstName = values[1].(string)
		user.LastName = values[2].(string)
		user.Password = values[3].(string)
		user.Email = values[4].(string)
	}
	return user, nil
}

//Получить данные пользователя по его айди в бд
func GetUserInfoById(userId int, conn *pgxpool.Pool) (UserT, error) {
	var user UserT
	rows, err := conn.Query(context.Background(), "Select * from overflow.users where Id = $1", userId)
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
func AddUser(user UserT, conn *pgxpool.Pool) error {
	_, err := conn.Query(context.Background(), "insert into overflow.users(first_name, last_name, Password, email) values ($1, $2, $3, $4);", user.FirstName, user.LastName, user.Password, user.Email)
	return err
}

// Добавить почту
func AddMail(email Mail, conn *pgxpool.Pool) error {
	_, err := conn.Query(context.Background(), "insert into overflow.mails(client_id, sender, addressee,theme,  text, files, date) values($1, $2, $3, $4, $5, $6, $7);", email.Client_id, email.Sender, email.Addressee, email.Text, email.Files, email.Date)
	return err
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
		mails.Addressee = values[0].(string)
		mails.Files = values[3].(string)
		mails.Theme = values[1].(string)
		mails.Text = values[2].(string)
		mails.Date = values[4].(time.Time)
		results = append(results, mails)

	}
	return results, nil
}
