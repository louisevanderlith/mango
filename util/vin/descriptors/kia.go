package descriptors

type KIA struct {
}

func (d KIA) GetData(vinNo string) string {
	return ""
}

func groupsKIA() {
	groupm := NewWMIGroup("M")
	groupm.Add("S0", "KIA", NotSpecified, KIA{})
}
