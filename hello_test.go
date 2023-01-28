package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World! FAIL"
    if got != want {
        t.Errorf("got %q, want %q\n", got, want)
    }
}
