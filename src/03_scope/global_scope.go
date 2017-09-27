package main

import "fmt"

var count int = 10

func main() {
    fmt.Println("total ", count)
    foo()
}

func foo() {
    fmt.Println("total outside ", count)
}