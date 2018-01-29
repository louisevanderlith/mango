package common

type VINDetail struct {
	Country       string
	Type          string
	Manufacturer  string
	FullVIN       string
	Year          int
	ProducttionNo int64
	IsValid       bool
}

var types []string

func init() {
	types = getTypes()
	classes = getClasses()
}

func getTypes() []string {
	return []string{
		"Passenger car",
		"Multipurpose passenger vehicle",
		"Truck",
		"Bus",
		"Trailer",
		"Motorcycle",
		"Incomplete vehicle other than trailer",
		"Low speed vehicle",
	}
}

func Deserialize(val string) VINDetail {
	var result VINDetail

	return result
}
