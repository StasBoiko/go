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
	Print(v interface{})
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

//в сортинг можно было бы конструктор свой добавить, чтобы сервис через New инициализировать вместо просто создания структуры
func NewService(log Logger) *Service {
	return &Service{log: log}
}

func (s *Service) Print(v interface{}) {
	s.log.Println(v)
}

func (s *Service) Sort(dataToSort interface{}) {
	switch t := dataToSort.(type) {
	case []Person:
		sort.Sort(ByAge(t))
	case []Place:
		sort.Sort(ByCity(t))
	default:
		s.log.Println("unexpected type %T", t)
	}
}

func (s *Service) Decode(data []byte) ([]Person, []Place) {
	var persons []Person
	var places []Place
	var jsonData Things

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, nil
	}
	for _, record := range jsonData.Data {
		name, has_name := record["name"]
		age, has_age := record["age"]
		city, has_city := record["city"]
		country, has_country := record["country"]
		if has_name && has_age {
			newPerson := Person{name.(string), int(age.(float64))}
			persons = append(persons, newPerson)
		}
		if has_country && has_city {
			newPlace := Place{city.(string), country.(string)}
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

	logger := log.New(os.Stdout, "INFO: ", 0)

	var hD HumanDecoder
	hD = NewService(logger)
	// hD = &Service{logger}
	personsSlice, placesSlice := hD.Decode(jsonStr)

	hD.Sort(personsSlice)
	hD.Sort(placesSlice)

	hD.Print(personsSlice)
	hD.Print(placesSlice)
}

//еще вариант обработки JSON
type Mixed struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func solutionB(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]Mixed
	if err := json.Unmarshal(jsonStr, &data); err != nil {
		fmt.Println(err)
		return persons, places
	}
	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{item.Name, item.Age})
		} else {
			places = append(places, Place{item.City, item.Country})
		}
	}
	return persons, places
}

//еще вариант обработки JSON
func solutionD(data []byte) ([]Person, []Place) {
	var doc struct {
		Things []json.RawMessage `json:"things"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		log.Fatal(err)
	}
	people, places := make([]Person, 0, len(doc.Things)), make([]Place, 0, len(doc.Things))
	for _, thing := range doc.Things {
		peep := &Person{}
		if peep.unmarshal(thing) {
			people = append(people, *peep)
			continue
		}
		spot := &Place{}
		if spot.unmarshal(thing) {
			places = append(places, *spot)
			continue
		}
		fmt.Printf("unable to unmarshal: %v\n", thing)
	}
	return people, places
}
func (p *Person) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return p.Name != "" && p.Age != 0
}
func (p *Place) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return p.City != "" && p.Country != ""
}


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
	Print(v interface{})
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

//в сортинг можно было бы конструктор свой добавить, чтобы сервис через New инициализировать вместо просто создания структуры
func NewService(log Logger) *Service {
	return &Service{log: log}
}

func (s *Service) Print(v interface{}) {
	s.log.Println(v)
}

func (s *Service) Sort(dataToSort interface{}) {
	switch t := dataToSort.(type) {
	case []Person:
		sort.Sort(ByAge(t))
	case []Place:
		sort.Sort(ByCity(t))
	default:
		s.log.Println("unexpected type %T", t)
	}
}

func (s *Service) Decode(data []byte) ([]Person, []Place) {
	var persons []Person
	var places []Place
	var jsonData Things

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, nil
	}
	for _, record := range jsonData.Data {
		name, has_name := record["name"]
		age, has_age := record["age"]
		city, has_city := record["city"]
		country, has_country := record["country"]
		if has_name && has_age {
			newPerson := Person{name.(string), int(age.(float64))}
			persons = append(persons, newPerson)
		}
		if has_country && has_city {
			newPlace := Place{city.(string), country.(string)}
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

	logger := log.New(os.Stdout, "INFO: ", 0)

	var hD HumanDecoder
	hD = NewService(logger)
	// hD = &Service{logger}
	personsSlice, placesSlice := hD.Decode(jsonStr)

	hD.Sort(personsSlice)
	hD.Sort(placesSlice)

	hD.Print(personsSlice)
	hD.Print(placesSlice)
}

//еще вариант обработки JSON
type Mixed struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func solutionB(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]Mixed
	if err := json.Unmarshal(jsonStr, &data); err != nil {
		fmt.Println(err)
		return persons, places
	}
	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{item.Name, item.Age})
		} else {
			places = append(places, Place{item.City, item.Country})
		}
	}
	return persons, places
}

//еще вариант обработки JSON
func solutionD(data []byte) ([]Person, []Place) {
	var doc struct {
		Things []json.RawMessage `json:"things"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		log.Fatal(err)
	}
	people, places := make([]Person, 0, len(doc.Things)), make([]Place, 0, len(doc.Things))
	for _, thing := range doc.Things {
		peep := &Person{}
		if peep.unmarshal(thing) {
			people = append(people, *peep)
			continue
		}
		spot := &Place{}
		if spot.unmarshal(thing) {
			places = append(places, *spot)
			continue
		}
		fmt.Printf("unable to unmarshal: %v\n", thing)
	}
	return people, places
}
func (p *Person) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return p.Name != "" && p.Age != 0
}
func (p *Place) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return p.City != "" && p.Country != ""
}
