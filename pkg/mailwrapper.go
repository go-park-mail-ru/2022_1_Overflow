package pkg

import (
	"OverflowBackend/internal/models"
	"bufio"
	"strings"
)

var innerMailSeparator = ">"

func MailWrapper(mailOuter models.MailForm, mailInner models.Mail) models.MailForm {
	scanner := bufio.NewScanner(strings.NewReader(mailInner.Text))
	for scanner.Scan() {
		mailOuter.Text = ("\n" + innerMailSeparator + scanner.Text())
	}
	return mailOuter
}