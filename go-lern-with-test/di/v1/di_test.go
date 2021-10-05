package main

import (
	"bytes"
	"testing"
	"fmt"
)

func TestGreet(t *testing.T) {
    buffer := &bytes.Buffer{}
	
    Greet(buffer, "Chris")

    got := buffer.String()
	fmt.Println(got)
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
