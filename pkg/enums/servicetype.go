package enums

type ServiceType int

const (
	API ServiceType = iota
	APP
	ANY
)

var servicetypes = [...]string{
	"API",
	"APP",
	"ANY"}

func (s ServiceType) String() string {
	return servicetypes[s]
}
