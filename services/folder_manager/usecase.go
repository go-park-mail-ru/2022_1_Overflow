package folder_manager

import (
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type FolderManagerServiceInterface interface {
	AddFolder(context context.Context, request *folder_manager_proto.AddFolderRequest) (*utils_proto.JsonResponse, error)
	AddMailToFolder(context context.Context, request *folder_manager_proto.AddMailToFolderRequest) (*utils_proto.JsonResponse, error)
	ChangeFolder(context context.Context, request *folder_manager_proto.ChangeFolderRequest) (*utils_proto.JsonResponse, error)
	//GetFolderByName(context context.Context, request *folder_manager_proto.GetFolderByNameRequest) (*folder_manager_proto.ResponseFolder, error)
	//GetFolderById(context context.Context, request *folder_manager_proto.GetFolderByIdRequest) (*folder_manager_proto.ResponseFolder, error)
	DeleteFolder(context context.Context, request *folder_manager_proto.DeleteFolderRequest) (*utils_proto.JsonResponse, error)
	ListFolders(context context.Context, request *folder_manager_proto.ListFoldersRequest) (*folder_manager_proto.ResponseFolders, error)
	ListFolder(context context.Context, request *folder_manager_proto.ListFolderRequest) (*folder_manager_proto.ListFolderRequest, error)
}