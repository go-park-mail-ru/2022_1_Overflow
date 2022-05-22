package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/proto/repository_proto"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (c *Database) AddAttachLink(context context.Context, request *repository_proto.AddAttachLinkRequest) (*repository_proto.Nothing, error) {
	res, err := c.Conn.Query(context, "INSERT INTO overflow.attaches(mail_id, filename) VALUES ($1, $2);", request.MailID, request.Filename)
	if err != nil {
		return &repository_proto.Nothing{
			Status: false,
		}, err
	}
	res.Close()
	return &repository_proto.Nothing{
		Status: true,
	}, nil
}

func (c *Database) CheckAttachLink(context context.Context, request *repository_proto.GetAttachRequest) (*repository_proto.Nothing, error) {
	row := c.Conn.QueryRow(context, "SELECT 'exist' FROM overflow.attaches WHERE mail_id = $1 AND filename = $2;", request.MailID, request.Filename)
	var isFind string
	if err := row.Scan(&isFind); err != nil {
		return &repository_proto.Nothing{
			Status: false,
		}, nil
	}
	return &repository_proto.Nothing{
		Status: true,
	}, nil
}

func (c *Database) ListAttaches(context context.Context, request *repository_proto.GetAttachRequest) (*repository_proto.ResponseAttaches, error) {
	rows, err := c.Conn.Query(context, "SELECT filename FROM overflow.attaches WHERE mail_id = $1;", request.MailID)
	if err != nil {
		log.Error(err)
		return &repository_proto.ResponseAttaches{
			Filenames: nil,
		}, err
	}
	defer rows.Close()

	var attaches models.AttachList
	for rows.Next() {
		var attachShort models.AttachShort
		if err := rows.Scan(&attachShort.Filename); err != nil {
			log.Error(err)
			return &repository_proto.ResponseAttaches{
				Filenames: nil,
			}, err
		}
		attaches.Attaches = append(attaches.Attaches, attachShort)
	}

	if err := rows.Err(); err != nil {
		log.Error(err)
		return &repository_proto.ResponseAttaches{
			Filenames: nil,
		}, err
	}

	filenamesBytes, err := json.Marshal(attaches)
	if err != nil {
		log.Error(err)
		return &repository_proto.ResponseAttaches{
			Filenames: nil,
		}, err
	}

	return &repository_proto.ResponseAttaches{
		Filenames: filenamesBytes,
	}, nil
}
