package models

import (
	"bytes"
	"encoding/json"
)

type BioLink struct {

	ID string `json:"id"`
	Name string `json:"name"`
	From string `json:"from"`
	To string `json:"to"`

}

func (link BioLink) ToJson() string {
	var buffer *bytes.Buffer = &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if encoder != nil{
		encoder.Encode(link)
	}
	return buffer.String()
}
