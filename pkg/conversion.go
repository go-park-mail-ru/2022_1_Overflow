package pkg

import "OverflowBackend/internal/models"

func ConvertToUser(data models.SignUpForm) (user models.User, err error) {
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.Password = HashPassword(data.Password)
	return
}