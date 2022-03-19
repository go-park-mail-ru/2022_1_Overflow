package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/validation"
	"OverflowBackend/pkg"
	"fmt"
)

type SessionManager struct {}

func (uc *UseCase) SignIn(data models.SignInForm) error {
	if err := validation.CheckSignIn(data); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) SignUp(data models.SignUpForm) error {
	if err := validation.CheckSignUp(data); err != nil {
		return err
		//w.Write(pkg.CreateJsonResponse(1, err.Error(), nil))
	}
	user, err := pkg.ConvertToUser(data)
	if err != nil {
		return err
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	userFind, _ := uc.db.GetUserInfoByEmail(data.Email)
	if (userFind != models.User{}) {
		err = fmt.Errorf("Пользователь %v уже существует.", data.Email)
		return err
		//w.Write(pkg.CreateJsonResponse(2, err.Error(), nil))
	}
	if err = uc.db.AddUser(user); err != nil {
		return err
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}