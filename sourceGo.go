package main

import (
	"fmt"
)

func getName() string {
	return "World!"
}

func main() {
	fmt.Println("Hello", getName())
}
