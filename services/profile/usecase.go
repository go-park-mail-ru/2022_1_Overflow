package profile

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
)

type ProfileServiceInterface interface {
	Init(config *config.Config, db repository_proto.DatabaseRepositoryClient)
	GetInfo(data *utils_proto.Session) *profile_proto.GetInfoResponse
	SetAvatar(request *profile_proto.SetAvatarRequest) *utils_proto.JsonResponse
	SetInfo(request *profile_proto.SetInfoRequest) *utils_proto.JsonResponse
	GetAvatar(request *profile_proto.GetAvatarRequest) *profile_proto.GetAvatarResponse
}