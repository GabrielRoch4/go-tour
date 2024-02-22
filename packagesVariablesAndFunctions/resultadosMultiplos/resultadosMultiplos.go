package main

import "fmt"

// Função com dois retornos
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	
	a := "Hello"
	b := "World"
	fmt.Println(swap(a, b))
	
}
