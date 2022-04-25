package auth

import (
	"OverflowBackend/internal/config"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
)

type AuthServiceInterface interface {
	Init(config *config.Config, db repository_proto.DatabaseRepositoryClient)
	SignIn(data *auth_proto.SignInForm) *utils_proto.JsonResponse
	SignUp(data *auth_proto.SignUpForm) *utils_proto.JsonResponse
}