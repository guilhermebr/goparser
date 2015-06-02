package parser

import (
	"encoding/csv"
	"fmt"
	"io"
)

type InformationCSV struct {
	Name    interface{}
	Email   interface{}
	Gender  interface{}
	Age     interface{}
	Unknown map[string]interface{}
}

func (inf *InformationCSV) UnmarshalCSV(data []string, header []string) error {
	inf.Unknown = map[string]interface{}{}

	for i := 0; i < len(data); i++ {
		switch header[i] {
		case "name":
			{
				inf.Name = data[i]
			}
		case "email":
			{
				inf.Email = data[i]
			}
		case "gender":
			{
				inf.Gender = data[i]
			}
		case "age":
			{
				inf.Age = data[i]
			}
		default:
			{
				inf.Unknown[header[i]] = data[i]
			}
		}
	}

	return nil
}

func ParseCSV(file io.Reader) ([]InformationCSV, error) {
	var informations []InformationCSV
	inf := InformationCSV{}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	// First line is the fields name
	header, err := reader.Read()

	if err != nil {
		return nil, err
	}

	data, _ := reader.ReadAll()

	for _, line := range data {
		inf.UnmarshalCSV(line, header)
		informations = append(informations, inf)
	}

	fmt.Printf("%#v\n", informations)
	return informations, nil
}
