package descriptors

import "testing"

func TestGetData_Civic_Success(t *testing.T) {
	input := "1HGEJ8144XL019972"
	actual := Honda{}.GetData(input)

	t.Log(actual)
	t.Fail()
}

func TestGetData_CRV_Success(t *testing.T) {
	input := "JHLRD68405C200888"
	actual := Honda{}.GetData(input)

	t.Log(actual)
	t.Fail()
}

func TestGetData_Acura_Success(t *testing.T) {
	input := "JH4NA1158NT000999"
	actual := Honda{}.GetData(input)

	t.Log(actual)
	t.Fail()
}
