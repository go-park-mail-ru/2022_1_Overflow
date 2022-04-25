package delivery

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"

	"google.golang.org/grpc"
)

type Delivery struct {
	auth auth_proto.AuthClient
	profile profile_proto.ProfileClient
	mailbox mailbox_proto.MailboxClient
	config *config.Config
}

func (d *Delivery) Init(config *config.Config, authDial grpc.ClientConnInterface, profileDial grpc.ClientConnInterface, mailboxDial grpc.ClientConnInterface) {
	d.auth = auth_proto.NewAuthClient(authDial)
	d.profile = profile_proto.NewProfileClient(profileDial)
	d.mailbox = mailbox_proto.NewMailboxClient(mailboxDial)
	d.config = config
}