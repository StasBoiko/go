
How to read data from std input?
```
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// when you hit enter - the data is read
	scanner.Scan()
	// this will read data from the scanner
	yourData := scanner.Text()
	// will print it to stdin
	fmt.Println(yourData)
}

```

Try to use it in your tasks

***0***

Print a chessboard with the specified dimensions of height and width, according to the example height - 4 and wigth 6:

```
^  ^  ^  ^  ^  ^
  ^  ^  ^  ^  ^  ^
^  ^  ^  ^  ^  ^
  ^  ^  ^  ^  ^  ^
```

Input parameters: height, width, symbol to print
Output: board

***1)***


scan array string

use strings split to split numbers and Atoi

add logic that counts the number of positive even numbers in the array and prints it.

        Example of input data: 30,-1,-6,90,-6,22,52,123,2,35,6

***2)***

1. Enter valid bank card number
2. Validate it
3. Print string with all strings covered except for the last 4

        Input example: 4539 1488 0343 6467


***3)***
Closures

```
package main

import "fmt"

/*
Go supports anonymous functions, which can form closures.

A closure is a special type of anonymous function that
references variables declared outside of the function itself. It is similar to accessing global variables which are available before the declaration of the function.

Anonymous functions are useful when you want to define a function inline without having to name it.

We call intSeq, assigning the result (a function) to nextInt.
This function value captures its own i value, which will be updated each time we call nextInt.
 */

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt()) // 1 
	fmt.Println(nextInt()) // 2 
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}

```

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, …)
The next number is found by adding up the two numbers before it:

the 2 is found by adding the two numbers before it (1+1),
the 3 is found by adding the two numbers before it (1+2),
the 5 is (2+3)
