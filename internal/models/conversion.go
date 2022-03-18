package models

import "OverflowBackend/pkg"

func ConvertToUser(data SignUpForm) (user User, err error) {
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.Password = pkg.HashPassword(data.Password)
	return
}
