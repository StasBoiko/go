Write a program that parses given input data into 2 list of structs and sorts them. Persons should be sorted by age and Country data should be sorted by City name length
Define service structure that will implement `Decoder` interface. Inject into the service logger that implements our logger interface

```
// Data to decode
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

```

```
type Decoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})

	Print(interface{})
	Printlen(persons []Person, places []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

func main() {
    // logger to Inject 
    logger := log.New(os.Stdout, "INFO: ", 0)
	
    // rest of the code
   
}

```


1. Define structs to parse code into
2. Define service struct with logger interface inside 
2. Parse json
3. Put data into corresonding slices
4. Pritn slices len
5. Write sort functions for each slice
6. Call sort on resulted slices
7. Print the slices
8. Output should be equal to:

```
INFO: 4 4
INFO: [{Bob 36} {Alice 37} {Roman 32} {Albert 3}]
INFO: [{Ipoh Malaysia} {Dnipro Ukraine} {Northampton England} {New York City US}
```
