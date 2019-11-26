package main

import (
        "testing"
)

func TestFoo(t *testing.T) {
        a := 10
        expected := 45
        t.Logf("Start number: %d", a)
        x := Foo(a)
        if x != expected {
                t.Errorf("Incorrect output: %d", x)
        } else {
                t.Logf("Correct output: %d", x)
        }
}
