package handlers

import (
	"fmt"
	"strings"
)

type Validators struct {}

func (v *Validators) EmailValidator(email string) error {
	if err := v.CheckEmptyField(email); err != nil {
		return fmt.Errorf("Поле адреса почты является пустым.")
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("Поле адреса почты не содержит символа @.")
	}
	return nil
}

func (v *Validators) PasswordValidator(password string) error {
	if err := v.CheckEmptyField(password); err != nil {
		return fmt.Errorf("Поле пароля является пустым.")
	}
	return nil
}

func (v *Validators) SamePasswordValidator(password, passwordConfirmation string) error {
	if err := v.CheckEmptyField(passwordConfirmation); err != nil {
		return fmt.Errorf("Поле повтора пароля является пустым.")
	}
	return nil
}

func (v *Validators) CheckSignup(email, password, passwordConfirmation string) error {
	if err := v.CheckSignin(email, password); err != nil {
		return err
	}
	if err := v.SamePasswordValidator(password, passwordConfirmation); err != nil {
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

func (v *Validators) CheckEmptyField(value string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("Поле является пустым.")
	}
	return nil
}