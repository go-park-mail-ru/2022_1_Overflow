package pkg

import (
	"errors"
	"strings"
)

func ValidateFormatFolderName(folderName string) (string, error) {
	if IsEmpty(folderName) {
		return "", errors.New("Имя папки не может быть пустым.")
	}
	folderName = strings.TrimSpace(folderName)
	return folderName, nil
}

func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}