package format_

import (
	"encoding/json"
	"os"
	"sort"
)

type animals struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(path, dir string) error {
	file, err := readJSON(path)
	if err != nil {
		return err
	}

	sort.Slice(file, func(i, j int) bool {
		return file[i].Age < file[j].Age
	})

	f, err := os.CreateTemp(dir, "json-v1.1.0-")
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

func readJSON(path string) ([]animals, error) {
	res := make([]animals, 0, 6)

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
