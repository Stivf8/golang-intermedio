package main

import (
	"fmt"
	"strconv"
)

func main2() {

	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Capturando valor y error
	myValue, err := strconv.ParseInt("XD", 0, 64)

	// Validando errores.
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Mapa clave valor.
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice de enteros.
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	//agregar valor nuevo al slice se usa append
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
}
