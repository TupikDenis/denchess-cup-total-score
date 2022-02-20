package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type Data struct {
	URL      string
	filename string
}

func (data *Data) download() error {
	response, err := http.Get(data.URL)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Not 200")
	}

	file, err := os.Create(data.filename)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	return nil

}

func rewriteJSON(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	jsonContent := "["

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		jsonContent += fileScanner.Text() + ",\n"
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	if last := len(jsonContent) - 1; last >= 0 && jsonContent[last] == ',' {
		jsonContent = jsonContent[:last]
	}
	jsonContent = jsonContent[:len(jsonContent)-2]
	jsonContent += "]"
	file.Close()

	file1, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	_, err = file1.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("write error: %s", err)
		return
	}

	err = file1.Truncate(0)
	if err != nil {
		log.Fatalf("write error: %s", err)
		return
	}

	_, err = io.WriteString(file1, jsonContent)

	if err != nil {
		log.Fatalf("write error: %s", err)
		return
	}

	file1.Close()
}
