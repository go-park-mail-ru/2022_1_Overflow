package validation

import (
	"fmt"
	"strings"
)


func EmailValidator(email string) error {
	if err := CheckEmptyField(email, "email"); err != nil {
		return fmt.Errorf("Поле адреса почты является пустым.")
	}
	if strings.Contains(email, "@") {
		return fmt.Errorf("Поле почты содержит символ @.")
	}
	return nil
}

func PasswordValidator(password string) error {
	if err := CheckEmptyField(password, "password"); err != nil {
		return fmt.Errorf("Поле пароля является пустым.")
	}
	return nil
}

func SamePasswordValidator(password, passwordConfirmation string) error {
	if err := CheckEmptyField(passwordConfirmation, "password_confirmation"); err != nil {
		return fmt.Errorf("Поле повтора пароля является пустым.")
	}
	if password != passwordConfirmation {
		return fmt.Errorf("Поля пароля и повтора пароля не совпадают.")
	}
	return nil
}

func CheckSignup(data map[string]string) error {
	if err := CheckEmptyField(data["first_name"], "Имя"); err != nil {
		return err
	}
	if err := CheckEmptyField(data["last_name"], "Фамилия"); err != nil {
		return err
	}
	if err := CheckSignin(data["email"], data["password"]); err != nil {
		return err
	}
	if err := SamePasswordValidator(data["password"], data["password_confirmation"]); err != nil {
		return err
	}
	return nil
}

func CheckSignin(email, password string) error {
	if err := EmailValidator(email); err != nil {
		return err
	}
	if err := PasswordValidator(password); err != nil {
		return err
	}
	return nil
}

func CheckEmptyField(value string, key string) error {
	if len(strings.TrimSpace(value)) == 0 {
		if len(strings.TrimSpace(key)) == 0 {
			return fmt.Errorf("Поле является пустым.")
		} else {
			return fmt.Errorf("Поле %v является пустым.", key)
		}
	}
	return nil
}