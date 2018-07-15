package things

import (
	"fmt"

	"github.com/louisevanderlith/husk"
)

type Manufacturer struct {
	Name        string `hsk:"size(50)"`
	Description string `hsk:"null;size(255)"`
	Models      Models
}

func (o Manufacturer) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func seedManufacturer() {
	data := Manufacturers{}
	data.Add(Manufacturer{Name: "Acura"})
	data.Add(Manufacturer{Name: "Alfa Romeo"})
	data.Add(Manufacturer{Name: "Aptera"})
	data.Add(Manufacturer{Name: "Aston Martin"})
	data.Add(Manufacturer{Name: "Audi"})
	data.Add(Manufacturer{Name: "Austin"})
	data.Add(Manufacturer{Name: "Bentley"})
	data.Add(Manufacturer{Name: "BMW"})
	data.Add(Manufacturer{Name: "Bugatti"})
	data.Add(Manufacturer{Name: "Buick"})
	data.Add(Manufacturer{Name: "Cadillac"})
	data.Add(Manufacturer{Name: "Chevrolet"})
	data.Add(Manufacturer{Name: "Chrysler"})
	data.Add(Manufacturer{Name: "Citroën"})
	data.Add(Manufacturer{Name: "Corbin"})
	data.Add(Manufacturer{Name: "Daewoo"})
	data.Add(Manufacturer{Name: "Daihatsu"})
	data.Add(Manufacturer{Name: "Dodge"})
	data.Add(Manufacturer{Name: "Eagle"})
	data.Add(Manufacturer{Name: "Fairthorpe"})
	data.Add(Manufacturer{Name: "Ferrari"})
	data.Add(Manufacturer{Name: "FIAT"})
	data.Add(Manufacturer{Name: "Fillmore"})
	data.Add(Manufacturer{Name: "Foose"})
	data.Add(Manufacturer{Name: "Ford"})
	data.Add(Manufacturer{Name: "Geo"})
	data.Add(Manufacturer{Name: "GMC"})
	data.Add(Manufacturer{Name: "Hillman"})
	data.Add(Manufacturer{Name: "Holden"})
	data.Add(Manufacturer{Name: "Honda"})
	data.Add(Manufacturer{Name: "HUMMER"})
	data.Add(Manufacturer{Name: "Hyundai"})
	data.Add(Manufacturer{Name: "Infiniti"})
	data.Add(Manufacturer{Name: "Isuzu"})
	data.Add(Manufacturer{Name: "Jaguar"})
	data.Add(Manufacturer{Name: "Jeep"})
	data.Add(Manufacturer{Name: "Jensen"})
	data.Add(Manufacturer{Name: "Kia"})
	data.Add(Manufacturer{Name: "Lamborghini"})
	data.Add(Manufacturer{Name: "Land Rover"})
	data.Add(Manufacturer{Name: "Lexus"})
	data.Add(Manufacturer{Name: "Lincoln"})
	data.Add(Manufacturer{Name: "Lotus"})
	data.Add(Manufacturer{Name: "Maserati"})
	data.Add(Manufacturer{Name: "Maybach"})
	data.Add(Manufacturer{Name: "Mazda"})
	data.Add(Manufacturer{Name: "McLaren"})
	data.Add(Manufacturer{Name: "Mercedes-Benz"})
	data.Add(Manufacturer{Name: "Mercury"})
	data.Add(Manufacturer{Name: "Merkur"})
	data.Add(Manufacturer{Name: "MG"})
	data.Add(Manufacturer{Name: "MINI"})
	data.Add(Manufacturer{Name: "Mitsubishi"})
	data.Add(Manufacturer{Name: "Morgan"})
	data.Add(Manufacturer{Name: "Nissan"})
	data.Add(Manufacturer{Name: "Oldsmobile"})
	data.Add(Manufacturer{Name: "Panoz"})
	data.Add(Manufacturer{Name: "Peugeot"})
	data.Add(Manufacturer{Name: "Plymouth"})
	data.Add(Manufacturer{Name: "Pontiac"})
	data.Add(Manufacturer{Name: "Porsche"})
	data.Add(Manufacturer{Name: "Ram"})
	data.Add(Manufacturer{Name: "Rambler"})
	data.Add(Manufacturer{Name: "Renault"})
	data.Add(Manufacturer{Name: "Rolls-Royce"})
	data.Add(Manufacturer{Name: "Saab"})
	data.Add(Manufacturer{Name: "Saturn"})
	data.Add(Manufacturer{Name: "Scion"})
	data.Add(Manufacturer{Name: "Shelby"})
	data.Add(Manufacturer{Name: "Smart"})
	data.Add(Manufacturer{Name: "Spyker"})
	data.Add(Manufacturer{Name: "Spyker Cars"})
	data.Add(Manufacturer{Name: "Studebaker"})
	data.Add(Manufacturer{Name: "Subaru"})
	data.Add(Manufacturer{Name: "Suzuki"})
	data.Add(Manufacturer{Name: "Tesla"})
	data.Add(Manufacturer{Name: "Toyota"})
	data.Add(Manufacturer{Name: "Volkswagen"})
	data.Add(Manufacturer{Name: "Volvo"})

	var items Manufacturers
	err := Ctx.Manufacturers.Read(&Manufacturer{}, &items)

	if err == nil {
		if len(items) == 0 {
			_, err = Ctx.Manufacturers.CreateMulti(data)

			if err != nil {
				fmt.Println("seedManufacturer:", err)
			}
		}
	} else {
		fmt.Println("seedManufacturer:", err)
	}
}
