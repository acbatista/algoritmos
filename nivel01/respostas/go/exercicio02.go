package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite um valor inteiro: ")

	input, _ := reader.ReadString('\n')

	valor, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Entrada inválida. Por favor, insira um número inteiro.")
	}

	fmt.Printf("O antecessor de %d é %d\n", valor, valor-1)
}
