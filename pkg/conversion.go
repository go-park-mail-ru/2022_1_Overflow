package pkg

import (
	"OverflowBackend/internal/models"
	"errors"
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

func IsLocalEmail(email string) bool {
	if strings.Contains(email, "@") {
		parts := strings.Split(email, "@")
		return parts[1] == domain
	} else {
		return true
	}
}

func ParseDomain(email string) (string, error) {
	if strings.Contains(email, "@") {
		return strings.Split(email, "@")[1], nil
	} else {
		return "", errors.New("Не удалось распознать домен почтового адреса.")
	}
}

func ConvertDomain(domain string) string {
	if domain == "gmail.com" {
		return "smtp.gmail.com:587"
	} else {
		return domain + ":25"
	}
}

func ThemeToAvatarName(theme string) string {
	if IsThemeReserved(theme) {
		return "dummy_" + theme
	} else {
		return "dummy_blue"
	}
}