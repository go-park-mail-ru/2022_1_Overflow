package pkg

import (
	"OverflowBackend/internal/models"
	"strings"
)

const domain = "overmail.online"

func ConvertToUser(data *models.SignUpForm) (user models.User, err error) {
	user.Firstname = data.Firstname
	user.Lastname = data.Lastname
	user.Username = data.Username
	user.Password = HashPassword(data.Password)
	return
}

func EmailToUsername(email string) string {
	if strings.Contains(email, "@") {
		parts := strings.Split(email, "@")
		if parts[1] == domain {
			return parts[0]
		} else {
			return email
		}
	} else {
		return email
	}
}