package main

import (
	"errors"
	"fmt"
)

// Interfaz
type Mostrable interface {
	MostrarInformacion()
}

// Estructura Libro (Encapsulación)
type Libro struct {
	titulo     string
	autor      string
	categoria  string
	disponible bool
}

// Constructor
func NuevoLibro(titulo, autor, categoria string) Libro {
	return Libro{
		titulo:     titulo,
		autor:      autor,
		categoria:  categoria,
		disponible: true,
	}
}

// Métodos Getters
func (l Libro) ObtenerTitulo() string {
	return l.titulo
}

func (l Libro) ObtenerAutor() string {
	return l.autor
}

func (l Libro) ObtenerCategoria() string {
	return l.categoria
}

// Método para mostrar información
func (l Libro) MostrarInformacion() {
	fmt.Println("---------------------------")
	fmt.Println("Título:", l.titulo)
	fmt.Println("Autor:", l.autor)
	fmt.Println("Categoría:", l.categoria)

	if l.disponible {
		fmt.Println("Estado: Disponible")
	} else {
		fmt.Println("Estado: Prestado")
	}
	fmt.Println("---------------------------")
}

// Método para cambiar disponibilidad
func (l *Libro) CambiarDisponibilidad(estado bool) {
	l.disponible = estado
}

// Estructura de datos
var catalogo []Libro

func main() {

	var opcion int

	for {

		fmt.Println("\n=== Sistema de Gestión de Libros ===")
		fmt.Println("1. Registrar libro")
		fmt.Println("2. Mostrar libros")
		fmt.Println("3. Buscar libro")
		fmt.Println("4. Registrar préstamo")
		fmt.Println("5. Registrar devolución")
		fmt.Println("6. Salir")

		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			err := registrarLibro()
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			mostrarLibros()

		case 3:
			buscarLibro()

		case 4:
			registrarPrestamo()

		case 5:
			registrarDevolucion()

		case 6:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción no válida")
		}
	}
}

// Manejo de errores
func registrarLibro() error {

	var titulo string
	var autor string
	var categoria string

	fmt.Print("Título: ")
	fmt.Scan(&titulo)

	fmt.Print("Autor: ")
	fmt.Scan(&autor)

	fmt.Print("Categoría: ")
	fmt.Scan(&categoria)

	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}

	libro := NuevoLibro(titulo, autor, categoria)
	catalogo = append(catalogo, libro)

	fmt.Println("Libro registrado correctamente")
	return nil
}

func mostrarLibros() {

	if len(catalogo) == 0 {
		fmt.Println("No existen libros registrados")
		return
	}

	fmt.Println("\n=== Catálogo de Libros ===")

	for _, libro := range catalogo {
		libro.MostrarInformacion()
	}
}

func buscarLibro() {

	var titulo string

	fmt.Print("Ingrese el título a buscar: ")
	fmt.Scan(&titulo)

	for _, libro := range catalogo {

		if libro.ObtenerTitulo() == titulo {

			fmt.Println("Libro encontrado:")
			libro.MostrarInformacion()
			return
		}
	}

	fmt.Println("Libro no encontrado")
}

func registrarPrestamo() {

	var titulo string

	fmt.Print("Título del libro a prestar: ")
	fmt.Scan(&titulo)

	for i := range catalogo {

		if catalogo[i].ObtenerTitulo() == titulo {

			if !catalogo[i].disponible {
				fmt.Println("El libro ya está prestado")
				return
			}

			catalogo[i].CambiarDisponibilidad(false)
			fmt.Println("Préstamo registrado correctamente")
			return
		}
	}

	fmt.Println("Libro no encontrado")
}

func registrarDevolucion() {

	var titulo string

	fmt.Print("Título del libro a devolver: ")
	fmt.Scan(&titulo)

	for i := range catalogo {

		if catalogo[i].ObtenerTitulo() == titulo {

			catalogo[i].CambiarDisponibilidad(true)
			fmt.Println("Devolución registrada correctamente")
			return
		}
	}

	fmt.Println("Libro no encontrado")
}
