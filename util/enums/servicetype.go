package enums

type ServiceType int

const (
	DB ServiceType = iota
	API
	APP
	ANY
)

var servicetypes = [...]string{
	"DB",
	"API",
	"APP",
	"ANY"}

func (s ServiceType) String() string {
	return servicetypes[s]
}
