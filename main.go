package format_

import (
	"encoding/xml"
	"fmt"
	"os"
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
		return err
	}

	f, err := os.CreateTemp(dir, "xml-v2.0.0-")
	if err != nil {
		return err
	}
	defer f.Close()

	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	err = enc.Encode(p)
	if err != nil {
		return err
	}

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
