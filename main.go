
//*Ejercicio:
 package main

 
 //importación de paquetes
//"fmt" es para formatear e imprimir el texto
 //se importa el paquete "bufio"
 import (
	"bufio"
	"fmt"
	"os"
	"strconv" //sirve para convertir de texto a numero
	"strings"
 )

 
//funcion main es el punto de entrada del programa
func main() {
	//se usa el paquete fmt 
	fmt.Println("Hola Mundo Concurrente")

	//se declara una variable para contar
	mensaje := "Este es el taller 0"

	//llamamos la funcion y definimos mas abajo la funcion
	imprimirMensaje(mensaje)

	fmt.Println("Contando hasta 3:")
	for i := 1; i <= 3; i++ {
		fmt.Println("Numero %d\n", i) //permite el formateo con %d para numeros
	}
}

//Nuevo codigo paso#3

//creacion de un lector de entrada
reader := bufio.NewReader(os.Stdin)

fmt.Print("Por favor, ingresa un numero para calcular su doble: ")

//Se lee hasta que el usuario presione Enter
entrada, _ := reader .reader.ReadString('\n')

//Limpiar el caracter de nueva linea y espacios
entrada = strings.TrimSpace(entrada)

//convertie el string a # entero
numero, err := strconv.Atoi(entrada)
//Manejo de errores
if err != nil {
	fmt.Println("Error: Debes ingresar un numero válido.", err)
	return
}

resultado := calcularDoble(numero)
fmt.Printf("El doble de %d es %d\n", numero, resultado)


 //definimos la nueva función impmensaje
//toma el argumento mensaje de tipo string
func imprimirMensaje(mensaje string) {
	fmt.Println("Mensaje de la función:", mensaje)
 } 

 

 func calcularDoble(num int) int {
	return num * 2
 }





//   func main() {
// 	numero := 7
// 	resultado := calcularDoble(numero)
// 	fmt.Print("El doble de %d es %d\n", numero, resultado)
//  }

 