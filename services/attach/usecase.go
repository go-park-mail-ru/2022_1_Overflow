package attach

import (
	"OverflowBackend/proto/attach_proto"
	"context"
)

type ProfileServiceInterface interface {
	SaveAttach(context context.Context, request *attach_proto.SaveAttachRequest) (*attach_proto.Nothing, error)
	GetAttach(context context.Context, request *attach_proto.GetAttachRequest) (*attach_proto.AttachResponse, error)
	ListAttaches(context context.Context, request *attach_proto.GetAttachRequest) (*attach_proto.AttachListResponse, error)
}
