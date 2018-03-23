package things

import (
	"fmt"

	"github.com/louisevanderlith/db"
	"github.com/louisevanderlith/mango/util"
)

type Manufacturer struct {
	db.Record
	Name        string   `orm:"size(50)"`
	Description string   `orm:"null;size(255)"`
	Models      []*Model `orm:"reverse(many)"`
}

func (o Manufacturer) Validate() (bool, error) {
	return util.ValidateStruct(&o)
}

func seedManufacturer() {
	data := []Manufacturer{
		Manufacturer{Name: "Acura"},
		Manufacturer{Name: "Alfa Romeo"},
		Manufacturer{Name: "Aptera"},
		Manufacturer{Name: "Aston Martin"},
		Manufacturer{Name: "Audi"},
		Manufacturer{Name: "Austin"},
		Manufacturer{Name: "Bentley"},
		Manufacturer{Name: "BMW"},
		Manufacturer{Name: "Bugatti"},
		Manufacturer{Name: "Buick"},
		Manufacturer{Name: "Cadillac"},
		Manufacturer{Name: "Chevrolet"},
		Manufacturer{Name: "Chrysler"},
		Manufacturer{Name: "Citroën"},
		Manufacturer{Name: "Corbin"},
		Manufacturer{Name: "Daewoo"},
		Manufacturer{Name: "Daihatsu"},
		Manufacturer{Name: "Dodge"},
		Manufacturer{Name: "Eagle"},
		Manufacturer{Name: "Fairthorpe"},
		Manufacturer{Name: "Ferrari"},
		Manufacturer{Name: "FIAT"},
		Manufacturer{Name: "Fillmore"},
		Manufacturer{Name: "Foose"},
		Manufacturer{Name: "Ford"},
		Manufacturer{Name: "Geo"},
		Manufacturer{Name: "GMC"},
		Manufacturer{Name: "Hillman"},
		Manufacturer{Name: "Holden"},
		Manufacturer{Name: "Honda"},
		Manufacturer{Name: "HUMMER"},
		Manufacturer{Name: "Hyundai"},
		Manufacturer{Name: "Infiniti"},
		Manufacturer{Name: "Isuzu"},
		Manufacturer{Name: "Jaguar"},
		Manufacturer{Name: "Jeep"},
		Manufacturer{Name: "Jensen"},
		Manufacturer{Name: "Kia"},
		Manufacturer{Name: "Lamborghini"},
		Manufacturer{Name: "Land Rover"},
		Manufacturer{Name: "Lexus"},
		Manufacturer{Name: "Lincoln"},
		Manufacturer{Name: "Lotus"},
		Manufacturer{Name: "Maserati"},
		Manufacturer{Name: "Maybach"},
		Manufacturer{Name: "Mazda"},
		Manufacturer{Name: "McLaren"},
		Manufacturer{Name: "Mercedes-Benz"},
		Manufacturer{Name: "Mercury"},
		Manufacturer{Name: "Merkur"},
		Manufacturer{Name: "MG"},
		Manufacturer{Name: "MINI"},
		Manufacturer{Name: "Mitsubishi"},
		Manufacturer{Name: "Morgan"},
		Manufacturer{Name: "Nissan"},
		Manufacturer{Name: "Oldsmobile"},
		Manufacturer{Name: "Panoz"},
		Manufacturer{Name: "Peugeot"},
		Manufacturer{Name: "Plymouth"},
		Manufacturer{Name: "Pontiac"},
		Manufacturer{Name: "Porsche"},
		Manufacturer{Name: "Ram"},
		Manufacturer{Name: "Rambler"},
		Manufacturer{Name: "Renault"},
		Manufacturer{Name: "Rolls-Royce"},
		Manufacturer{Name: "Saab"},
		Manufacturer{Name: "Saturn"},
		Manufacturer{Name: "Scion"},
		Manufacturer{Name: "Shelby"},
		Manufacturer{Name: "Smart"},
		Manufacturer{Name: "Spyker"},
		Manufacturer{Name: "Spyker Cars"},
		Manufacturer{Name: "Studebaker"},
		Manufacturer{Name: "Subaru"},
		Manufacturer{Name: "Suzuki"},
		Manufacturer{Name: "Tesla"},
		Manufacturer{Name: "Toyota"},
		Manufacturer{Name: "Volkswagen"},
		Manufacturer{Name: "Volvo"},
	}

	var items []Manufacturer
	err := Ctx.Manufacturer.Read(&Manufacturer{}, &items)

	if err == nil {
		if len(items) == 0 {
			batchCount := len(data)
			_, err = Ctx.Manufacturer.CreateMulti(batchCount, data)

			if err != nil {
				fmt.Println("seedManufacturer:", err)
			}
		}
	} else {
		fmt.Println("seedManufacturer:", err)
	}
}
