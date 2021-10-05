// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"sort"
// )

// // Write a function that parses given input data into 2 list of structs and sorts them.
// // Persons should be sorted by age and Country data should be sorted by City name length
// // Define service structure that will implement HumanDecoder interface.
// // Inject into the service logger that implements our logger interface
// // Use dependency injection to inject logger that will print data
// // Interface structure to be used

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
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// type Place struct {
// 	City    string `json:"city"`
// 	Country string `json:"country"`
// }

// type Data struct {
// 	Person
// 	Place
// }

// type MainData struct {
// 	Things []Data
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%s: %d", p.Name, p.Age)
// }

// func Decode(data []byte) ([]Person, []Place) {
// 	someData := MainData{}
// 	err := json.Unmarshal([]byte(data), &someData)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//primes := [6]int{2, 3, 5, 7, 11, 13}

// 	//var s []int
// 	//fmt.Println(someData)
// 	//a := make(map[string]interface{})
// 	//	var foo []Data = someData.Things
// 	//personCheck := Person{}
// 	// personCheck := someData.Things[0].Person.Age
// 	Sort(someData.Things)
// 	for i := range someData.Things {

// 		//j, _ := json.Marshal(val)
// 		//fmt.Println(j)

// 		// data := map[string]string{
// 		// 	"a": "aaa",
// 		// 	"b": "bbb",
// 		// }
// 		//someData.Things[i].Person = data
// 		if someData.Things[i].Person.Age != 0 {

// 			fmt.Println(someData.Things[i].Person.Age)
// 		}
// 		if someData.Things[i].Person.Name != "" {
// 			fmt.Println(someData.Things[i].Person.Name)
// 		}
// 		// fmt.Println(someData.Things[i].Person.Name)
// 		//fmt.Println(someData.Things[i].Person)
// 		//fmt.Println(someData.Things[i].Place)
// 		//fmt.Println(i)
// 		//fmt.Println(val)
// 	}
// 	return person, place
// }
// func Sort(dataToSort interface{}) {

// }
// // func Print(interface{}) {

// // }

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

//     Decode(jsonStr)

// }

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

}
