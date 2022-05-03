package profile

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type ProfileServiceInterface interface {
	Init(config *config.Config, db repository_proto.DatabaseRepositoryClient)
	GetInfo(context context.Context, request *profile_proto.GetInfoRequest) (*profile_proto.GetInfoResponse, error)
	SetAvatar(context context.Context, request *profile_proto.SetAvatarRequest) (*utils_proto.JsonResponse, error)
	SetInfo(context context.Context, request *profile_proto.SetInfoRequest) (*utils_proto.JsonResponse, error)
	GetAvatar(context context.Context, request *profile_proto.GetAvatarRequest) (*profile_proto.GetAvatarResponse, error)
	ChangePassword(context context.Context, request *profile_proto.ChangePasswordRequest) (*utils_proto.JsonResponse, error)
}