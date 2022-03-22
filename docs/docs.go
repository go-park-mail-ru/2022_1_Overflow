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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/logout": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "Завершение сессии пользователя",
                "responses": {
                    "200": {
                        "description": "Успешное завершение сессии.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Сессия отсутствует, сессия не валидна."
                    },
                    "500": {
                        "description": ""
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
                        "description": "OK"
                    },
                    "401": {
                        "description": "Письмо не принадлежит пользователю, сессия отсутствует или сессия не валидна."
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка БД, неверные GET параметры."
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
                        "description": "Сессия отсутствует или сессия не валидна."
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка БД."
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
                        "description": "Сессия отсутствует или сессия не валидна."
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка БД."
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
                        "description": "OK"
                    },
                    "401": {
                        "description": "Письмо не принадлежит пользователю, сессия отсутствует или сессия не валидна."
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка БД, неверные GET параметры."
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
                        "description": "Сессия отсутствует, сессия не валидна."
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка БД, пользователь не найден, неверные данные сессии."
                    }
                }
            }
        },
        "/profile/avatar": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "Получение ссылки на аватарку пользователя",
                "responses": {
                    "200": {
                        "description": "Ссылка на аватарку в формате /{static_dir}/{file}.{ext}.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка БД, пользователь не найден или сессия не валидна."
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
                    "text/plain"
                ],
                "summary": "Установка/смена аватарки пользователя",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Файл аватарки.",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное установка аватарки.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка валидации формы, БД или сессия не валидна."
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
                    "text/plain"
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное изменение настроек.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": "Ошибка валидации формы, БД или сессия не валидна."
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
                    "text/plain"
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная аутентификация пользователя.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Пользователь не существует, ошибка БД или валидации формы."
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
                    "text/plain"
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная регистрация пользователя.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка валидации формы, БД или пользователь уже существует."
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
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
