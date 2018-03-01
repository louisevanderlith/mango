package common

import "errors"

type Class struct {
	Name      string
	MinWeight int
	MaxWeight int
}

var classes []Class

func init() {
	classes = getClasses()
}

func getClassByWeight(weight int) (result Class, err error) {
	for _, v := range classes {
		if v.MinWeight < weight && v.MaxWeight >= weight {
			result = v
			break
		}
	}

	if result.Name == "" {
		err = errors.New("no class found for the weight provided")
	}

	return result, err
}

func getClasses() []Class {
	var result []Class

	result = append(result, Class{"A", 0, 1360})
	result = append(result, Class{"B", 1360, 1814})
	result = append(result, Class{"C", 1814, 2268})
	result = append(result, Class{"D", 2268, 2722})
	result = append(result, Class{"E", 2722, 3175})
	result = append(result, Class{"F", 3175, 3629})
	result = append(result, Class{"G", 3629, 4082})
	result = append(result, Class{"H", 4082, 4536})
	result = append(result, Class{"3", 4536, 6350})
	result = append(result, Class{"4", 6350, 7257})
	result = append(result, Class{"5", 7257, 8845})
	result = append(result, Class{"6", 8845, 11793})
	result = append(result, Class{"7", 11793, 14968})
	result = append(result, Class{"8", 14968, 99999})

	return result
}
