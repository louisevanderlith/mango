package enums

type ServiceType int

const (
	DB ServiceType = iota
	API
	APP
	PROXY
	ANY
)

var servicetypes = [...]string{
	"DB",
	"API",
	"APP",
	"PROXY",
	"ANY"}

func (s ServiceType) String() string {
	return servicetypes[s]
}
