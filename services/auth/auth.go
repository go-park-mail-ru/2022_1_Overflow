package auth

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/internal/validation"
	"OverflowBackend/pkg"
	auth_proto "OverflowBackend/proto/auth_proto"
	repository_proto "OverflowBackend/proto/repository_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"github.com/mailru/easyjson"

	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	config *config.Config
	db     repository_proto.DatabaseRepositoryClient
}

func (s *AuthService) Init(config *config.Config, db repository_proto.DatabaseRepositoryClient) {
	s.config = config
	s.db = db
}

func (s *AuthService) SignIn(context context.Context, request *auth_proto.SignInRequest) (*utils_proto.JsonResponse, error) {
	log.Info("SignIn: ", "handling usecase")
	log.Info("SignIn: ", "handling validation")
	var data models.SignInForm                     // создаем пустую форму для входа
	err := easyjson.Unmarshal(request.Form, &data) // распаковываем запрос в форму
	if err != nil {                                // ошибка распаковки
		log.Errorf("SignIn: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if err := validation.CheckSignIn(&data); err != nil { // валидируем форму
		log.Errorf("SignIn: %v", err) // ошибка валидации
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()).Bytes(),
		}, err
	}
	// получаем информацию о пользователе по его логину внутри объекта формы
	log.Info("SignIn: ", "handling db")
	req := repository_proto.GetUserInfoByUsernameRequest{ // форма запроса в микросервис БД
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context, &req) // выполняем запрос
	if err != nil {
		log.Errorf("SignIn: %v", err) // ошибка выполнения запроса
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(),
		}, err
	}
	var userFind models.User                           // создаем пустой объект найденного пользователя
	userFindBytes := resp.User                         // объект пользователя в байтовом представлении
	err = easyjson.Unmarshal(userFindBytes, &userFind) // распаковываем объект пользователя
	if err != nil {                                    // ошибка распаковки
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, nil
	}
	if (userFind == models.User{}) { // если объект всё ещё пустой (пользователь не найден)
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(), // возвращаем ответ о неверной паре логин/пароль
		}, nil
	}
	if userFind.Password != pkg.HashPassword(data.Password) { // если объект не пустой, но пароли не совпадают
		return &utils_proto.JsonResponse{
			Response: pkg.WRONG_CREDS_ERR.Bytes(), // возвращаем ответ о неверной паре логин/пароль
		}, nil
	}
	log.Info("SignIn, username: ", data.Username)
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(), // успешный вход, возвращаем ответ об отсуствии ошибки (сессия пока не создана)
	}, nil
}

func (s *AuthService) SignUp(context context.Context, request *auth_proto.SignUpRequest) (*utils_proto.JsonResponse, error) {
	log.Info("SignUp: ", "handling usecase")
	log.Info("SignUp: ", "handling validaton")
	var data models.SignUpForm                     // создаем пустую форму регистрации пользователя
	err := easyjson.Unmarshal(request.Form, &data) // распаковываем данные запроса в форму
	if err != nil {                                // ошибка распаковки
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if err := validation.CheckSignUp(&data); err != nil { // валидируем полученную форму
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_BAD_VALIDATION, err.Error()).Bytes(), // ошибка валидации
		}, err
	}
	user, err := pkg.ConvertToUser(&data) // преобразуем полученную форму в объект пользователя
	if err != nil {                       // ошибка преобразования
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.INTERNAL_ERR.Bytes(),
		}, err
	}
	// получаем информацию о пользователе по его логину
	req := repository_proto.GetUserInfoByUsernameRequest{ // форма запроса в БД
		Username: data.Username,
	}
	resp, err := s.db.GetUserInfoByUsername(context, &req) // выполняем запрос
	if err != nil {                                        // ошибка выполнения запроса
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	var userFind models.User                           // создаем пустой объект найденного пользователя
	userFindBytes := resp.User                         // объект пользователя в байтовом представлении
	err = easyjson.Unmarshal(userFindBytes, &userFind) // распаковываем байтовое представление в объект
	if err != nil {                                    // ошибка распаковки
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.JSON_ERR.Bytes(),
		}, err
	}
	if (userFind != models.User{}) { // если объект найденного пользователя не пустой (т.е. пользователь уже существует)
		return &utils_proto.JsonResponse{
			Response: pkg.CreateJsonErr(pkg.STATUS_USER_EXISTS, "Пользователь уже существует.").Bytes(),
		}, nil
	}
	userBytes, _ := easyjson.Marshal(user)   // преобразуем объект пользователя в байтовое представление
	req2 := repository_proto.AddUserRequest{ // форма запроса добавления пользователя в БД
		User: userBytes,
	}
	resp2, err := s.db.AddUser(context, &req2) // добавляем пользователя в БД
	if err != nil {                            // ошибка выполнения запроса
		log.Errorf("SignUp: %v", err)
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, err
	}
	if resp2.Status != utils_proto.DatabaseStatus_OK { // ошибка выполнения запроса
		return &utils_proto.JsonResponse{
			Response: pkg.DB_ERR.Bytes(),
		}, nil
	}
	log.Info("SignUp, username: ", data.Username)
	return &utils_proto.JsonResponse{
		Response: pkg.NO_ERR.Bytes(),
	}, err
}
