
//*Ejercicio:

 func main() {
	numero := 7
	resultado := calcularDoble(numero)
	fmt.Print("El doble de %d es %d\n", numero, resultado)
 }

 func calcularDoble(num int) int {
	return num * 2
 }

 package main

 //importación de paquetes
//"fmt" es para formatear e imprimir el texto
 //se importa el paquete "bufio"
 impor (
	"bufio"
	"fmt"
	"os"
	"strconv" //sirve para convertir de texto a numero
	"strings"
 )


//main-g0

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
		fmt.Printf("Numero %d\n", i) //permite el formateo con %d para numeros
	}
}

//Nuevo codigo paso#3

//creacion de un lector de entrada
reader := bufio.NewReader(os.Stdin)

fmt.Print("Por favor, ingresa un numero para calcular su doble: ")

//Se lee hasta que el usuario presione Enter
entrada, _ := reader .reader.ReadString('\n')



//definimos la nueva función impmensaje
//toma el argumento mensaje de tipo string
func imprimirMensaje(mensaje string) {
	fmt.Println("Mensaje de la función:", mensaje)
 } 

 