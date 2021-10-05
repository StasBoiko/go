package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

//tak0
// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	var height string
// 	fmt.Print("Enter height: ")
// 	scanner.Scan()
// 	height = scanner.Text()

// 	var width string
// 	fmt.Print("Enter width: ")
// 	scanner.Scan()
// 	width = scanner.Text()

// 	fmt.Print("Enter symbol: ")
// 	scanner.Scan()
// 	symbol := scanner.Text()

// 	intWidth, err := strconv.Atoi(width)
// 	if err != nil {
// 		panic(err)
// 	}
// 	intHeight, err := strconv.Atoi(height)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var firstLine string
// 	for i := 0; i < intWidth; i++ {
// 		firstLine += symbol + " "
// 	}

// 	for j := 0; j < intHeight; j++ {
// 		if j%2 == 0 {
// 			fmt.Printf("%q\n  ", firstLine)
// 		} else {
// 			fmt.Printf("%q\n", firstLine)
// 		}
// 	}
// }

// Task 1
// scan array string
// use strings split to split numbers and Atoi
// add logic that counts the number of positive even numbers in the array and prints it.
// Example of input data: 30,-1,-6,90,-6,22,52,123,2,35,6
// func main() {
// 	s := strings.Split("30,-1,-6,90,-6,22,52,123,2,35,6", ",")
// 	var count int
// 	for _, n := range s {
// 		intN, _ := strconv.Atoi(n)
// 		if intN > 0 {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }

// Task 2
// Enter valid bank card number
// Validate it
// Print string with all strings covered except for the last 4
// Input example: 4539 1488 0343 6467
// func main() {
// 	s := strings.Split("4539 1488 0343 6467", " ")
// 	s[3] = "xxxx"
// 	fmt.Println(s)
// }

// Task 3
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, …) The next number is found by adding up the two numbers before it:
// the 2 is found by adding the two numbers before it (1+1), the 3 is found by adding the two numbers before it (1+2), the 5 is (2+3)
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

// Task 4

// Check if a given number or part of this number is a palindrome.
//For example, the number 1234437 is not a palindrome, but its part 3443 is a palindrome.
//Numbers less than 10 are counted as invalid input.

// Easy level:
// Find if a number is a palindrome or not

// Mid level:
// You need to find only first subpalindrome

// Hard:
// Find all subpalindromes
// Input parameters: number

// Output: the palindrome/s extracted from the number, or 0 if the extraction failed or no palindromes was found.
// func main() {
// Easy level:
// Find if a number is a palindrome or not
// s := "123456"
// s := "12345678901234";
// s := "123456654321";
// s := "12219876543223"
// easy := "1221"
// foo := IsPalindrome(easy)

// fmt.Println(len(s))
// if len(s) < 10 {
// 	fmt.Println("math: square root of negative number %g")
// }

// intS2, err := strconv.Atoi(s2)
// intS3, err := strconv.Atoi(s3)
// intS4, err := strconv.Atoi(s4)

// s1 := strings.Split(s1, " ")
// s2 := strings.Split(s2, " ")
// s3 := strings.Split(s3, " ")
// s4 := strings.Split(s4, " ")
// s[3] = "xxxx"
// 	fmt.Println(foo)
// }

// Easy level:
// Find if a number is a palindrome or not
// func IsPalindrome(str string) bool {
// 	reversedStr := ""
// 	for i := len(str) - 1; i >= 0; i-- {
// 		reversedStr += string(str[i])
// 	}
// 	for i := range str {
// 		if str[i] != reversedStr[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var number, remainder, temp int
// 	var reverse int = 0

// 	fmt.Print("Enter any positive integer : ")
// 	fmt.Scan(&number)

// 	// fmt.Println(number)
// 	temp = number

// 	// For Loop used in format of While Loop
// 	for {
// 		remainder = number % 10
// 		// fmt.Println(number % 10)
// 		reverse = reverse*10 + remainder

// 		number /= 10

// 		if number == 0 {
// 			break // Break Statement used to exit from loop
// 		}
// 	}

// 	fmt.Println(number)
// 	fmt.Println(reverse)
// 	if temp == reverse {
// 		fmt.Printf("%d is a Palindrome", temp)
// 	} else {
// 		fmt.Printf("%d is not a Palindrome", temp)
// 	}

// }

// Task 5

// There are 2 ways to count lucky tickets:
// Simple - if a six-digit number is printed on the ticket,
// and the sum of the first three digits is equal to the sum of the last three digits, then this ticket is lucky.

