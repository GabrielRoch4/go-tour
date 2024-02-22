package main

import "fmt"

/*
	Dentro de uma função a instrução de atribuição curta := pode ser utilizada
	em lugar de uma declaração var com o tipo implícito.

	Fora de uma função cada estrutura começa com uma palavra-chave
	(var, func, e assim por diante) e não é possível usar o :=.

	https://go-tour-br.appspot.com/basics/10
*/

func main() {
	num1 := 2
	num2 := 3
	soma := num1 + num2

	fmt.Println(soma)
}
