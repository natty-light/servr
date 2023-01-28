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
