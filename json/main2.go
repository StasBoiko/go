// package main2

// Write a function that parses given input data into 2 list of structs and sorts them.
// Persons should be sorted by age and Country data should be sorted by City name length
// Define service structure that will implement HumanDecoder interface.
// Inject into the service logger that implements our logger interface
// Use dependency injection to inject logger that will print data
// Interface structure to be used

// type HumanDecoder interface {
// 	Decode(data []byte) ([]Person, []Place)
// 	Sort(dataToSort interface{})
// 	Print(interface{})
// }

// type Logger interface {
// 	Println(v ...interface{})
// 	Fatalf(format string, v ...interface{})
// }

// type Service struct {
// 	log Logger
// }

// type Person struct {
// 	name string
// 	age  int
// }

// type Place struct {
// 	city    string
// 	country string
// }

// type Obj struct {
//     Things []interface{}
// }

// type Things struct {
// 	person Person
// 	place Place
// }

// const settingsFilename = "settings.json"

// func (t *Thing) GetLeague() League {
//     sort.Slice(t.Person.age, func(i, j int) bool {
//         return t.Person.age[i].Wins > t.Person.age[j].Wins
//     })
//     return f.league
// }

// // func SortPerson(p Person) {

// // }

// type HumanDecoder interface {
// 	Decode(data []byte) ([]Person, []Place)
// 	Sort(dataToSort interface{})
// 	Print(interface{})
// }

// func Decode(r io.Reader) *Decoder {

// }

// type Person struct {
// 	name string
// 	age  int
// }

// type Place struct {
// 	city    string
// 	country string
// }

// type Thing struct {
// 	Person
// 	Place
// }

// type Person []Person
// type Place []Place

// type Obj struct {
// Data []byte
// Id   string `json:"id"`
// Things json.RawMessage `json:"things"`
// }

// type romanNumeral struct {
// 	Value  uint16
// 	Symbol string
// }

// type romanNumerals []Obj

// func (r romanNumerals) ValueOf(symbols ...byte) uint16 {
// 	symbol := string(symbols)
// 	for _, s := range r {
// 		// if s.Data == symbol {
// 		// return s.Data
// 		// }
// 	}

// 	return 0
// }

// type MainData struct {
// 	Things []struct {
// 		Name    string `json:"name"`
// 		Age     int    `json:"age"`
// 		City    string `json:"city"`
// 		Country string `json:"country"`
// 	} `json:"things"`
// }

// func main() {
// 	var jsonStr = []byte(`
// {
//     "things": [
//         {
//             "name": "Alice",
//             "age": 37
//         },
//         {
//             "city": "Ipoh",
//             "country": "Malaysia"
//         },
//         {
//             "name": "Bob",
//             "age": 36
//         },
//         {
//             "city": "Northampton",
//             "country": "England"
//         },
//  		{
//             "name": "Albert",
//             "age": 3
//         },
// 		{
//             "city": "Dnipro",
//             "country": "Ukraine"
//         },
// 		{
//             "name": "Roman",
//             "age": 32
//         },
// 		{
//             "city": "New York City",
//             "country": "US"
//         }
//     ]
// }`)

	// 	rawDataOut, err := json.MarshalIndent(string(jsonStr), "", "  ")
	// 	fmt.Println(string(rawDataOut))
	// 	if err != nil {
	// 		log.Fatal("JSON marshaling failed:", err)
	// 	}

	// 	err = ioutil.WriteFile(settingsFilename, rawDataOut, 0)
	// 	if err != nil {
	// 		log.Fatal("Cannot write updated settings file:", err)
	// 	}
	// }

	// ------------------------------------

	// package main

	// import (
	// 	"encoding/json"
	// 	"io/ioutil"
	// 	"log"
	// )

	// type Client struct {
	// 	ClientId string
	// 	Date     string
	// }

	// type Settings struct {
	// 	Clients []Client
	// }

	// const settingsFilename = "settings.json"

	// func main() {
	// 	rawDataIn, err := ioutil.ReadFile(settingsFilename)
	// 	if err != nil {
	// 		log.Fatal("Cannot load settings:", err)
	// 	}

	// 	var settings Settings
	// 	err = json.Unmarshal(rawDataIn, &settings)
	// 	if err != nil {
	// 		log.Fatal("Invalid settings format:", err)
	// 	}

	// 	newClient := Client{
	// 		ClientId: "123",
	// 		Date:     "2016-11-17 12:34",
	// 	}

	// 	settings.Clients = append(settings.Clients, newClient)

	// 	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	// 	if err != nil {
	// 		log.Fatal("JSON marshaling failed:", err)
	// 	}

	// 	err = ioutil.WriteFile(settingsFilename, rawDataOut, 0)
	// 	if err != nil {
	// 		log.Fatal("Cannot write updated settings file:", err)
	// 	}
	// }

	// func main() {
	// byt := []byte(`{"id":"someID","data":["str1","str2"]}`)

	// var obj Obj
	// if err := json.Unmarshal(byt, &obj); err != nil {
	//     panic(err)
	// }

	// fmt.Println(obj)

	// byt := jsonStr
	// byt := []byte(`{"things":[{"name": "Alice", "age": 37},{"city": "Northampton", "country": "England"}]}`)
	// byt := []byte(`{"data":[{"my": "obj", "id": 42},{"my": "obj", "id": 42}]}`)

	// v := Obj{jsonStr}

	// fmt.Printf("%+v\n", v.Data)
	// var d []interface{}
	// if err := json.Unmarshal(jsonStr, &d); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", d)

	// var obj Obj

	// obj.Data := jsonStr

	// if err := json.Unmarshal(byt, &obj); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", obj)
	// fmt.Printf("Things: %s\n", obj.Things)

	// var d []interface{}
	// if err := json.Unmarshal(obj.Things, &d); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", d)

	// for _, v := range d {
	// 	// you need a type switch to deterine the type and be able to use most of these
	// 	switch real := v.(type) {
	// 	case string:
	// 		fmt.Println("I'm a string!", real)
	// 	case float64:
	// 		fmt.Println("I'm a number!", real)
	// 	default:
	// 		fmt.Printf("Unaccounted for: %+v\n", v)
	// 	}

	// }
}

// type Object struct {
// 	MapMutex *sync.Mutex
// 	Map      map[string]string
// }

// ЭТО ПРЕОБРАЗОВАТЬ НЕ МОЖЕТ:
// {
// 	"sequence": 3977318850,
// 	"bids": [
// 			  [ "price": "4625.78", "size": "0.80766325", "order_id": 3 ]
// 			],
// 	"asks": [
// 			  [ "price": "4625.79", "size": "3.0154341", "order_id": 3 ]
// 			]
// 	}
