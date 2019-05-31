package enums

//ServiceType Services must identify as an API, APX, APP
//API: Lowest logic layer. API can not call another API.
//APX: Workflow Executor
//APP: Presentation layer. 
type ServiceType int

const (
	API ServiceType = iota
	APP
	ANY
	APX
)

var servicetypes = [...]string{
	"API",
	"APP",
	"ANY",
	"APX"}

func (s ServiceType) String() string {
	return servicetypes[s]
}
