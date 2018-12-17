package main

import (
	"encoding/json"
	"log"
	"os"
)

func fileCreate(path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func writeFile(path string, line string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte(line))
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func fileLoad(file string) ([]string, error) {
	var urls []string
	urlsFile, err := os.Open(file)
	defer urlsFile.Close()
	if err != nil {
		return urls, err
	}
	jsonParser := json.NewDecoder(urlsFile)
	jsonParser.Decode(&urls)
	return urls, nil
}
