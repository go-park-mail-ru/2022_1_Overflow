package usecase

import (
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/usecase/session"
	"OverflowBackend/internal/usecase/validation"
	"OverflowBackend/pkg"
	"fmt"
	"net/http"
)

type SessionManager struct {}

func (uc *UseCase) SignIn(w http.ResponseWriter, r *http.Request, data models.SignInForm) {
	if err := validation.CheckSignIn(data); err != nil {
		w.Write(pkg.CreateJsonResponse(1, err.Error(), nil))
		return
	}

	err := session.CreateSession(w, r, data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(pkg.CreateJsonResponse(0, "OK", nil))
}

func (uc *UseCase) SignUp(w http.ResponseWriter, r *http.Request, data models.SignUpForm) {
	if err := validation.CheckSignUp(data); err != nil {
		w.Write(pkg.CreateJsonResponse(1, err.Error(), nil))
		return
	}
	user, err := pkg.ConvertToUser(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	userFind, _ := uc.db.GetUserInfoByEmail(data.Email)
	if (userFind != models.User{}) {
		err = fmt.Errorf("Пользователь %v уже существует.", data.Email)
		w.Write(pkg.CreateJsonResponse(2, err.Error(), nil))
		return
	}
	if err = uc.db.AddUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Write(pkg.CreateJsonResponse(0, "OK", nil))
}