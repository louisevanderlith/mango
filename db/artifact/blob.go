package artifact

import "github.com/louisevanderlith/mango/db"

type Blob struct {
	db.Record
	Data []byte
}

func (o Blob) Validate() (bool, error){
	return true, nil
}