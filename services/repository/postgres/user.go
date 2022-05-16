package postgres

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"

	//log "github.com/sirupsen/logrus"
)

// Получить данные пользователя по его почте
func (c *Database) GetUserInfoByUsername(context context.Context, request *repository_proto.GetUserInfoByUsernameRequest) (response *repository_proto.ResponseUser, err error) {
	var user models.User
	userBytes, _ := json.Marshal(user)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
		}
		response = &repository_proto.ResponseUser{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			User: userBytes,
		}
	}()
	rows, err := c.Conn.Query(context, "Select id, first_name, last_name, password, username from overflow.users where username = $1;", request.Username)
	if err != nil {
		return &repository_proto.ResponseUser{
			User: userBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseUser{
				User: userBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		user.Id = values[0].(int32)
		user.Firstname = values[1].(string)
		user.Lastname = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	userBytes, _ = json.Marshal(user)
	return &repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Получить данные пользователя по его айди в бд
func (c *Database) GetUserInfoById(context context.Context, request *repository_proto.GetUserInfoByIdRequest) (response *repository_proto.ResponseUser, err error) {
	var user models.User
	userBytes, _ := json.Marshal(user)
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errRecover.(error)
		}
		response = &repository_proto.ResponseUser{
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
			User: userBytes,
		}
	}()
	rows, err := c.Conn.Query(context, "Select id, first_name, last_name, password, username from overflow.users(id, first_name, last_name, password, username) where id = $1;", request.UserId)
	if err != nil {
		return &repository_proto.ResponseUser{
			User: userBytes,
			Response: &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			},
		}, err
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return &repository_proto.ResponseUser{
				User: userBytes,
				Response: &utils_proto.DatabaseResponse{
					Status: utils_proto.DatabaseStatus_ERROR,
				},
			}, err
		}
		user.Id = values[0].(int32)
		user.Firstname = values[1].(string)
		user.Lastname = values[2].(string)
		user.Password = values[3].(string)
		user.Username = values[4].(string)
	}
	userBytes, _ = json.Marshal(user)
	return &repository_proto.ResponseUser{
		User: userBytes,
		Response: &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		},
	}, nil
}

// Добавить пользователя
func (c *Database) AddUser(context context.Context, request *repository_proto.AddUserRequest) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	res, err := c.Conn.Query(context, "INSERT INTO overflow.users(first_name, last_name, password, username) VALUES ($1, $2, $3, $4);", user.Firstname, user.Lastname, user.Password, user.Username)
	if err == nil {
		res.Close()
		resp, err := c.GetUserInfoByUsername(context, &repository_proto.GetUserInfoByUsernameRequest{
			Username: user.Username,
		})
		if err != nil {
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		err = json.Unmarshal(resp.User, &user)
		if err != nil{
			return &utils_proto.DatabaseResponse{
				Status: utils_proto.DatabaseStatus_ERROR,
			}, err
		}
		c.UserConfig(context, user.Id) // конфигурация профиля пользователя
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Изменить пароль
func (c *Database) ChangeUserPassword(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.users SET password = $1 WHERE id = $2;", request.Data, user.Id)
	if err == nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Изменить имя
func (c *Database) ChangeUserFirstName(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.users SET first_name = $1 WHERE id = $2;", request.Data, user.Id)
	if err == nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}

// Изменить фамилию
func (c *Database) ChangeUserLastName(context context.Context, request *repository_proto.ChangeForm) (*utils_proto.DatabaseResponse, error) {
	var user models.User
	err := json.Unmarshal(request.User, &user)
	if err != nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
	_, err = c.Conn.Exec(context, "UPDATE overflow.users SET last_name = $1 WHERE id = $2;", request.Data, user.Id)
	if err == nil {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_OK,
		}, err
	} else {
		return &utils_proto.DatabaseResponse{
			Status: utils_proto.DatabaseStatus_ERROR,
		}, err
	}
}