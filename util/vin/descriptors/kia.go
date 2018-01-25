package descriptors

type KIA struct {
}

func (d KIA) GetData(vin string) interface{} {
	return 0
}

func groupsKIA() {
	groupm := NewWMIGroup("M")
	groupm.Add("S0", "KIA", NotSpecified, KIA{})
}
