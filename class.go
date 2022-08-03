//constructor es el que instancia los objetos
package main

import "fmt"

type Empleado struct {
	Id   int
	Name string
}

//manera recomendada de emular los constructores
//recordar que go trata los structs como si fueran copias, es decir trabaja con el apuntador de memora &
//crea funcion que recibe los parametros del struct empleados y devuelve datos con apuntador *
func NewEmpleado(id int, nombre string) *Empleado {
	return &Empleado{
		Id:   id,
		Name: nombre,
	}
}

func main() {
	//instanciar manera 1
	e := Empleado{}
	fmt.Println(e)
	//instanciar manera 2 explicita
	e2 := Empleado{
		Id:   1,
		Name: "Nombrexd",
	}
	fmt.Println(e2)
	//instanciar manera 3
	e3 := new(Empleado)
	fmt.Println(*e3)
	//manera 4 recomendada
	e4 := NewEmpleado(3, "Stiven")
	fmt.Println(*e4)
}
