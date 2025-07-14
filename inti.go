package main //punto de entrada

import (
	"fmt" // paquete de formateo
	"reflect"
)

//comentario simple
/*
Comentarios de doble lineas
*/

func main() { //funcion principal
	fmt.Print("print con el modulo fmt \n") //print sobre la misma linea
	fmt.Println("print con salto de linea")
	print("Hola Mundo! con Go \n")

	//Variables

	var myString string = "My String \n"
	fmt.Print(myString)

	myString = "Other String \n"

	fmt.Print(myString)

	// Tipos de datos

	var name string = "Emanuel"
	fmt.Println(name)

	var age int32 = 29 //multiples tipos de int, int32, int64, etc
	fmt.Println(age)

	fmt.Printf("%s %d", name, age) //Print con formate, Println lo puede manejar
	fmt.Println(name, age)

	//Mostrar tipos de datos
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(name))

	//Numeros con numero flotante

	var pi float32 = 3.14
	fmt.Println(pi)

	//Booleanos

	var myBool bool = false
	println(myBool)
	myBool = true
	println(myBool)

	//auto init vars :=

	myString2 := "auto init vars"
	println(myString2)

	//constantes

	const myRol = "Full Stack"
	println((myRol))

	//Control de flujos

	if age >= 18 {
		println("Soy mayor de edad")
	} else {
		println("No soy mayor de edad")
	}

	//Los IF's no encapsulan las comparaciones en parentesis

	/*
		Operadores de Go
		Aritméticos:
		+ - * / % (suma, resta, multiplicación, división, módulo)
		Comparación:
		== != < > <= >=
		Lógicos:
		&& || ! (AND, OR, NOT)
		Asignación:
		= += -= *= /= %= &= |= ^= <<= >>=
		Incremento/Decremento:
		++ --
		Bitwise:
		& | ^ << >> &^ (AND, OR, XOR, shift izq/der, AND NOT)
		Punteros:
		& (dirección) * (desreferencia)
		Channel:
		<- (enviar/recibir)
		Declaración corta:
		:= (declarar e inicializar)
	*/

	//Arrays
}
