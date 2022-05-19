package pkg

const (
	FOLDER_SPAM = "Спам"
	FOLDER_DRAFTS = "Черновики"
)

const (
	DELETED_USERNAME = "[DELETED]"
)

func IsFolderReserved(folderName string) bool {
	return folderName == FOLDER_SPAM || folderName == FOLDER_DRAFTS
}