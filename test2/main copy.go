package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Things struct {
	Data []map[string]interface{} `json:"things"`
}

type HumanDecoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})
	Print(interface{})
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

func (s *Service) Print(interface{}) {
	// fmt.Println(interface{})
}

func (s *Service) Sort(dataToSort interface{}) {
	switch t := dataToSort.(type) {
	case []Person:
		sort.Sort(ByAge(t))
	case []Place:
		sort.Sort(ByCity(t))
	default:
		fmt.Printf("unexpected type %T", t)

	}
}

func (s *Service) Decode(data []byte) ([]Person, []Place) {
	var persons []Person
	var places []Place
	var jsonData Things

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println(err)
	}
	for _, record := range jsonData.Data {
		_, has_name := record["name"]
		_, has_age := record["age"]
		_, has_city := record["city"]
		_, has_country := record["country"]
		if has_name && has_age {
			var newPerson = Person{string(record["name"].(string)), int(record["age"].(float64))}
			persons = append(persons, newPerson)
		}
		if has_country && has_city {
			var newPlace = Place{string(record["city"].(string)), string(record["country"].(string))}
			places = append(places, newPlace)
		}
	}

	return persons, places
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByCity []Place

func (a ByCity) Len() int           { return len(a) }
func (a ByCity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCity) Less(i, j int) bool { return len(a[i].City) < len(a[j].City) }

func main() {
	logger := log.New(os.Stdout, "INFO: ", 0)

	var jsonStr = []byte(`
	{
		"things": [
			{
				"name": "Alice",
				"age": 37
			},
			{
				"city": "Ipoh",
				"country": "Malaysia"
			},
			{
				"name": "Bob",
				"age": 36
			},
			{
				"city": "Northampton",
				"country": "England"
			},
			{
				"name": "Albert",
				"age": 3
			},
			{
				"city": "Dnipro",
				"country": "Ukraine"
			},
			{
				"name": "Roman",
				"age": 32
			},
			{
				"city": "New York City",
				"country": "US"
			}
		]
	}`)

	var hD HumanDecoder
	hD = &Service{logger}
	personsSlice, placesSlice := hD.Decode(jsonStr)

	hD.Sort(personsSlice)
	hD.Sort(placesSlice)

	logger.Println(personsSlice)
	logger.Println(placesSlice)
}
