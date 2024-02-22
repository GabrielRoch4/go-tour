package main

import (
	"fmt"
)

func main() {
	num := 5

	if num%2 == 0 {
		fmt.Println("O número", num, "é par")
	} else {
		fmt.Println("O número", num, "não é par")
	}
}
