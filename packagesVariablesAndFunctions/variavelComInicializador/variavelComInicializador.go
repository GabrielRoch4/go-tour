package main

import "fmt"

/*
	A declaração var pode incluir inicializadores, uma por variável.
	Se um inicializador está presente, o tipo pode ser omitido; a variável terá o tipo do inicializador.	
*/

var i, j int = 1, 2

var num = 90

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java, num)
}
