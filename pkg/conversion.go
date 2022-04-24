package pkg

import (
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/utils_proto"
)

func ConvertToUser(data *auth_proto.SignUpForm) (user utils_proto.User, err error) {
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Username = data.Username
	user.Password = HashPassword(data.Password)
	return
}
