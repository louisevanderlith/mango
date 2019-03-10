package mango

import "testing"

func Test_MarshalToResult_MustDeserializeToObj(t *testing.T) {
	res := `{"Error":"","Data":["master.html","modal.html","navigation.html","scripts.html","sidebar.html"]}`
	buff := []string{}
	resp, err := marshalToResult([]byte(res), &buff)

	if err != nil {
		t.Error(err)
	}

	result := resp.Data.(*[]string)

	if len(buff) != len(*result) {
		t.Errorf("Buff len %v not equal to Data len %v", len(buff), len(*result))
	}
}
