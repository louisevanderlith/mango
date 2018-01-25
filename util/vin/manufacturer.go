package vin

type Descriptor interface {
	GetData(vin string) interface{}
}

type Manufacturer struct {
	IsMicro    bool
	Category   string
	Name       string
	Descriptor Descriptor
}
