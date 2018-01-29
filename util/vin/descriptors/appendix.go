package descriptors

import (
	"github.com/louisevanderlith/mango/util/vin/common"
)

type WMICategory int

type Descriptors map[string]WMIDescriptor
type WMIGroup map[string]Descriptors

func NewWMIGroup(groupName string) Descriptors {
	group, ok := groups[groupName]

	if ok {
		return group
	}

	groups[groupName] = make(Descriptors)
	return groups[groupName]
}

func (g Descriptors) Add(manufacturerCode, manufacturerName string, category WMICategory, descriptor common.Descriptor) {
	g[manufacturerCode] = buildDescriptor(manufacturerName, category, descriptor)
}

var groups WMIGroup

func init() {
	groups = make(WMIGroup)
	buildGroups()
}

func buildGroups() {
	groupsAlfa()
	groupsAudi()
	groupsBMW()
	groupsFerrari()
	groupsFord()
	groupsGM()
	groupsHonda()
	groupsHyundai()
	groupsIsuzu()
	groupsKIA()
	groupsLamborghini()
	groupsLandRover()
	groupsMaserati()
	groupsMercedes()
	groupsMitsubishi()
	groupsSEAT()
	groupsSkoda()
	groupsSubaru()
	groupsToyota()
	groupsVW()
	groupsVolvo()
}

type WMIDescriptor struct {
	Manufacturer string
	Category     WMICategory
	Descriptor   common.Descriptor
}

const (
	NotSpecified WMICategory = iota
	PassengerCar
	MPV
	Truck
	Bus
	Motorcycle
	ATV
	IncompleteCar
	BasicChassis
)

func GetWMIDescriptor(wmiCode string) WMIDescriptor {
	var result WMIDescriptor
	countryCode := wmiCode[:1]
	manufacturerCode := wmiCode[1:]

	group, hasGroup := groups[countryCode]

	if hasGroup {
		result = group[manufacturerCode]
	}

	return result
}

func buildDescriptor(manufacturerName string, category WMICategory, descriptor common.Descriptor) WMIDescriptor {
	result := WMIDescriptor{
		Manufacturer: manufacturerName,
		Category:     category,
		Descriptor:   descriptor,
	}

	return result
}
