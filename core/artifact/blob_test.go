package artifact

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/louisevanderlith/mango/pkg/enums"
)

func getImage(location string) []byte {
	dat, err := ioutil.ReadFile(location)

	if err != nil {
		log.Println(err)
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

	data := Blob{}
	data.Data = getImage("./test.png")

	err := data.OptimizeFor(enums.Logo)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Logo_JPG2PNG(t *testing.T) {
	resultName := "jpg2png_logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.jpg")

	err := data.OptimizeFor(enums.Logo)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Banner_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_banner.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.png")

	err := data.OptimizeFor(enums.Banner)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Banner_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_banner.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.jpg")

	err := data.OptimizeFor(enums.Banner)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Ad_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_ad.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.png")

	err := data.OptimizeFor(enums.Ad)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Ad_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_ad.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.jpg")

	err := data.OptimizeFor(enums.Ad)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Thumb_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_thumb.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.png")

	err := data.OptimizeFor(enums.Thumb)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Thumb_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_thumb.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./test.jpg")

	err := data.OptimizeFor(enums.Thumb)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, data.Data)
}

func TestBlob_OptimizeFor_Ad(t *testing.T) {
	resultName := "logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := Blob{}
	data.Data = getImage("./logo.png")

	err := data.OptimizeFor(enums.Logo)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	objData := data.Data
	t.Log(len(objData))
	if len(objData) > 0 {
		saveImage(writeLocation, objData)
	} else {
		t.Error("Image Zero")
	}

}
