package general

import (
	"encoding/json"
)

/*
Создать ответ на запрос в формате JSON, где:
	status - статус ответа;
	message - доп. сообщение ответа;
	content - передаваемое содержимое ответа;

	Статусы:
		0 - запрос прошел успешно,
		1 - невозможно обработать данные формы в запросе,
		2 - данные формы не прошли валидацию,
		3 - ошибка при преобразовании данных формы в структуру БД,
		4 - ошибка при записи в БД.
*/
func CreateJsonResponse(status int, message string, content interface{}) []byte {
	resp, _ := json.Marshal(
		map[string]interface{}{
			"status":  status,
			"message": message,
			"content": content,
		},
	)
	return resp
}