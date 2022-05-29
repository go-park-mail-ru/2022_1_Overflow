package folder_manager

import (
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
)

type FolderManagerServiceInterface interface {
	AddFolder(context context.Context, request *folder_manager_proto.AddFolderRequest) (*folder_manager_proto.ResponseFolder, error)
	AddMailToFolderById(context context.Context, request *folder_manager_proto.AddMailToFolderByIdRequest) (*utils_proto.JsonResponse, error)
	AddMailToFolderByObject(context context.Context, request *folder_manager_proto.AddMailToFolderByObjectRequest) (*utils_proto.JsonResponse, error)
	ChangeFolder(context context.Context, request *folder_manager_proto.ChangeFolderRequest) (*utils_proto.JsonResponse, error)
	//GetFolderByName(context context.Context, request *folder_manager_proto.GetFolderByNameRequest) (*folder_manager_proto.ResponseFolder, error)
	//GetFolderById(context context.Context, request *folder_manager_proto.GetFolderByIdRequest) (*folder_manager_proto.ResponseFolder, error)
	DeleteFolder(context context.Context, request *folder_manager_proto.DeleteFolderRequest) (*utils_proto.JsonResponse, error)
	ListFolders(context context.Context, request *folder_manager_proto.ListFoldersRequest) (*folder_manager_proto.ResponseFolders, error)
	ListFolder(context context.Context, request *folder_manager_proto.ListFolderRequest) (*folder_manager_proto.ListFolderRequest, error)
	DeleteFolderMail(context context.Context, request *folder_manager_proto.DeleteFolderMailRequest) (*utils_proto.JsonResponse, error)
	MoveFolderMail(context context.Context, request *folder_manager_proto.MoveFolderMailRequest) (*utils_proto.JsonResponse, error)
	UpdateFolderMail(context context.Context, request *folder_manager_proto.UpdateFolderMailRequest) (*utils_proto.JsonResponse, error)
}