package secure

type LoginTrace struct {
	Record
	Location string `orm:"null;size(128)"`
	IP       string `orm:"null;size(50)"`
	Allowed  bool   `orm:"default(false)"`
	User     *User  `orm:"rel(fk)"`
}

func createLoginTrace() {

}
