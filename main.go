
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
    fmt.Println("Hola Mundo Concurrente")

    mensaje := "Este es el taller 0"
    imprimirMensaje(mensaje)

    fmt.Println("Contando hasta 3:")
    for i := 1; i <= 3; i++ {
        fmt.Printf("Numero %d\n", i)
    }

    // Bloque que estaba fuera de funciones ahora está dentro de main
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Por favor, ingresa un numero para calcular su doble: ")
    entrada, _ := reader.ReadString('\n')
    entrada = strings.TrimSpace(entrada)

    numero, err := strconv.Atoi(entrada)
    if err != nil {
        fmt.Println("Error: Debes ingresar un numero válido.", err)
        return
    }

    resultado := calcularDoble(numero)
    fmt.Printf("El doble de %d es %d\n", numero, resultado)
}

 //definimos la nueva función impmensaje
//toma el argumento mensaje de tipo string
func imprimirMensaje(mensaje string) {
	fmt.Println("Mensaje de la función:", mensaje)
 } 


 func calcularDoble(num int) int {
	return num * 2
 }

 