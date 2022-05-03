package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"time"

	//log "github.com/sirupsen/logrus"
)


// Добавить письмо
func (c *Database) AddMail(context context.Context, request *repository_proto.AddMailRequest) (*utils_proto.DatabaseResponse, error) {
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res, err := c.Conn.Query(context, "INSERT INTO overflow.mails(client_id, sender, addressee, theme, text, files, date) VALUES ($1, $2, $3, $4, $5, $6, $7);", mail.ClientId, mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date)
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
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	username := request.Username
	res, err := c.Conn.Query(context, "UPDATE overflow.mails SET sender = 'null' WHERE id = $1 AND sender = $2;", mail.Id, username)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res.Close()
	res, err = c.Conn.Query(context, "UPDATE overflow.mails SET addressee = 'null' WHERE id = $1 AND addressee = $2;", mail.Id, username)
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
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res, err := c.Conn.Query(context, "UPDATE overflow.mails SET read = $1 WHERE id = $2;", request.Read, mail.Id)
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
	var mail models.Mail
	mailBytes, _ := json.Marshal(mail)
	rows, err := c.Conn.Query(context, "SELECT * FROM overflow.mails WHERE Id = $1;", request.MailId)
	if err != nil {
		return &repository_proto.ResponseMail{
			Mail: mailBytes,
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
				Mail: mailBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Id = values[0].(int32)
		mail.ClientId = values[1].(int32)
		mail.Sender = values[2].(string)
		mail.Addressee = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Theme = values[5].(string)
		mail.Text = values[6].(string)
		mail.Files = values[7].(string)
		mail.Read = values[8].(bool)
	}
	mailBytes, _ = json.Marshal(mail)
	return &repository_proto.ResponseMail{
		Mail: mailBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Получить входящие сообщения пользователя
func (c *Database) GetIncomeMails(context context.Context, request *repository_proto.GetIncomeMailsRequest) (*repository_proto.ResponseMails, error) {
	var results []models.Mail
	resultsBytes, _ := json.Marshal(results)
	rows, err := c.Conn.Query(context, "SELECT * FROM getIncomeMails($1);", request.UserId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: resultsBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail models.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Mails: resultsBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Sender = values[0].(string)
		mail.Theme = values[1].(string)
		mail.Text = values[2].(string)
		mail.Files = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Read = values[5].(bool)
		mail.Id = values[6].(int32)
		results = append(results, mail)
	}
	resultsBytes, _ = json.Marshal(results)
	return &repository_proto.ResponseMails{
		Mails: resultsBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

//Получить отправленные пользователем сообщения
func (c *Database) GetOutcomeMails(context context.Context, request *repository_proto.GetOutcomeMailsRequest) (*repository_proto.ResponseMails, error) {
	var results []models.Mail
	resultsBytes, _ := json.Marshal(results)
	rows, err := c.Conn.Query(context, "SELECT * FROM getOutcomeMails($1);", request.UserId)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: resultsBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail models.Mail
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseMails{
				Mails: resultsBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		mail.Addressee = values[0].(string)
		mail.Theme = values[1].(string)
		mail.Text = values[2].(string)
		mail.Files = values[3].(string)
		mail.Date = values[4].(time.Time)
		mail.Id = values[5].(int32)
		results = append(results, mail)
	}
	resultsBytes, _ = json.Marshal(results)
	return &repository_proto.ResponseMails{
		Mails: resultsBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}
