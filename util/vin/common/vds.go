package common

type VDSControl interface {
	GetPassengerCar(sections VINSections, year int) VDS
}

type VDS struct {
	Model         string
	BodyStyle     string
	Doors         string
	EngineModel   string
	EngineSize    string
	Trim          string
	Transmission  string
	Gears         string
	Extras        []string
	AssemblyPlant string
}

func (v VDS) GetPassengerCar(sections VINSections, year int) VDS {
	return v
}
