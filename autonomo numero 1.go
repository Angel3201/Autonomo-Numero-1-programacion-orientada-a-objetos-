package main

import "fmt"

func main() {

	var opcion int

	for {

		fmt.Println("=== Sistema de Gestión de Libros ===")
		fmt.Println("1. Registrar libro")
		fmt.Println("2. Mostrar libros")
		fmt.Println("3. Salir")

		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			registrarLibro()

		case 2:
			mostrarLibros()

		case 3:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción no válida")
		}
	}
}

func registrarLibro() {
	fmt.Println("Registro de libro")
}

func mostrarLibros() {
	fmt.Println("Lista de libros")
}
