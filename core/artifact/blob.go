package artifact

import (
	"bufio"
	"bytes"
	"errors"
	"image"

	"github.com/disintegration/imaging"
	"github.com/louisevanderlith/mango/pkg/enums"
)

type Blob struct {
	Data []byte
}

type optimFunc func(data image.Image) (result []byte, mimetype string, err error)
type optmizer map[enums.OptimizeType]optimFunc

var optimizers optmizer

func init() {
	optimizers = getOptimizers()
}

func (o Blob) Valid() (bool, error) {
	return true, nil
}

func NewBLOB(data []byte, purpose string) (*Blob, string, error) {
	result := &Blob{Data: data}

	targetType := enums.GetOptimizeType(purpose)
	mime, err := result.OptimizeFor(targetType)

	return result, mime, err
}

func (o *Blob) OptimizeFor(oType enums.OptimizeType) (string, error) {
	reader := bytes.NewReader(o.Data)
	decoded, err := imaging.Decode(reader)

	if err != nil {
		return "", err
	}

	opt, hasOpt := optimizers[oType]

	if !hasOpt {
		return "", errors.New("optimizer Type not found")
	}

	data, mime, err := opt(decoded)

	if err != nil {
		return "", err
	}

	o.Data = data

	return mime, err
}

func getOptimizers() optmizer {
	result := make(optmizer)

	result[enums.Logo] = optimizeLogo
	result[enums.Banner] = optimizeBanner
	result[enums.Ad] = optimizeAd
	result[enums.Thumb] = optimizeThumb

	return result
}

func optimizeAd(data image.Image) ([]byte, string, error) {
	return optimize(data, 700, 450, imaging.JPEG)
}

func optimizeBanner(data image.Image) ([]byte, string, error) {
	return optimize(data, 1536, 864, imaging.JPEG)
}

func optimizeLogo(data image.Image) ([]byte, string, error) {
	return optimize(data, 256, 128, imaging.PNG)
}

func optimizeThumb(data image.Image) ([]byte, string, error) {
	return optimize(data, 350, 145, imaging.JPEG)
}

func optimize(data image.Image, width, height int, format imaging.Format) ([]byte, string, error) {
	var b bytes.Buffer

	writer := bufio.NewWriter(&b)
	defer writer.Flush()

	optImage := imaging.Fit(data, width, height, imaging.Lanczos)

	err := imaging.Encode(writer, optImage, format)

	mimetype := "image/" + format.String()

	return b.Bytes(), mimetype, err
}
