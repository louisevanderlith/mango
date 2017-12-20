package artifact

import (
	"testing"
	"io/ioutil"
	"fmt"
	"github.com/louisevanderlith/mango/util/enums"
	"os"
)

func getImage(location string) []byte {
	dat, err := ioutil.ReadFile(location)

	if err != nil {
		fmt.Println(err)
	}

	return dat
}

func saveImage(location string, file []byte) {
	ioutil.WriteFile(location, file, 0644)
}

func TestBlob_OptimizeFor_Logo_PNG2PNG(t *testing.T) {
	resultName := "png2png_logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{
		Data: getImage("./test.png"),
	}

	err := data.OptimizeFor(enums.Logo)

	if err != nil {
		t.Error("Error occured:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Logo_JPG2PNG(t *testing.T) {
	resultName := "jpg2png_logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{
		Data: getImage("./test.jpg"),
	}

	err := data.OptimizeFor(enums.Logo)

	if err != nil {
		t.Error("Error occured:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Banner_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_banner.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{
		Data: getImage("./test.png"),
	}

	err := data.OptimizeFor(enums.Banner)

	if err != nil {
		t.Error("Error occured:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Banner_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_banner.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{
		Data: getImage("./test.jpg"),
	}

	err := data.OptimizeFor(enums.Banner)

	if err != nil {
		t.Error("Error occured:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Ad_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_ad.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{
		Data: getImage("./test.png"),
	}

	err := data.OptimizeFor(enums.Ad)

	if err != nil {
		t.Error("Error occured:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Ad_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_ad.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{
		Data: getImage("./test.jpg"),
	}

	err := data.OptimizeFor(enums.Ad)

	if err != nil {
		t.Error("Error occured:", err)
	}

	saveImage(writeLocation, data.Data)
}
