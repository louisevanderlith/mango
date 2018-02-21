package comms

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/db"
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
		body := buildMessage(m)
		sendErr := sendEmail(body, m.Name, m.To)

		if sendErr != nil {
			m.Sent = false
			m.Error = sendErr.Error()
		}
	}

	_, err := Ctx.Message.Create(&m)

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

func buildMessage(m Message) string {
	result := ""

	result += fmt.Sprintf("Name: %s <br/>", m.Name)
	result += fmt.Sprintf("Email: %s <br/>", m.Email)
	result += fmt.Sprintf("Phone: %s <br/>", m.Phone)
	result += fmt.Sprintf("Message: %s <br/>", m.Body)

	return result
}
