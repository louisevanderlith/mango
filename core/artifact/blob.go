package artifact

import (
	"bufio"
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"github.com/louisevanderlith/mango/util/enums"
)

type Blob struct {
	Data []byte
}

type optimFunc func(data image.Image) ([]byte, error)
type optmizer map[enums.OptimizeType]optimFunc

var optimizers optmizer

func init() {
	optimizers = getOptimizers()
}

func (o Blob) Valid() (bool, error) {
	return true, nil
}

func (o *Blob) OptimizeFor(oType enums.OptimizeType) error {
	reader := bytes.NewReader(o.Data)
	decoded, err := imaging.Decode(reader)

	if err == nil {
		opt, hasOpt := optimizers[oType]

		if hasOpt {
			var data []byte
			data, err = opt(decoded)

			if err == nil {
				o.Data = data
			}
		}
	}

	return err
}

func getOptimizers() optmizer {
	result := make(optmizer)

	result[enums.Logo] = optimizeLogo
	result[enums.Banner] = optimizeBanner
	result[enums.Ad] = optimizeAd
	result[enums.Thumb] = optimizeThumb

	return result
}

func optimizeAd(data image.Image) ([]byte, error) {
	return optimize(data, 700, 450, imaging.JPEG)
}

func optimizeBanner(data image.Image) ([]byte, error) {
	return optimize(data, 1536, 864, imaging.JPEG)
}

func optimizeLogo(data image.Image) ([]byte, error) {
	return optimize(data, 256, 128, imaging.PNG)
}

func optimizeThumb(data image.Image) ([]byte, error) {
	return optimize(data, 350, 145, imaging.JPEG)
}

func optimize(data image.Image, width, height int, format imaging.Format) ([]byte, error) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	optImage := imaging.Fit(data, width, height, imaging.Lanczos)

	err := imaging.Encode(writer, optImage, format)

	defer writer.Flush()

	return b.Bytes(), err
}
