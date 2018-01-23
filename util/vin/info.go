package vin

type Info struct {
	WMI WMI //[0:2]
	VDS VDS //[3:8]
	VIS VIS //[9:16]
}

func getInfo(vin string) Info {
	var result Info

	return result
}
