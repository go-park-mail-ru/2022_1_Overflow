package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/validation"
	"OverflowBackend/pkg"
	"fmt"
)

type SessionManager struct {}

func (uc *UseCase) SignIn(data models.SignInForm) pkg.JsonResponse {
	if err := validation.CheckSignIn(data); err != nil {
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error())
	}
	userFind, err := uc.db.GetUserInfoByEmail(data.Email)
	if (err != nil || userFind == models.User{}) {
		return pkg.WRONG_CREDS_ERR
	}
	if (userFind.Password != pkg.HashPassword(data.Password)) {
		return pkg.WRONG_CREDS_ERR
	}
	return pkg.NO_ERR
}

func (uc *UseCase) SignUp(data models.SignUpForm) pkg.JsonResponse {
	if err := validation.CheckSignUp(data); err != nil {
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error())
	}
	user, err := pkg.ConvertToUser(data)
	if err != nil {
		return pkg.INTERNAL_ERR
	}
	
	userFind, _ := uc.db.GetUserInfoByEmail(data.Email)
	if (userFind != models.User{}) {
		return pkg.CreateJsonErr(pkg.STATUS_USER_EXISTS, fmt.Sprintf("Пользователь %v уже существует.", data.Email))
	}
	if err = uc.db.AddUser(user); err != nil {
		return pkg.DB_ERR
	}
	return pkg.NO_ERR
}