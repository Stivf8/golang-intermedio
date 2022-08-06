package main

import (
	"fmt"
	"time"
)

// Funciones anónimas
func main() {
	//funcion anonima que dice hello, se usa cuando estamos muy seguros que solo la vamos ausar una vez
	func() {
		println("Hello")
	}()

	x := 5
	//otra funcion anonima
	y := func() int {
		return x * 2
	}()
	//imprimos funcion anonima
	fmt.Println(y)
	//creamos canal que espera un int
	c := make(chan string)
	//ejecutamos goRoutine que ejecuta funcion anonima, luego responde al canal c creado
	go func() {
		fmt.Println("Starting function")
		time.Sleep(2 * time.Second)
		fmt.Println("Finishing function")
		c <- "Go Routine Finalizada"
	}()
	//imprimimos el canal, para que la rutina actual espere a la goRoutine
	//tambien puede ser asi
	//<-c
	fmt.Println(<-c)
}
