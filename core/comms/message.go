package comms

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"

	"gopkg.in/gomail.v2"
)

type Message struct {
	Name  string `hsk:"size(50)"`
	Email string `hsk:"size(128)"`
	Phone string `hsk:"size(15)"`
	Body  string `hsk:"size(1024)"`
	To    string `hsk:"null;size(128)"`
	Sent  bool   `hsk:"default(false)"`
	Error string `hsk:"null;size(2048)"`
}

func (m Message) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}

func GetMessages(page, size int) (messageSet, error) {
	return ctx.Messages.Find(page, size, func(obj Message) bool {
		return true
	})
}

func (m Message) SendMessage() error {
	if beego.BConfig.RunMode != "dev" {
		body := populatTemplate(m)
		sendErr := sendEmail(body, m.Name, m.To)

		if sendErr != nil {
			m.Sent = false
			m.Error = sendErr.Error()
		} else {
			m.Sent = true
		}
	}

	_, err := ctx.Messages.Create(m)

	return err
}

func sendEmail(body, name, to string) error {
	smtpUser := beego.AppConfig.String("smtpUsername")
	smtpPass := beego.AppConfig.String("smtpPassword")
	smtpAddress := beego.AppConfig.String("smtpAddress")
	smtpPort, _ := beego.AppConfig.Int("smtpPort")

	gm := gomail.NewMessage()
	gm.SetHeader("From", smtpUser)
	gm.SetHeader("To", to)
	gm.SetHeader("Subject", "Contact Us - "+name)
	gm.SetBody("text/html", body)

	d := gomail.NewDialer(smtpAddress, smtpPort, smtpUser, smtpPass)

	err := d.DialAndSend(gm)

	if err != nil {
		log.Println("sendMail:", err)
	}

	return err
}
