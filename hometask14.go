package gohometask14

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type animal struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type animalList struct {
	List []animal `json:"animals"`
}

func Do(pathRead, pathWrite string) error {
	res, err := readFile(pathRead)
	if err != nil {
		return fmt.Errorf("failed read file: %w", err)
	}

	err = writeFile(pathWrite, res)
	if err != nil {
		return fmt.Errorf("failed write to file: %w", err)
	}
	return nil
}

func readFile(path string) (animalList, error) {
	file, err := os.Open(path)
	if err != nil {
		return animalList{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var animals = animalList{}

	dec := json.NewDecoder(file)

	for dec.More() {
		var animal animal
		err := dec.Decode(&animal)
		if err != nil {
			return animalList{}, fmt.Errorf("failed to decode json: %w", err)
		}

		animals.List = append(animals.List, animal)
	}
	return animals, nil
}

func writeFile(path string, animals animalList) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	file.WriteString(xml.Header)

	dec := xml.NewEncoder(file)
	dec.Indent("", " ")
	dec.Encode(animals.List)
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}
	return nil
}
