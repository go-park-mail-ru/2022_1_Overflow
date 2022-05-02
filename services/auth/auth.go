package auth

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/validation"
	"OverflowBackend/pkg"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type AuthService struct {
	config *config.Config
	db repository_proto.DatabaseRepositoryClient
}

func (s *AuthService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	s.config = config
	s.db = db
}

func (s *AuthService) SignIn(context context.Context, data *auth_proto.SignInForm) (*utils_proto.JsonResponse, error) {
	log.Info("SignIn: ", "handling usecase")
	log.Info("SignIn: ", "handling validation")
	if err := validation.CheckSignIn(data); err != nil {
		log.Errorf("SignIn: %v", err)
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()), err
	}
	log.Info("SignIn: ", "handling db")
	req := repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context, &req)
	if (err != nil) {
		log.Errorf("SignIn: %v", err)
		return &pkg.WRONG_CREDS_ERR, err
	}
	userFind := resp.User
	if (proto.Equal(userFind, &utils_proto.User{})) {
		return &pkg.WRONG_CREDS_ERR, nil
	}
	if userFind.Password != pkg.HashPassword(data.Password) {
		return &pkg.WRONG_CREDS_ERR, nil
	}
	log.Info("SignIn, username: ", data.Username)
	return &pkg.NO_ERR, nil
}

func (s *AuthService) SignUp(context context.Context, data *auth_proto.SignUpForm) (*utils_proto.JsonResponse, error) {
	log.Info("SignUp: ", "handling usecase")
	log.Info("SignUp: ", "handling validaton")
	if err := validation.CheckSignUp(data); err != nil {
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()), err
	}
	user, err := pkg.ConvertToUser(data)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &pkg.INTERNAL_ERR, err
	}
	req := repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context, &req)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &pkg.INTERNAL_ERR, err
	}
	userFind := resp.User
	if (!proto.Equal(userFind, &utils_proto.User{})) {
		return pkg.CreateJsonErr(pkg.STATUS_USER_EXISTS, fmt.Sprintf("Пользователь %v уже существует.", data.Username)), nil
	}
	req2 := repository_proto.AddUserRequest{
		User: &user,
	}
	resp2, err := s.db.AddUser(context, &req2); 
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &pkg.DB_ERR, err
	}
	if resp2.Status == 1 {
		return &pkg.DB_ERR, nil
	}
	log.Info("SignUp, username: ", data.Username)
	return &pkg.NO_ERR, nil
}