// Hard - if the sum of the even digits of the ticket is equal to the sum of the odd digits of the ticket,
// then the ticket is considered lucky.
// Determine programmatically which variant of counting lucky tickets will give them a greater number at a given interval.

// Task: Calculate how many tickets are lucky within provided min and max tickets
// Input parameters: 2 values min digit of ticket and max digit of ticket
// Output: information about the winning method and the number of lucky tickets for each counting method.
// Input numbers contains exactly 6 digits. Not more or less

// Example:

// Min: 120123
// Max: 320320

// --Result--
// EasyFormula: 11187
// HardFormula: 5790

// Simple:
// func main() {
// 	var countLucky int
// 	// Min := 120123
// 	min := 120123
// 	max := 320320
// 	for {
// 		min++
// 		strMin := strconv.Itoa(min)
// 		var firstThree int
// 		var lastThree int
// 		var intVal int
// 		for k, val := range strMin {
// 			if k < 3 {
// 				intVal, _ = strconv.Atoi(string(val))
// 				firstThree += intVal
// 			} else {
// 				intVal, _ = strconv.Atoi(string(val))
// 				lastThree += intVal
// 			}
// 		}
// 		if firstThree == lastThree {
// 			countLucky++
// 		}
// 		if min == max {
// 			break
// 		}
// 	}
// 	fmt.Println(countLucky)
// }

// Hard:
// func main() {
// 	var countLucky int
// 	min := 120123
// 	max := 320320
// 	for {
// 		min++
// 		strMin := strconv.Itoa(min)
// 		var evenNumbers int
// 		var oddNumbers int
// 		var intVal int
// 		for _, val := range strMin {
// 			intVal, _ = strconv.Atoi(string(val))
// 			if intVal%2 == 0 {
// 				evenNumbers += intVal
// 			} else {
// 				oddNumbers += intVal
// 			}
// 		}
// 		if evenNumbers == oddNumbers {
// 			countLucky++
// 		}

// 		if min == max {
// 			break
// 		}
// 	}
// 	fmt.Println(countLucky)
// }

// func main() {

// 	var buf bytes.Buffer
// 	reader := bufio.NewReader(os.Stdin)

// 	// for {
// 	line, err := reader.ReadString('\n')
// 	if err != nil {
// 		// if err == io.EOF {
// 		// 	buf.WriteString(line)
// 		// 	// break // end of the input
// 		// } else {
// 			fmt.Println(err.Error())
// 			os.Exit(1) // something bad happened
// 		// }
// 	}
// 	buf.WriteString(line)

// 	// }

// 	fmt.Printf("valid json? %v\n", json.Valid(buf.Bytes()))

// 	type MedicalRecord struct {
// 		Name string `json:"name"`
// 		Age  int    `json:"age"`
// 	}

// 	var record MedicalRecord
// 	err = json.Unmarshal(buf.Bytes(), &record)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1) // something bad happened
// 	}

// 	fmt.Printf("name: %s, age: %d\n", record.Name, record.Age)
// }

// type MedicalRecord struct{}

// func main() {
// 	var record MedicalRecord

// 	err := json.NewDecoder(os.Stdin).Decode(&record)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(record)
// }

type Things struct {
	Data []map[string]interface{} `json:"things"`
}

func main() {

	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	// for {
	line, err := reader.ReadString('\n')
	if err != nil {
		// if err == io.EOF {
		// buf.WriteString(line)
		// break // end of the input
		// } else {
		fmt.Println(err.Error())
		os.Exit(1) // something bad happened
		// }
	}
	buf.WriteString(line)

	// }

	fmt.Printf("valid json? %v\n", json.Valid(buf.Bytes()))

	var jsonData Things

	err = json.Unmarshal([]byte(buf.Bytes()), &jsonData)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // something bad happened
	}
	for _, record := range jsonData.Data {

		fmt.Println(record)
	}

	// var doc map[string][]Mixed
	// var jsonStr = []byte(buf.Bytes())
	// // var record MedicalRecord
	// err = json.Unmarshal(jsonStr, &doc)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1) // something bad happened
	// } else {
	// 	fmt.Println(doc)
	// 	fmt.Println(doc.Data)
	// 	fmt.Println(buf.Bytes)
	// 	fmt.Println([]byte(buf.Bytes()))
	// }

	// foo := doc.Data
	// for i, data := range doc.Data {

	// 	fmt.Printf("sss: %s, ddd: %d, vvv %v\n", i, i, i)
	// 	fmt.Printf("sss: %s, ddd: %d, vvv %v\n", data, data, data)
	// }
	// fmt.Printf("sss: %s, ddd: %d, vvv %v\n", foo, foo, foo)
}
