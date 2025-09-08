//main-go

package main //paqute "main"

//importación de paquetes
//"fmt" es para formatear e imprimir el texto
import "fmt"

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

//definimos la nueva función impmensaje
//toma el argumento mensaje de tipo string
func imprimirMensaje(mensaje string) {
	fmt.Println("Mensaje de la función:", texto)
 } 