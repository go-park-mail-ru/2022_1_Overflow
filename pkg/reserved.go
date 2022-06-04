package pkg

const (
	FOLDER_SPAM = "Спам"
	FOLDER_DRAFTS = "Черновики"
)

const (
	DELETED_USERNAME = "[DELETED]"
)

const (
	THEME_PINK = "pink"
	THEME_ORANGE = "orange"
	THEME_GREEN = "green"
	THEME_BLUE = "blue"
)

var THEMES = [...]string{
	THEME_PINK,
	THEME_ORANGE,
	THEME_GREEN,
	THEME_BLUE,
}

func IsThemeReserved(theme string) bool {
	for _, themeRes := range THEMES {
		if themeRes == theme {
			return true
		}
	}
	return false
}

func IsFolderReserved(folderName string) bool {
	return folderName == FOLDER_SPAM || folderName == FOLDER_DRAFTS
}