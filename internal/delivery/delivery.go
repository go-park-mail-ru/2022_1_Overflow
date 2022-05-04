package delivery

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
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
	auth auth_proto.AuthClient,
	profile profile_proto.ProfileClient,
	mailbox mailbox_proto.MailboxClient,
	folderManager folder_manager_proto.FolderManagerClient,
	) {
	d.config = config
	d.auth = auth
	d.profile = profile
	d.mailbox = mailbox
	d.folderManager = folderManager
}