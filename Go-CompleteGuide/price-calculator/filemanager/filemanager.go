package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([] string, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Failed to open file!")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Unable to read file content!")
	}

	file.Close()

	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	
	file, err:=	os.Create(path)

	if err != nil {
		return errors.New("Failed to crate file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert date to JSON.")
	}

	file.Close()
	return nil
	
}