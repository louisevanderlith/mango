package artifact

import (
	"github.com/louisevanderlith/mango/db"
	"github.com/disintegration/imaging"
	"github.com/louisevanderlith/mango/util/enums"
	"image"
	"bytes"
	"bufio"
)

type Blob struct {
	db.Record
	Data string
}

type optmizer map[enums.OptimizeType]func(data image.Image) []byte

var optimizers optmizer

func init() {
	optimizers = getOptimizers()
}

func (o Blob) Validate() (bool, error) {
	return true, nil
}

func (o *Blob) OptimizeFor(oType enums.OptimizeType) error {
	reader := bytes.NewReader(o.GetData())
	decoded, err := imaging.Decode(reader)

	if err == nil {
		opt, hasOpt := optimizers[oType]

		if hasOpt {
			o.SetData(opt(decoded))
		}
	}

	return err
}

func (o *Blob) GetData() []byte {
	return []byte(o.Data)
}

func (o *Blob)SetData(blob []byte) {
	o.Data = string(blob)
}

func getOptimizers() optmizer {
	result := make(optmizer)

	result[enums.Logo] = optimizeLogo
	result[enums.Banner] = optimizeBanner
	result[enums.Ad] = optimizeAd

	return result
}

func optimizeAd(data image.Image) []byte {
	var b bytes.Buffer

	writer := bufio.NewWriter(&b)
	optImage := imaging.Fill(data, 700, 450, imaging.Center, imaging.Lanczos)
	imaging.Encode(writer, optImage, imaging.JPEG)

	defer writer.Flush()

	return b.Bytes()
}

func optimizeBanner(data image.Image) []byte {
	var b bytes.Buffer

	writer := bufio.NewWriter(&b)
	optImage := imaging.Fill(data, 1536, 864, imaging.Center, imaging.Lanczos)
	imaging.Encode(writer, optImage, imaging.JPEG)

	defer writer.Flush()

	return b.Bytes()
}

func optimizeLogo(data image.Image) []byte {
	var b bytes.Buffer

	writer := bufio.NewWriter(&b)
	optImage := imaging.Fill(data, 128, 64, imaging.Center, imaging.Lanczos)
	imaging.Encode(writer, optImage, imaging.PNG)

	defer writer.Flush()

	return b.Bytes()
}
