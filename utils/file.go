package utils

import (
	"bufio"
	"os"
)

func Read() string {
	f, _ := os.Open(`./data.txt`)
	var line string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line = scanner.Text()
	}
	f.Close()
	return line
}

func Write(schools []string, id string) (string, error) {
	var fileName string = id + ".csv"

	f, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}
	defer f.Close()

	for _, school := range schools {
		_, err := f.WriteString(school + ",\n")
		if err != nil {
			return fileName, err
		}
	}
	f.Close()
	return fileName, nil
}
