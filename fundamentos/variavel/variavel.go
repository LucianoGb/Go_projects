package main

import "fmt"

func main()  {
	//maneira shorthand para declarar a variável. ":=" é chamado de gopher
	inteiro := 10
	texto := "Luciano"
	logico := true

	fmt.Println("Inteiro:", inteiro, " Texto:", texto, "Logico:", logico)

	// simulando erro de tentar mudar o valor da variavel com o operador gopher

	// inteiro := 20

	// fmt.Println(inteiro)
}