package comms

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
	"gopkg.in/gomail.v2"
)

type Message struct {
	db.Record
	Name  string `orm:"size(50)"`
	Email string `orm:"size(128)"`
	Phone string `orm:"size(15)"`
	Body  string `orm:"size(1024)"`
	To    string `orm:"null;size(128)"`
	Sent  bool   `orm:"default(false)"`
	Error string `orm:"null;size(2048)"`
}

func (m Message) Validate() (bool, error) {
	return util.ValidateStruct(&m)
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

	_, err := Ctx.Messages.Create(&m)

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
