package comms

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	gomail "gopkg.in/gomail.v2"
)

type Message struct {
	Record
	Name  string `orm:"size(50)"`
	Email string `orm:"size(128)"`
	Phone string `orm:"size(15)"`
	Body  string `orm:"size(1024)"`
	Sent  bool   `orm:"default(false)"`
	Error string `orm:"null;size(2048)"`
}

func SendMessage(m Message) error {
	//body := buildMessage(m)
	//sendErr := sendEmail(body)

	/*if sendErr != nil {
		m.Sent = false
		m.Error = sendErr.Error()
	}*/

	err := saveMessageLog(m)

	return err
}

func sendEmail(body string) error {
	smtpUser := beego.AppConfig.String("smtpUsername")
	smtpPass := beego.AppConfig.String("smtpPassword")
	smtpAddress := beego.AppConfig.String("smtpAddress")
	smtpPort, _ := beego.AppConfig.Int("smtpPort")

	gm := gomail.NewMessage()
	gm.SetHeader("From", smtpUser)
	gm.SetHeader("To", "abc@gmail.com")
	gm.SetHeader("Subject", "Contact Us - avosaweb")
	gm.SetBody("text/html", body)

	d := gomail.NewDialer(smtpAddress, smtpPort, smtpUser, smtpPass)

	err := d.DialAndSend(gm)

	if err != nil {
		log.Panic(err)
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

func saveMessageLog(m Message) error {
	o := orm.NewOrm()
	_, err := o.Insert(&m)

	if err != nil {
		log.Panic(err)
	}

	return err
}
