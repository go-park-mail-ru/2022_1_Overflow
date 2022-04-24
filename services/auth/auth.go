package auth

import (
	"OverflowBackend/internal/validation"
	"OverflowBackend/pkg"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	db repository_proto.DatabaseRepositoryClient
}

func (s *AuthService) Init(db repository_proto.DatabaseRepositoryClient) {
	s.db = db
}

func (s *AuthService) SignIn(data *auth_proto.SignInForm) *utils_proto.JsonResponse {
	log.Info("SignIn: ", "handling usecase")
	log.Info("SignIn: ", "handling validation")
	if err := validation.CheckSignIn(data); err != nil {
		log.Errorf("SignIn: %v", err)
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error())
	}
	log.Info("SignIn: ", "handling db")
	req := repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context.Background(), &req)
	if (err != nil) {
		log.Errorf("SignIn: %v", err)
		return &pkg.WRONG_CREDS_ERR
	}
	userFind := resp.User
	if (proto.Equal(userFind, &utils_proto.User{})) {
		return &pkg.WRONG_CREDS_ERR
	}
	if userFind.Password != pkg.HashPassword(data.Password) {
		log.Errorf("SignIn: %v", err)
		return &pkg.WRONG_CREDS_ERR
	}
	log.Info("SignIn, username: ", data.Username)
	return &pkg.NO_ERR
}

func (s *AuthService) SignUp(data *auth_proto.SignUpForm) *utils_proto.JsonResponse {
	log.Info("SignUp: ", "handling usecase")
	log.Info("SignUp: ", "handling validaton")
	if err := validation.CheckSignUp(data); err != nil {
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error())
	}
	user, err := pkg.ConvertToUser(data)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &pkg.INTERNAL_ERR
	}
	req := repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context.Background(), &req)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &pkg.INTERNAL_ERR
	}
	userFind := resp.User
	if (proto.Equal(userFind, &utils_proto.User{})) {
		return pkg.CreateJsonErr(pkg.STATUS_USER_EXISTS, fmt.Sprintf("Пользователь %v уже существует.", data.Username))
	}
	req2 := repository_proto.AddUserRequest{
		User: &user,
	}
	resp2, err := s.db.AddUser(context.Background(), &req2); 
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &pkg.DB_ERR
	}
	if resp2.Status == 1 {
		return &pkg.DB_ERR
	}
	log.Info("SignUp, username: ", data.Username)
	return &pkg.NO_ERR
}
