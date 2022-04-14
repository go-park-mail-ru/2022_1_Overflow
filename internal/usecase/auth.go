package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/validation"
	"OverflowBackend/pkg"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type SessionManager struct{}

func (uc *UseCase) SignIn(data models.SignInForm) pkg.JsonResponse {
	log.Info("SignIn: ", "handling usecase")
	log.Info("SignIn: ", "handling validation")
	if err := validation.CheckSignIn(data); err != nil {
		log.Error("SignIn: ", "bad validation")
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error())
	}
	log.Info("SignIn: ", "handling db")
	userFind, err := uc.db.GetUserInfoByUsername(data.Username)
	if (err != nil || userFind == models.User{}) {
		log.Error("SignIn: ", "wrong username")
		return pkg.WRONG_CREDS_ERR
	}
	if userFind.Password != pkg.HashPassword(data.Password) {
		log.Error("SignIn: ", "wrong password")
		return pkg.WRONG_CREDS_ERR
	}
	log.Info("SignIn, username: ", data.Username)
	return pkg.NO_ERR
}

func (uc *UseCase) SignUp(data models.SignUpForm) pkg.JsonResponse {
	log.Info("SignUp: ", "handling usecase")
	if err := validation.CheckSignUp(data); err != nil {
		return pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error())
	}
	user, err := pkg.ConvertToUser(data)
	if err != nil {
		log.Error(err)
		return pkg.INTERNAL_ERR
	}

	userFind, _ := uc.db.GetUserInfoByUsername(data.Username)
	if (userFind != models.User{}) {
		return pkg.CreateJsonErr(pkg.STATUS_USER_EXISTS, fmt.Sprintf("Пользователь %v уже существует.", data.Username))
	}
	if err = uc.db.AddUser(user); err != nil {
		log.Error(err)
		return pkg.DB_ERR
	}
	log.Info("SignUp, username: ", data.Username)
	return pkg.NO_ERR
}
