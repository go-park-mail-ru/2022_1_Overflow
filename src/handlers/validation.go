package handlers

import (
	"fmt"
	"strings"
)

type Validators struct {}

func (v *Validators) EmailValidator(email string) error {
	if err := v.CheckEmptyField(email, "email"); err != nil {
		return fmt.Errorf("Поле адреса почты является пустым.")
	}
	return nil
}

func (v *Validators) PasswordValidator(password string) error {
	if err := v.CheckEmptyField(password, "password"); err != nil {
		return fmt.Errorf("Поле пароля является пустым.")
	}
	return nil
}

func (v *Validators) SamePasswordValidator(password, passwordConfirmation string) error {
	if err := v.CheckEmptyField(passwordConfirmation, "password_confirmation"); err != nil {
		return fmt.Errorf("Поле повтора пароля является пустым.")
	}
	if password != passwordConfirmation {
		return fmt.Errorf("Поля пароля и повтора пароля не совпадают.")
	}
	return nil
}

func (v *Validators) CheckSignup(data map[string]string) error {
	if err := v.CheckEmptyField(data["first_name"], "Имя"); err != nil {
		return err
	}
	if err := v.CheckEmptyField(data["last_name"], "Фамилия"); err != nil {
		return err
	}
	if err := v.CheckSignin(data["email"], data["password"]); err != nil {
		return err
	}
	if err := v.SamePasswordValidator(data["password"], data["password_confirmation"]); err != nil {
		return err
	}
	return nil
}

func (v *Validators) CheckSignin(email, password string) error {
	if err := v.EmailValidator(email); err != nil {
		return err
	}
	if err := v.PasswordValidator(password); err != nil {
		return err
	}
	return nil
}

func (v *Validators) CheckEmptyField(value string, key string) error {
	if len(strings.TrimSpace(value)) == 0 {
		if len(strings.TrimSpace(key)) == 0 {
			return fmt.Errorf("Поле является пустым.")
		} else {
			return fmt.Errorf("Поле %v является пустым.", key)
		}
	}
	return nil
}