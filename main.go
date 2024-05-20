package format_

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
)

type Patient struct {
	XMLName xml.Name `xml:"Patient"`
	Name    string   `xml:"Name"`
	Age     int      `xml:"Age"`
	Email   string   `xml:"Email"`
}

type Patients struct {
	XMLName  xml.Name  `xml:"patients"`
	Patients []Patient `xml:"Patient"`
}

func Do(path, dir string) error {
	p, err := readFile(path)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return err
	}

	sort.Slice(p.Patients, func(i, j int) bool {
		return p.Patients[i].Age < p.Patients[j].Age
	})

	f, err := os.CreateTemp(dir, "xml-v2.1.0-")
	if err != nil {
		return err
	}
	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	err = enc.Encode(p)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}

func readFile(path string) (Patients, error) {
	var res Patients

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return res, err
	}
	defer file.Close()

	err = xml.NewDecoder(file).Decode(&res)
	if err != nil {
		fmt.Println("Ошибка декодирования XML:", err)
		return res, err
	}

	return res, nil
}
