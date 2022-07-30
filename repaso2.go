package main

//repaso general:
//GoRotines y apuntadores
import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Capturando valor y error
	myValue, err := strconv.ParseInt("NaN", 0, 64)

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
	} //agregar valor con append al slices
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	//el canal monitoreara al goRoutine, que envia tipo entero
	c := make(chan int)
	//luego enviamos c a la goRoutine para que tenga que responder(el canal como parametro)
	go doSomething(c)
	//se espera hasta que la goRoutine responda el 1
	<-c

	//Apuntadores
	g := 25
	fmt.Println(g) // imprime el valor entero 25
	h := &g
	fmt.Println(h) // imprimer la direccion de memoria.
	i := *h
	fmt.Println(i) // Imprime el valor por de g
}

//funcion que recibe el valor c tipo channel y duerme 3 segundos
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	//aqui enviara al canal c el valor de 1 para confirmar
	c <- 1
}
