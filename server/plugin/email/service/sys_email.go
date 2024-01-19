package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
)

type EmailService struct{}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: Send email test
//@return: err error

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: Send email test
//@return: err error
//@params to string recipient
//@params subject string title (topic)
//@params body string email content

func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
