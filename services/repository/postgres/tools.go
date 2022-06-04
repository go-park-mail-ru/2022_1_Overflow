package postgres

import (
	"OverflowBackend/pkg"
	"context"
	//log "github.com/sirupsen/logrus"
)

func (c *Database) UserConfig(context context.Context, user_id int32) error {
	rows, err := c.Conn.Query(context, "INSERT INTO overflow.folders(name, user_id) VALUES ($1, $2);", pkg.FOLDER_SPAM, user_id)
	if err != nil {
		return err
	}
	rows.Close()
	rows, err = c.Conn.Query(context, "INSERT INTO overflow.folders(name, user_id) VALUES ($1, $2);", pkg.FOLDER_DRAFTS, user_id)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}