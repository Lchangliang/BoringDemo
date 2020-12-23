package serial

import (
	"encoding/json"
	"fmt"
)

type Serialzable interface {
	Serialize(object interface{}) string
}

type Deserializable interface {
	Deserializable(text string, object interface{}) error
}

type Serialzation struct {}

func (serialzation *Serialzation) Serialize(object interface{}) string {
	data, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return ""
	}
	return string(data)
}

func (serialzable *Serialzation) Deserializable(text string, object interface{}) error {
	err := json.Unmarshal([]byte(text), object)
	return err
}




