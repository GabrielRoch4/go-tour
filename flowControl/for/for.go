package main

// a instrução inicial: executada antes da primeira iteração
// a expressão de condição: avaliada antes de cada iteração
// a pós-instrução: executado no final de cada iteração
// não há parenteses

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Teste")
	}

}
