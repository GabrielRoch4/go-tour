package main

import "fmt"

// Tipo vem após a variável
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2, 2))
}
