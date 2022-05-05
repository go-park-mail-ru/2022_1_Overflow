package pkg

import (
	"OverflowBackend/internal/models"
)

func ConvertToUser(data *models.SignUpForm) (user models.User, err error) {
	user.Firstname = data.Firstname
	user.Lastname = data.Lastname
	user.Username = data.Username
	user.Password = HashPassword(data.Password)
	return
}
