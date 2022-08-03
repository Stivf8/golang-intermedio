package main

import "fmt"

//la herencia en go es la composicion.

type Persona struct {
	Nombre string
	Edad   int
}

type Empleado struct {
	Id int
}

//como decimos que el FullTiempoEmpleado tiene propiedades de Persona y Propiedades de Empleado
type FullTiempoEmpleado struct {
	Persona
	Empleado
}

func main() {
	ftEmpleado := FullTiempoEmpleado{}
	ftEmpleado.Nombre = "Stiven"
	ftEmpleado.Edad = 21
	ftEmpleado.Id = 1
	fmt.Println(ftEmpleado)
}
