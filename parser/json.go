package parser

import (
	"encoding/json"
	"fmt"
)

type InformationJson struct {
	Name    interface{}
	Email   interface{}
	Gender  interface{}
	Age     interface{}
	Unknown map[string]json.RawMessage
}

func (inf *InformationJson) UnmarshalJSON(data []byte) error {
	inf.Unknown = map[string]json.RawMessage{}
	err := json.Unmarshal(data, &inf.Unknown)
	if err != nil {
		return err
	}

	err = json.Unmarshal(inf.Unknown["name"], &inf.Name)
	if err != nil {
		fmt.Println(err)
	}
	delete(inf.Unknown, "name")

	err = json.Unmarshal(inf.Unknown["email"], &inf.Email)
	if err != nil {
		fmt.Println(err)
	}
	delete(inf.Unknown, "email")

	err = json.Unmarshal(inf.Unknown["gender"], &inf.Gender)
	if err != nil {
		fmt.Println(err)
	}
	delete(inf.Unknown, "gender")

	err = json.Unmarshal(inf.Unknown["age"], &inf.Age)
	if err != nil {
		fmt.Println(err)
	}
	delete(inf.Unknown, "age")
	return nil
}

func ParseJSON(file []byte) []InformationJson {
	var inf []InformationJson

	//inf := &InformationJson{}
	err := json.Unmarshal(file, &inf)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("%#v\n", inf)
	return inf
}
