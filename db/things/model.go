package things

import (
	"fmt"

	"github.com/louisevanderlith/mango/db"
	"github.com/louisevanderlith/mango/util"
)

type Model struct {
	db.Record
	Name         string        `orm:"size(50)"`
	Manufacturer *Manufacturer `orm:"rel(fk)"`
}

func (o Model) Validate() (bool, error) {
	return util.ValidateStruct(o)
}

func seedModel() {
	data := []Model{
		Model{Name: "1000", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "280ZX", Manufacturer: &Manufacturer{Record: db.Record{ID: 55}}},
		Model{Name: "2CV", Manufacturer: &Manufacturer{Record: db.Record{ID: 14}}},
		Model{Name: "5000S", Manufacturer: &Manufacturer{Record: db.Record{ID: 5}}},
		Model{Name: "600", Manufacturer: &Manufacturer{Record: db.Record{ID: 8}}},
		Model{Name: "6000", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "626", Manufacturer: &Manufacturer{Record: db.Record{ID: 46}}},
		Model{Name: "900", Manufacturer: &Manufacturer{Record: db.Record{ID: 66}}},
		Model{Name: "914", Manufacturer: &Manufacturer{Record: db.Record{ID: 61}}},
		Model{Name: "944", Manufacturer: &Manufacturer{Record: db.Record{ID: 61}}},
		Model{Name: "Accord", Manufacturer: &Manufacturer{Record: db.Record{ID: 30}}},
		Model{Name: "Alliance", Manufacturer: &Manufacturer{Record: db.Record{ID: 64}}},
		Model{Name: "Aries", Manufacturer: &Manufacturer{Record: db.Record{ID: 18}}},
		Model{Name: "Aspen", Manufacturer: &Manufacturer{Record: db.Record{ID: 18}}},
		Model{Name: "Avanti", Manufacturer: &Manufacturer{Record: db.Record{ID: 73}}},
		Model{Name: "Beetle", Manufacturer: &Manufacturer{Record: db.Record{ID: 78}}},
		Model{Name: "Bel Air", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Biturbo", Manufacturer: &Manufacturer{Record: db.Record{ID: 44}}},
		Model{Name: "Bonneville", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Brat", Manufacturer: &Manufacturer{Record: db.Record{ID: 74}}},
		Model{Name: "Bronco II", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Bronco", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Camaro", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Capri", Manufacturer: &Manufacturer{Record: db.Record{ID: 49}}},
		Model{Name: "Caprice", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Celica", Manufacturer: &Manufacturer{Record: db.Record{ID: 77}}},
		Model{Name: "Century", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "Charger", Manufacturer: &Manufacturer{Record: db.Record{ID: 18}}},
		Model{Name: "Citation", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Civic", Manufacturer: &Manufacturer{Record: db.Record{ID: 30}}},
		Model{Name: "Classic", Manufacturer: &Manufacturer{Record: db.Record{ID: 63}}},
		Model{Name: "Continental", Manufacturer: &Manufacturer{Record: db.Record{ID: 42}}},
		Model{Name: "Cordia", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Corvair 500", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Corvair", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Corvette", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Cougar", Manufacturer: &Manufacturer{Record: db.Record{ID: 49}}},
		Model{Name: "Country", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "CR-X", Manufacturer: &Manufacturer{Record: db.Record{ID: 30}}},
		Model{Name: "CX", Manufacturer: &Manufacturer{Record: db.Record{ID: 13}}},
		Model{Name: "Daytona", Manufacturer: &Manufacturer{Record: db.Record{ID: 18}}},
		Model{Name: "E150", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "E250", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "E-Class", Manufacturer: &Manufacturer{Record: db.Record{ID: 48}}},
		Model{Name: "Electra", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "Escort", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "E-Series", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Esprit Turbo", Manufacturer: &Manufacturer{Record: db.Record{ID: 43}}},
		Model{Name: "EXP", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "F150", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "F250", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Fairlane", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Falcon", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Fiero", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Fillmore", Manufacturer: &Manufacturer{Record: db.Record{ID: 23}}},
		Model{Name: "Firebird", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Firefly", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Fleetwood", Manufacturer: &Manufacturer{Record: db.Record{ID: 11}}},
		Model{Name: "Fury", Manufacturer: &Manufacturer{Record: db.Record{ID: 59}}},
		Model{Name: "Galant", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Galaxie", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "GLC", Manufacturer: &Manufacturer{Record: db.Record{ID: 46}}},
		Model{Name: "Golf", Manufacturer: &Manufacturer{Record: db.Record{ID: 78}}},
		Model{Name: "Grand Marquis", Manufacturer: &Manufacturer{Record: db.Record{ID: 49}}},
		Model{Name: "Grand Prix", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "GT350", Manufacturer: &Manufacturer{Record: db.Record{ID: 69}}},
		Model{Name: "GT500", Manufacturer: &Manufacturer{Record: db.Record{ID: 69}}},
		Model{Name: "GTO", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Horizon", Manufacturer: &Manufacturer{Record: db.Record{ID: 59}}},
		Model{Name: "Imperial", Manufacturer: &Manufacturer{Record: db.Record{ID: 13}}},
		Model{Name: "Interceptor", Manufacturer: &Manufacturer{Record: db.Record{ID: 37}}},
		Model{Name: "Jetta", Manufacturer: &Manufacturer{Record: db.Record{ID: 78}}},
		Model{Name: "Laser", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "LeMans", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "LeSabre", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "LTD Crown Victoria", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "LTD", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "LUV", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Lynx", Manufacturer: &Manufacturer{Record: db.Record{ID: 49}}},
		Model{Name: "Mark VII", Manufacturer: &Manufacturer{Record: db.Record{ID: 42}}},
		Model{Name: "Marquis", Manufacturer: &Manufacturer{Record: db.Record{ID: 49}}},
		Model{Name: "MGB", Manufacturer: &Manufacturer{Record: db.Record{ID: 51}}},
		Model{Name: "Mini Cooper S", Manufacturer: &Manufacturer{Record: db.Record{ID: 6}}},
		Model{Name: "Mini Cooper", Manufacturer: &Manufacturer{Record: db.Record{ID: 6}}},
		Model{Name: "Mini", Manufacturer: &Manufacturer{Record: db.Record{ID: 6}}},
		Model{Name: "Minx Magnificent", Manufacturer: &Manufacturer{Record: db.Record{ID: 28}}},
		Model{Name: "Mirage", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Model T", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Monte Carlo", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Monza", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Mustang", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Omni", Manufacturer: &Manufacturer{Record: db.Record{ID: 18}}},
		Model{Name: "Pajero", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Parisienne", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Prelude", Manufacturer: &Manufacturer{Record: db.Record{ID: 30}}},
		Model{Name: "Quantum", Manufacturer: &Manufacturer{Record: db.Record{ID: 78}}},
		Model{Name: "Quattroporte", Manufacturer: &Manufacturer{Record: db.Record{ID: 44}}},
		Model{Name: "Ranger", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Reliant", Manufacturer: &Manufacturer{Record: db.Record{ID: 59}}},
		Model{Name: "Riviera", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "Roadrunner", Manufacturer: &Manufacturer{Record: db.Record{ID: 59}}},
		Model{Name: "Rockette", Manufacturer: &Manufacturer{Record: db.Record{ID: 20}}},
		Model{Name: "RX-7", Manufacturer: &Manufacturer{Record: db.Record{ID: 46}}},
		Model{Name: "Scirocco", Manufacturer: &Manufacturer{Record: db.Record{ID: 78}}},
		Model{Name: "S-Class", Manufacturer: &Manufacturer{Record: db.Record{ID: 48}}},
		Model{Name: "SJ 410", Manufacturer: &Manufacturer{Record: db.Record{ID: 75}}},
		Model{Name: "Skyhawk", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "SL-Class", Manufacturer: &Manufacturer{Record: db.Record{ID: 48}}},
		Model{Name: "SM", Manufacturer: &Manufacturer{Record: db.Record{ID: 13}}},
		Model{Name: "Somerset", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "Space", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Special", Manufacturer: &Manufacturer{Record: db.Record{ID: 10}}},
		Model{Name: "Starion", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Sunbird", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Tempest", Manufacturer: &Manufacturer{Record: db.Record{ID: 60}}},
		Model{Name: "Tempo", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Thunderbird", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Topaz", Manufacturer: &Manufacturer{Record: db.Record{ID: 49}}},
		Model{Name: "Torino", Manufacturer: &Manufacturer{Record: db.Record{ID: 25}}},
		Model{Name: "Toronado", Manufacturer: &Manufacturer{Record: db.Record{ID: 56}}},
		Model{Name: "Town Car", Manufacturer: &Manufacturer{Record: db.Record{ID: 42}}},
		Model{Name: "Tredia", Manufacturer: &Manufacturer{Record: db.Record{ID: 53}}},
		Model{Name: "Vanagon", Manufacturer: &Manufacturer{Record: db.Record{ID: 78}}},
		Model{Name: "Vega", Manufacturer: &Manufacturer{Record: db.Record{ID: 12}}},
		Model{Name: "Volare", Manufacturer: &Manufacturer{Record: db.Record{ID: 59}}},
		Model{Name: "W123", Manufacturer: &Manufacturer{Record: db.Record{ID: 48}}},
		Model{Name: "W126", Manufacturer: &Manufacturer{Record: db.Record{ID: 48}}},
		Model{Name: "W201", Manufacturer: &Manufacturer{Record: db.Record{ID: 48}}},
	}

	var items []Model
	err := Ctx.Model.Read(&Model{}, &items)

	if err == nil {
		if len(items) == 0 {
			batchCount := len(data)
			_, err = Ctx.Model.CreateMulti(batchCount, data)

			if err != nil {
				fmt.Println("seedModel:", err)
			}
		}
	} else {
		fmt.Println("seedModel:", err)
	}
}
