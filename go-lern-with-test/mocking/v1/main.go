package main

import (
	"fmt"
	"io"
	"time"
	// "os"
)

// Countdown prints a countdown from 3 to out.
// func Countdown(out io.Writer) {
// 	fmt.Fprint(out, "3")
// }

const finalWord = "Go!"
const countdownStart = 3

// type Sleeper interface {
//     Sleep()
// }

type SpySleeper struct {
    Calls int
}

// func (s *SpySleeper) Sleep() {
//     s.Calls++
// }

func Countdown(out io.Writer, sleeper *SpySleeper) {
    for i := countdownStart; i > 0; i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}

// func main() {
// 	Countdown(os.Stdout)
// }
