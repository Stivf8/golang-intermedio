package main

//Las interfaces dicen que debe ser implementado mas no dice el como
import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

//funcion que reutiliza funcion getMessage para FullTimeEmployee
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

//funcion que reutiliza funcion getMessage para TemporaryEmployee
func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

//definicion interfaz que dice que debe ser implementado getMessage
type PrintInfo interface {
	getMessage() string
}

//se debe definir el como se va a implementar la interfaz para el metodo getMessage
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
	//GetMessage(ftEmployee)
}
