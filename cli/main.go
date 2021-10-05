package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type video struct {
	Id string
}

type jsonData struct {
	json.RawMessage
}

type Message struct {
	Body string
}

func main() {
	//urfave/cli!!!
	// xmlCmd := flag.NewFlagSet("-xml", flag.ExitOnError)
	jsonCmd := flag.NewFlagSet("-json", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected '-xml' or '-json' commands")
		// os.Exit(1)
	}

	switch os.Args[1] {
	// case "-xml": // if its the 'xml' command
	// HandleXml(xmlCmd)
	case "-json": // if its the 'json' command
		HandleJson(jsonCmd)
		// fmt.Println("Hi!")
	default: // if we don't understand the input
		fmt.Println("unexpected value")
		// os.Exit(2)
	}
}

func HandleJson(jsonCmd *flag.FlagSet) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter json: ")
	scanner.Scan()
	jsonText := scanner.Text()

	err := ValidateJson(jsonText)
	if err != nil {
		log.Fatal(err)
	}

	err = SaveJson(jsonText)
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateJson(jsonText string) (err error) {
	var js jsonData
	err = json.Unmarshal([]byte(jsonText), &js)
	if err != nil {
		return err
	}
	return nil
}

// func getJson() (jsonData []jsonData) {

// 	fileBytes, err := ioutil.ReadFile("./data.json")

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = json.Unmarshal(fileBytes, &jsonData)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return jsonData
// }

func SaveJson(jsonText string) (err error) {
	bytes := []byte(jsonText)
	timeNow := time.Now().String()

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)

	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			fileBytes, err := ioutil.ReadFile(file.Name())
			if err != nil {
				return err
			}
			if strings.Compare(string(fileBytes), jsonText) == 0 {
				return fmt.Errorf("already exist")
			}
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf("./%s.json", timeNow), bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
