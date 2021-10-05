package main

// import "fmt"
// // Exercise: Loops and Functions
// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2*z)
// 	}

// 	return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }

// Exercise: Slices
// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {
// 	p := make([][]uint8, dy)
//     for y := range p {
//         p[y] = make([]uint8, dx)
//         for x := range p[y] {
// 			p[y][x] = uint8(x^y)
//         }
//     }
//     return p
// }

// func main() {
// 	pic.Show(Pic)
// }

// Exercise: Maps
// func WordCount(s string) map[string]int {
// 	m := make(map[string]int)
// 	for _, v := range strings.Fields(s) {
// 		m[v] += 1
// 	}
// 	return m
// }

// Exercise: Fibonacci closure
// func fibonacci() func(nextTerm int) int {
// 	return func(nextTerm int) int {
// 		return nextTerm
// 	}
// }

// func main() {
// 	t1 := 0
// 	t2 := 1
// 	nextTerm := t1 + t2
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f(nextTerm))
// 		t1 = t2
// 		t2 = nextTerm
// 		nextTerm = t1 + t2
// 	}
// }

// func main() {
// 	m := make(map[string]int)

// 	for _, v := range strings.Fields("i am good") {
// 		m[v] = len(v)
// 	}
// 	fmt.Println(m)
// }

// fibonacci is a function that returns
// a function that returns an int.
// func fibonacci() func(nextTerm int) int {
// 	return func(nextTerm int) int {
// 		return nextTerm
// 	}
// }

// func main() {
// 	t1 := 0
// 	t2 := 1
// 	nextTerm := t1 + t2
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f(nextTerm))
// 		t1 = t2
// 		t2 = nextTerm
// 		nextTerm = t1 + t2
// 	}
// }

// Exercise: Stringers
// type IPAddr [4]byte

// func (ip IPAddr) String() string {
// 	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
// }

// func (key StasMap) String() string {
// 	return fmt.Sprintf("%v", "myNewKey")
// }

// type StasMap string

// func main() {
// 	hosts := map[StasMap]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for key, ip := range hosts {
// 		fmt.Printf("%v %v \n", key, ip)
// 	}
// }

// Не получилось:
// Exercise: Readers
// package main

// import (
// 	"golang.org/x/tour/reader"
// )

// type MyReader struct{}

// // TODO: Add a Read([]byte) (int, error) method to MyReader.
// func (r MyReader) Read(b []byte) (int, error) {
// 	b[0] = 'A'
// 	return 0, nil
// }

// func main() {
// 	reader.Validate(MyReader{})
// }
