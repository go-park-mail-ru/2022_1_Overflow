package pkg

import (
	"encoding/json"
)

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
