package main

import "fmt"

func main() {
	fmt.Println("I can call count, count is", count)
}

// rules order deklarasi variable untuk level package tidak berlaku
var count int = 5
