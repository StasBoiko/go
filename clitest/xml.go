package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// type Xml int

func HandleXml() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter xml: ")
	scanner.Scan()
	xmlText := scanner.Text()

	err := ValidateXml(xmlText)
	if err != nil {
		log.Fatal(err)
	}

	err = SaveXml(xmlText)
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateXml(xmlText string) (err error) {
	var js interface{}
	err = xml.Unmarshal([]byte(xmlText), &js)
	if err != nil {
		return err
	}
	return nil
}

func SaveXml(xmlText string) (err error) {
	bytes := []byte(xmlText)
	timeNow := time.Now().String()

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)

	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".xml") {
			fileBytes, err := ioutil.ReadFile(file.Name())
			if err != nil {
				return err
			}
			if strings.Compare(string(fileBytes), xmlText) == 0 {
				return fmt.Errorf("already exist")
			}
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf("./%s.xml", timeNow), bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
