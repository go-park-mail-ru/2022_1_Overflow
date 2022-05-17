package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// Добавить письмо
func (c *Database) AddMail(context context.Context, request *repository_proto.AddMailRequest) (*utils_proto.DatabaseExtendResponse, error) {
	var mail models.Mail
	err := json.Unmarshal(request.Mail, &mail)
	if err != nil {
		return &utils_proto.DatabaseExtendResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
			Param:  "",
		}, err
	}
	res, err := c.Conn.Query(context, "INSERT INTO overflow.mails(sender, addressee, theme, text, files, date) VALUES ($1, $2, $3, $4, $5, $6);", mail.Sender, mail.Addressee, mail.Theme, mail.Text, mail.Files, mail.Date)
	if err == nil {
		res.Close()
		row := c.Conn.QueryRow(context, "SELECT max(id) FROM overflow.mails WHERE sender = $1", mail.Sender)
		var mailid int
		if err := row.Scan(&mailid); err != nil {
			return &utils_proto.DatabaseExtendResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
				Param:  "",
			}, err
		}
		return &utils_proto.DatabaseExtendResponse{
			Status: utils_proto.DatabaseStatus_OK,
			Param:  strconv.Itoa(mailid),
		}, nil
	} else {
		return &utils_proto.DatabaseExtendResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
			Param:  "",
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
	userId := request.UserId
	res, err := c.Conn.Query(context, "UPDATE overflow.mails SET sender = NULL WHERE id = $1 AND sender IN (SELECT username FROM overflow.users WHERE id=$2);", mail.Id, userId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res.Close()
	res, err = c.Conn.Query(context, "UPDATE overflow.mails SET addressee = NULL WHERE id = $1 AND addressee IN (SELECT username FROM overflow.users WHERE id=$2);", mail.Id, userId)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res.Close()
	res, err = c.Conn.Query(context, "DELETE FROM overflow.mails WHERE sender is NULL and addressee is NULL;")
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
func (c *Database) GetMailInfoById(context context.Context, request *repository_proto.GetMailInfoByIdRequest) (response *repository_proto.ResponseMail, err error) {
	var mail models.Mail
	mailBytes, _ := json.Marshal(mail)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseMail{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mail: mailBytes,
			}
		}
	}()
	rows, err := c.Conn.Query(context, "SELECT id, sender, addressee, date, theme, text, files, read FROM overflow.mails WHERE Id = $1;", request.MailId)
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
		mail.Sender = values[1].(string)
		mail.Addressee = values[2].(string)
		mail.Date = values[3].(time.Time)
		mail.Theme = values[4].(string)
		mail.Text = values[5].(string)
		mail.Files = values[6].(string)
		mail.Read = values[7].(bool)
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
func (c *Database) GetIncomeMails(context context.Context, request *repository_proto.GetIncomeMailsRequest) (response *repository_proto.ResponseMails, err error) {
	var results models.MailList
	resultsBytes, _ := json.Marshal(results)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: resultsBytes,
			}
		}
	}()
	var count int
	err = c.Conn.QueryRow(context, "SELECT COUNT(*) FROM overflow.mails WHERE id NOT IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id IN (SELECT id FROM overflow.folders WHERE user_id=$1) AND only_folder=true) AND addressee IN (SELECT username FROM overflow.users WHERE id=$1);", request.UserId).Scan(&count)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: resultsBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	results.Amount = count
	rows, err := c.Conn.Query(context, "SELECT sender, theme, text, files, date, read, id FROM overflow.mails WHERE id NOT IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id IN (SELECT id FROM overflow.folders WHERE user_id=$1) AND only_folder=true) AND addressee IN (SELECT username FROM overflow.users WHERE id=$1) ORDER BY date DESC OFFSET $3 LIMIT $2;", request.UserId, request.Limit, request.Offset)
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
		log.Debug("Получены значения: ", values)
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
		results.Mails = append(results.Mails, mail)
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
func (c *Database) GetOutcomeMails(context context.Context, request *repository_proto.GetOutcomeMailsRequest) (response *repository_proto.ResponseMails, err error) {
	var results models.MailList
	resultsBytes, _ := json.Marshal(results)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
			log.Error(err)
			response = &repository_proto.ResponseMails{
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
				Mails: resultsBytes,
			}
		}
	}()
	var count int
	err = c.Conn.QueryRow(context, "SELECT COUNT(*) FROM overflow.mails WHERE id NOT IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id IN (SELECT id FROM overflow.folders WHERE user_id=$1) AND only_folder=true) AND sender IN (SELECT username FROM overflow.users WHERE id=$1);", request.UserId).Scan(&count)
	if err != nil {
		return &repository_proto.ResponseMails{
			Mails: resultsBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	results.Amount = count
	rows, err := c.Conn.Query(context, "SELECT addressee, theme, text, files, date, id FROM overflow.mails WHERE id NOT IN (SELECT mail_id FROM overflow.folder_to_mail WHERE folder_id IN (SELECT id FROM overflow.folders WHERE user_id=$1) AND only_folder=true) AND sender IN (SELECT username FROM overflow.users WHERE id=$1) ORDER BY date DESC OFFSET $3 LIMIT $2;", request.UserId, request.Limit, request.Offset)
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
		results.Mails = append(results.Mails, mail)
	}
	resultsBytes, _ = json.Marshal(results)
	return &repository_proto.ResponseMails{
		Mails: resultsBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}
