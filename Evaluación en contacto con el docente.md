# Sistema-Gestion-Libros-Electronicos

Sistema de Gestión de Libros Electrónicos desarrollado en Go mediante servicios web, permitiendo administrar un catálogo de libros a través de operaciones de registro, consulta, actualización, eliminación, préstamos y devoluciones.

## Funcionalidades implementadas

* Registro de libros mediante servicios web.
* Consulta de todos los libros registrados.
* Búsqueda de libros por título.
* Actualización de información de libros.
* Eliminación de libros del catálogo.
* Registro de préstamos.
* Registro de devoluciones.
* Consulta del estado del servidor.
* Uso de estructuras (structs) para representar libros.
* Uso de slices para almacenar información.
* Implementación de interfaces.
* Encapsulación mediante métodos.
* Manejo de errores.
* Concurrencia mediante Mutex.
* Serialización y deserialización de datos en formato JSON.
* Implementación de servicios web utilizando HTTP.

## Servicios Web Disponibles

| Método | Endpoint         | Descripción                   |
| ------ | ---------------- | ----------------------------- |
| GET    | /libros          | Obtener todos los libros      |
| POST   | /libros          | Registrar un libro            |
| GET    | /libros/{titulo} | Buscar un libro               |
| PUT    | /libros/{titulo} | Actualizar un libro           |
| DELETE | /libros/{titulo} | Eliminar un libro             |
| POST   | /prestamos       | Registrar préstamo            |
| POST   | /devoluciones    | Registrar devolución          |
| GET    | /estado          | Verificar estado del servidor |

## Estructura del proyecto

```text
/main.go
```

## Tecnologías utilizadas

* Go (Golang)
* HTTP
* JSON
* Git
* GitHub

## Ejecución

```bash
go run main.go
```

Servidor disponible en:

```text
http://localhost:8080
```

## Autor

Angel Abrahan Iñiguez Chinchay

## Fecha

Junio 2026
