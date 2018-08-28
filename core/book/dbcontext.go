package book

type context struct {
	Vehicles     vehiclesTable
	VINS         vinsTable
	Services     servicesTable
	ServiceItems serviceItemsTable
}

var ctx context

func init() {
	ctx = context{
		Vehicles:     NewVehiclesTable(),
		VINS:         NewVINSTable(),
		Services:     NewServicesTable(),
		ServiceItems: NewServiceItemsTable(),
	}
}
