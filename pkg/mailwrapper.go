package pkg

import (
	"OverflowBackend/internal/models"
	"bufio"
	"strings"
)

var innerMailSeparator = "\n> "

func MailWrapper(mailOuter models.MailForm, mailInner models.Mail) models.MailForm {
	scanner := bufio.NewScanner(strings.NewReader(mailInner.Text))
	for scanner.Scan() {
		mailOuter.Text = (innerMailSeparator + scanner.Text())
	}
	return mailOuter
}