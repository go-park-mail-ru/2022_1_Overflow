package handlers

import (
	"fmt"
	"strings"
)

type Validators struct {
}

func (v *Validators) EmailValidator(email string) error {
	if err := v.CheckEmptyField(email); err != nil {
		return fmt.Errorf("Поле адреса почты является пустым.")
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

func (v *Validators) CheckAll(email, password, passwordConfirmation string) error {
	if err := v.EmailValidator(email); err != nil {
		return err
	}
	if err := v.PasswordValidator(password); err != nil {
		return err
	}
	if err := v.SamePasswordValidator(password, passwordConfirmation); err != nil {
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