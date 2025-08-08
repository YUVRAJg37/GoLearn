package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson[T any](path string) ([]T, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()
	var data []T
	err = json.NewDecoder(file).Decode(&data)

	if err != nil {
		if err.Error() == "EOF" {
			return []T{}, nil
		}
		return nil, err
	}

	return data, err
}

func WriteJson[T any](path string, data []T) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	return encoder.Encode(data)
}

//test
