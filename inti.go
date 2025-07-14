package main //punto de entrada

import (
	"container/list"
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

	/*
		Tipos de Datos en Go
		Tipos Básicos:

		Enteros: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
		Flotantes: float32, float64
		Complejos: complex64, complex128
		Booleano: bool
		String: string
		Byte: byte (alias de uint8)
		Rune: rune (alias de int32, para Unicode)

		Tipos Compuestos:

		Array: [5]int (tamaño fijo)
		Slice: []int (tamaño dinámico)
		Map: map[string]int (clave-valor)
		Struct: type Person struct { name string }
		Pointer: *int (referencia a memoria)
		Function: func(int) string (funciones como tipos)
		Interface: interface{} (define comportamiento)
		Channel: chan int (comunicación entre goroutines)

		Tipo Especial:

		Nil: valor cero para pointers, slices, maps, channels, functions e interfaces

		Conversiones:
		Go requiere conversiones explícitas: int(3.14), string(65)
		Los tipos no tienen jerarquía - no hay herencia, solo composición.
	*/

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

	var myArray [3]int //al definir longitud se inicializan en 0

	myArray[0] = 2 // se pueden reasignar

	fmt.Println(myArray) // el modulo fmt puede imprimir contenido de arrays , Println() no puede

	var arrayInit [2]int
	arrayInit[0] = 9
	arrayInit[1] = 2

	fmt.Println((arrayInit))

	// Map

	myMap := make(map[string]int) //Definicion de un Map vacio

	myMapInit := map[string]int{"age": 29, "cero": 0, "experienceYears": 1} //Map ya inicializado  con valores
	fmt.Println(myMapInit)

	fmt.Println(myMap)

	myMap["age"] = 29
	myMap["experienceYears"] = 1
	myMap["cero"] = 0

	fmt.Println(myMap)
	fmt.Println("---------------------")

	//LIST

	myList := list.New()
	myList.PushBack(1) //PushBack agrega al final de la lista
	myList.PushBack(2)
	myList.PushBack(3)

	myList.PushFront(0) // PushFront agrega al inicio de la lista

	fmt.Println(myList.Back().Value)  //muestra el ultimo valor
	fmt.Println(myList.Front().Value) //muestra el primer valor

	print("--------------------- \n")
	//Iteraciones / Bucles
	fmt.Println("Bucles")
	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index], index)
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	for index, value := range arrayInit {
		fmt.Println(index, value)
	}

	//Functions

	fmt.Println(myFunction())

	//Estructuras

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Emanuel", 29}

	fmt.Println((myStruct))
}

func myFunction() string {
	return "Mi function"
}
