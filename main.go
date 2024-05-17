package format_

import (
	"encoding/json"
	"log"
	"os"
)

type animals struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(path, dir string) error {
	file, err := readFile(path)
	if err != nil {
		return err
	}
	log.Println(file)

	f, err := os.CreateTemp(dir, "new_file-")
	if err != nil {
		return err
	}
	err = json.NewEncoder(f).Encode(file)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func readFile(path string) ([]animals, error) {
	res := make([]animals, 0, 3)

	f, err := os.Open(path)
	if err != nil {
		return res, err
	}

	defer f.Close()

	dec := json.NewDecoder(f)
	for dec.More() {
		var a animals
		err := dec.Decode(&a)
		if err != nil {
			return res, err
		}
		res = append(res, a)
	}

	return res, nil
}
