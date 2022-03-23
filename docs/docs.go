// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Роман Медников",
            "url": "https://vk.com/l____l____l____l____l____l",
            "email": "jellybe@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/get_token": {
            "get": {
                "description": "Токен приходит в header ответа в поле X-CSRF-Token",
                "produces": [
                    "application/json"
                ],
                "summary": "СSRF проверка",
                "responses": {
                    "200": {
                        "description": "Успешное получение CSRF токена.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Завершение сессии пользователя",
                "responses": {
                    "200": {
                        "description": "Успешное завершение сессии.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует, сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/mail/delete": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Удалить письмо по его id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID запрашиваемого письма.",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/mail/income": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Получение входящих сообщений",
                "responses": {
                    "200": {
                        "description": "Список входящих писем",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Mail"
                            }
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка БД.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/mail/outcome": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Получение исходящих сообщений",
                "responses": {
                    "200": {
                        "description": "Список исходящих писем",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Mail"
                            }
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка БД.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/mail/read": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Прочитать письмо по его id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID запрашиваемого письма.",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Письмо не принадлежит пользователю, ошибка БД, неверные GET параметры.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/mail/send": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Выполняет отправку письма получателю",
                "parameters": [
                    {
                        "description": "Форма письма",
                        "name": "MailForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MailForm"
                        }
                    },
                    {
                        "type": "string",
                        "description": "CSRF токен",
                        "name": "X-CSRF-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная отправка письма.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Получатель не существует, ошибка БД.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Получение данных пользователя",
                "responses": {
                    "200": {
                        "description": "Информация о пользователе",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует, сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка БД, пользователь не найден, неверные данные сессии.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/profile/avatar": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Получение ссылки на аватарку пользователя",
                "responses": {
                    "200": {
                        "description": "Ссылка на аватарку в формате /{static_dir}/{file}.{ext}.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка БД, пользователь не найден или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/profile/avatar/set": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Установка/смена аватарки пользователя",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Файл аватарки.",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CSRF токен",
                        "name": "X-CSRF-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное установка аватарки.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка валидации формы, БД или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/profile/set": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Изменение настроек пользователя",
                "parameters": [
                    {
                        "description": "Форма настроек пользователя.",
                        "name": "SettingsForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SettingsForm"
                        }
                    },
                    {
                        "type": "string",
                        "description": "CSRF токен",
                        "name": "X-CSRF-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное изменение настроек.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка валидации формы, БД или сессия не валидна.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Выполняет аутентификацию и выставляет сессионый cookie с названием OverflowMail",
                "parameters": [
                    {
                        "description": "Форма входа пользователя",
                        "name": "SignInForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInForm"
                        }
                    },
                    {
                        "type": "string",
                        "description": "CSRF токен",
                        "name": "X-CSRF-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная аутентификация пользователя.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Пользователь не существует, ошибка БД или валидации формы.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Выполняет регистрацию пользователя, НЕ выставляет сессионый cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Выполняет регистрацию пользователя",
                "parameters": [
                    {
                        "description": "Форма регистрации пользователя",
                        "name": "SignUpForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpForm"
                        }
                    },
                    {
                        "type": "string",
                        "description": "CSRF токен",
                        "name": "X-CSRF-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Вход уже выполнен, либо успешная регистрация пользователя.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка валидации формы, БД или пользователь уже существует.",
                        "schema": {
                            "$ref": "#/definitions/pkg.JsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Mail": {
            "description": "Структура письма",
            "type": "object",
            "properties": {
                "addressee": {
                    "type": "string"
                },
                "client_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "files": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "read": {
                    "type": "boolean"
                },
                "sender": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "models.MailForm": {
            "type": "object",
            "properties": {
                "addressee": {
                    "type": "string"
                },
                "files": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "models.SettingsForm": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.SignInForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 20
                },
                "password": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "models.SignUpForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 20
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 20
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 20
                },
                "password": {
                    "type": "string",
                    "maxLength": 20
                },
                "password_confirmation": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "models.User": {
            "description": "Структура пользователя",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "pkg.JsonResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "OverMail API",
	Description:      "API почтового сервиса команды Overflow.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
