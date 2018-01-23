package vin

type Descriptor interface {
	GetData() VDS
}

type Manufacturer struct {
	IsMicro    bool
	Category   string
	Name       string
	Descriptor Descriptor
}
