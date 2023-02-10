package utils

import (
	"bufio"
	"os"
)

func Read() string {
	f, _ := os.Open(`C:\Users\jagal\OneDrive\Desktop\serveR\data.txt`)
	var line string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line = scanner.Text()
	}
	return line
}

func Write(schools []string, id string) error {
	var fileName string = id + ".csv"

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, school := range schools {
		_, err := f.WriteString(school + ",\n")
		if err != nil {
			return err
		}
	}

	return nil
}
