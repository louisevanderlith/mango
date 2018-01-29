package common

type Descriptor interface {
	GetData(vin string) string
}

type Manufacturer struct {
	IsMicro    bool
	Category   string
	Name       string
	Descriptor Descriptor
}

func GetManufacturer(wmiCode string) Manufacturer {
	var result Manufacturer

	if len(wmiCode) == 1 {
		result.IsMicro = wmiCode == "9"
		//result.Descriptor = descriptors.GetDescriptor("", "")
	}

	return result
}
