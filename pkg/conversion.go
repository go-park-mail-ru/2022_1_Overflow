package pkg

import "OverflowBackend/internal/models"

func ConvertToUser(data models.SignUpForm) (user models.User, err error) {
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Username = data.Username
	user.Password = HashPassword(data.Password)
	return
}
