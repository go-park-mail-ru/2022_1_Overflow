package auth

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/validation"
	"OverflowBackend/pkg"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	config *config.Config
	db     repository_proto.DatabaseRepositoryClient
}

func (s *AuthService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	s.config = config
	s.db = db
}

func (s *AuthService) SignIn(context context.Context, request *auth_proto.SignInRequest) (*utils_proto.JsonResponse, error) {
	log.Info("SignIn: ", "handling usecase")
	log.Info("SignIn: ", "handling validation")
	var data models.SignInForm
	err := json.Unmarshal(request.Form, &data)
	if err != nil {
		log.Errorf("SignIn: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if err := validation.CheckSignIn(&data); err != nil {
		log.Errorf("SignIn: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()).Bytes(),
		}, err
	}
	log.Info("SignIn: ", "handling db")
	req := repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context, &req)
	if err != nil {
		log.Errorf("SignIn: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(),
		}, err
	}
	var userFind models.User
	userFindBytes := resp.User
	err = json.Unmarshal(userFindBytes, &userFind)
	if (userFind == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(),
		}, nil
	}
	if userFind.Password != pkg.HashPassword(data.Password) {
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(),
		}, nil
	}
	log.Info("SignIn, username: ", data.Username)
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, nil
}

func (s *AuthService) SignUp(context context.Context, request *auth_proto.SignUpRequest) (*utils_proto.JsonResponse, error) {
	log.Info("SignUp: ", "handling usecase")
	log.Info("SignUp: ", "handling validaton")
	var data models.SignUpForm
	err := json.Unmarshal(request.Form, &data)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if err := validation.CheckSignUp(&data); err != nil {
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()).Bytes(),
		}, err
	}
	user, err := pkg.ConvertToUser(&data)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.INTERNAL_ERR.Bytes(),
		}, err
	}
	req := repository_proto.GetUserInfoByUsernameRequest{
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context, &req)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	var userFind models.User
	userFindBytes := resp.User
	err = json.Unmarshal(userFindBytes, &userFind)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (userFind == models.User{}) {
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(),
		}, nil
	}
	userBytes, _ := json.Marshal(user)
	req2 := repository_proto.AddUserRequest{
		User: userBytes,
	}
	resp2, err := s.db.AddUser(context, &req2)
	if err != nil {
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK {
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	log.Info("SignUp, username: ", data.Username)
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, err
}
