package delivery

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"

	"google.golang.org/grpc"
)

type Delivery struct {
	auth auth_proto.AuthClient
	profile profile_proto.ProfileClient
	mailbox mailbox_proto.MailboxClient
	folderManager folder_manager_proto.FolderManagerClient
	config *config.Config
}

func (d *Delivery) Init(
	config *config.Config,
	authDial grpc.ClientConnInterface,
	profileDial grpc.ClientConnInterface,
	mailboxDial grpc.ClientConnInterface,
	folderManagerDial grpc.ClientConnInterface,
	) {
	d.config = config
	d.auth = auth_proto.NewAuthClient(authDial)
	d.profile = profile_proto.NewProfileClient(profileDial)
	d.mailbox = mailbox_proto.NewMailboxClient(mailboxDial)
	d.folderManager = folder_manager_proto.NewFolderManagerClient(folderManagerDial)
}