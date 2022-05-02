package auth

import (
	"OverflowBackend/internal/config"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type AuthServiceInterface interface {
	Init(config *config.Config, db repository_proto.DatabaseRepositoryClient)
	SignIn(context context.Context, data *auth_proto.SignInRequest) (*utils_proto.JsonResponse, error)
	SignUp(context context.Context, data *auth_proto.SignUpRequest) (*utils_proto.JsonResponse, error)
}