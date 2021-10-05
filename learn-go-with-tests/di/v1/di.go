// package main

// import (
// 	"fmt"
// 	"bytes"
// 	// "io"
// 	// "os"
// )

// Greet sends a personalised greeting to writer.
// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func Greet(writer *bytes.Buffer, name string) {
//     fmt.Fprintf(writer, "Hello, %s", name)
// }

// func main() {
// 	Greet(os.Stdout, "Elodie")
// }

package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, "world")
}

func main() {
    log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
