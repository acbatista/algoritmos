package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("Antes da troca: A = %d, B = %d\n", a, b)

	// Troca dos valores sem variável auxiliar
	// a = a + b
	// b = a - b
	// a = a - b

	// Troca simultânea em uma única linha
	a, b = b, a

	fmt.Printf("Após a troca: A = %d, B = %d\n", a, b)

}
