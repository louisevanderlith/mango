package secure

// Environment provides indicates in which environment a system is
type TraceType int

const (
	TraceRegister TraceType = iota
	TraceLogin
	TraceFail
	TraceLogout
)

var environments = [...]string{
	"TraceRegister",
	"LoTraceLogingin",
	"TraceFail",
	"TraceLogout"}

func (e TraceType) String() string {
	return environments[e]
}
