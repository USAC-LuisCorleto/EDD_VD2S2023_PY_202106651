package main

import (
	"EDD_VD2S2023_PY_202106651/Estructuras/ArbolAVL"
	"EDD_VD2S2023_PY_202106651/Estructuras/ColaPrioridad"
	"EDD_VD2S2023_PY_202106651/Estructuras/Listas"
	"EDD_VD2S2023_PY_202106651/Estructuras/MatrizDispersa"
	"fmt"
	"strconv"
)

var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var colaPrioridad *ColaPrioridad.Cola = &ColaPrioridad.Cola{Primero: nil, Longitud: 0}
var matrizDispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var arbolCursos *ArbolAVL.ArbolAVL = &ArbolAVL.ArbolAVL{Raiz: nil}

var loggeado_estudiante string = ""

func main() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("╔══════════════════════════╗")
		fmt.Println("║    Sistema de Inicio     ║")
		fmt.Println("╠══════════════════════════╣")
		fmt.Println("║ 1. Inicio de Sesion      ║")
		fmt.Println("║ 2. Salir                 ║")
		fmt.Println("╚══════════════════════════╝")

		fmt.Print("Selecciona una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			MenuLogin()
		case 2:
			salir = true
			fmt.Println("Saliendo...")
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func MenuLogin() {
	fmt.Print("\033[H\033[2J")
	usuario := ""
	password := ""
	fmt.Println("╔══════════════════════════╗")
	fmt.Println("║     Inicio de Sesión     ║")
	fmt.Println("╠══════════════════════════╣")
	fmt.Print("║ Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("║ Password: ")
	fmt.Scanln(&password)
	fmt.Println("╚══════════════════════════╝")

	if usuario == "ADMIN_202106651" && password == "admin" {
		fmt.Println("Inicio de sesión exitoso.")
		MenuAdmin()
	} else if listaDoble.Buscar(usuario, password) {
		loggeado_estudiante = usuario
		fmt.Println("Inicio de sesión exitoso.")
		fmt.Println("Bienvenido: ", loggeado_estudiante)
		MenuEstudiantes()
	} else {
		fmt.Println("Usuario o contraseña incorrecto")
	}
}

func MenuAdmin() {
	fmt.Print("\033[H\033[2J")
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║        Menú de Administrador       ║")
		fmt.Println("╠════════════════════════════════════╣")
		fmt.Println("║ [1]. Carga de Estudiantes Tutores  ║")
		fmt.Println("║ [2]. Carga de Estudiantes          ║")
		fmt.Println("║ [3]. Cargar de Cursos              ║")
		fmt.Println("║ [4]. Control de Estudiantes        ║")
		fmt.Println("║ [5]. Reportes                      ║")
		fmt.Println("║ [6]. Salir                         ║")
		fmt.Println("╚════════════════════════════════════╝")

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("\033[H\033[2J")
			CargaTutores()
		case 2:
			fmt.Print("\033[H\033[2J")
			CargaEstudiantes()
		case 3:
			fmt.Print("\033[H\033[2J")
			CargaCursos()
		case 4:
			fmt.Print("\033[H\033[2J")
			ControlEstudiantes()
		case 5:
			fmt.Print("\033[H\033[2J")
			MenuReportes()
		case 6:
			salir = true
		}
	}
}

func MenuReportes() {
	fmt.Print("\033[H\033[2J")
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║          Menú de Reportes          ║")
		fmt.Println("╠════════════════════════════════════╣")
		fmt.Println("║ [1]. Reporte de Tutores aceptados  ║")
		fmt.Println("║ [2]. Reporte de Estudiantes        ║")
		fmt.Println("║ [3]. Reporte de Cursos             ║")
		fmt.Println("║ [4]. Reporte de Asignaciones       ║")
		fmt.Println("║ [5]. Salir                         ║")
		fmt.Println("╚════════════════════════════════════╝")

		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			listaDoble.ReporteListaDoble()
		case 2:
			listaDoble.ReporteListaDobleCircular()
		case 3:
			arbolCursos.ReporteArbolAVL()
		case 4:
			matrizDispersa.Reporte("ReporteAsignaciones.jpg")
		case 5:
			salir = true
		}
	}
}

func MenuEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║         Menú de Estudiantes      ║")
		fmt.Println("╠══════════════════════════════════╣")
		fmt.Println("║ [1]. Ver Tutores Disponibles     ║")
		fmt.Println("║ [2]. Asignarse Tutores           ║")
		fmt.Println("║ [3]. Salir                       ║")
		fmt.Println("╚══════════════════════════════════╝")

		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Print("\033[H\033[2J")
			listaDobleCircular.Mostrar()
		case 2:
			AsignarCurso()
		case 3:
			salir = true
		}
	}
}

func AsignarCurso() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Println("Teclee el codigo del curso: ")
		fmt.Scanln(&opcion)
		if arbolCursos.Busqueda(opcion) {
			if listaDobleCircular.Buscar(opcion) {
				TutorBuscado := listaDobleCircular.BuscarTutor(opcion)
				estudiante, err := strconv.Atoi(loggeado_estudiante)
				if err != nil {
					fmt.Println("Error aquí", err)
					break
				}
				matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.Carnet, opcion)
				fmt.Println("El curso", opcion, "se asignó correctamente.")
				break
			} else {
				fmt.Println("No hay tutores disponibles para este curso.")
				break
			}
		} else {
			fmt.Println("No se encontró el curso.")
			break
		}
	}
}

func CargaTutores() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Ingrese el nombre del archivo para la carga de tutores: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerCSV(ruta)
	fmt.Println("Cargando tutores...")
	fmt.Println("Se cargaron los tutores correctamente.")
}

func CargaEstudiantes() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Ingrese el nombre del archivo para la carga de estudiantes: ")
	fmt.Scanln(&ruta)
	listaDoble.LeerCSV(ruta)
	fmt.Println("Cargando estudiantes...")
	fmt.Println("Se cargo los estudiantes correctamente.")
}

func CargaCursos() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Ingrese el nombre del archivo para la carga de cursos: ")
	fmt.Scanln(&ruta)
	arbolCursos.LeerJson(ruta)
	fmt.Println("Cargando cursos...")
	fmt.Println("Se cargaron los cursos correctamente.")
}

func ControlEstudiantes() {
	fmt.Print("\033[H\033[2J")
	opcion := 0
	salir := false

	for !salir {
		colaPrioridad.Primero_Cola()
		fmt.Println("[1]. Aceptar")
		fmt.Println("[2]. Rechazar")
		fmt.Println("[3]. Salir")
		fmt.Scanln(&opcion)
		fmt.Print("\033[H\033[2J")

		switch opcion {
		case 1:
			curso := colaPrioridad.Primero.Tutor.Curso
			comprobar := listaDobleCircular.Buscar(curso)

			if comprobar {
				tutorExistente := listaDobleCircular.BuscarTutor(curso).Tutor
				nuevoTutor := colaPrioridad.Primero.Tutor

				if nuevoTutor.Nota > tutorExistente.Nota {
					fmt.Println("═ Se sustituyó el tutor: ", tutorExistente.Nombre, "del curso: ", curso, "- por el tutor: ", nuevoTutor.Nombre, "═")
					listaDobleCircular.Eliminar(tutorExistente.Carnet)
					listaDobleCircular.Agregar(nuevoTutor.Carnet, nuevoTutor.Nombre, nuevoTutor.Curso, nuevoTutor.Nota)
				} else {
					fmt.Println("═El tutor tiene una nota menor al tutor actual. No se realizará la acción a menos que lo rechace.═")
					break
				}
			} else {
				fmt.Println("═Se registró el tutor: ", colaPrioridad.Primero.Tutor.Nombre, "en el curso:", colaPrioridad.Primero.Tutor.Curso, "con éxito.═")
				listaDobleCircular.Agregar(colaPrioridad.Primero.Tutor.Carnet, colaPrioridad.Primero.Tutor.Nombre, colaPrioridad.Primero.Tutor.Curso, colaPrioridad.Primero.Tutor.Nota)
			}
			colaPrioridad.Descolar()
		case 2:
			fmt.Println("═Tutor rechazado.═")
			colaPrioridad.Descolar()
		case 3:
			salir = true
		default:
			fmt.Println("═Opcion invalida═")
		}
	}
}
