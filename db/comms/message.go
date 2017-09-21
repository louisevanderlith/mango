package comms

type Message struct {
	db.Record
	Name       string    `orm:"size(50)"`
	Email      string    `orm:"size(128)"`
	Phone      string    `orm:"size(15)"`
	Body       string    `orm:"size(1024)"`
	Sent       bool      `orm:"default(false)"`
	Error      string    `orm:"null;size(2048)"`
}
