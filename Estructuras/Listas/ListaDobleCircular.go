package Listas

import "fmt"

type ListaDobleCircular struct {
	Inicio   *NodoListaCircular
	Longitud int
}

func (l *ListaDobleCircular) Agregar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutores{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoListaCircular{Tutor: nuevoTutor, Siguiente: nil, Anterior: nil}

	if l.Longitud == 0 {
		l.Inicio = nuevoNodo
		l.Inicio.Anterior = nuevoNodo
		l.Inicio.Siguiente = nuevoNodo
		l.Longitud++
	} else {
		aux := l.Inicio
		contador := 1
		for contador < l.Longitud {
			if l.Inicio.Tutor.Carnet > carnet {
				nuevoNodo.Siguiente = l.Inicio
				nuevoNodo.Anterior = l.Inicio.Anterior
				l.Inicio.Anterior = nuevoNodo
				l.Inicio = nuevoNodo
				l.Longitud++
				return
			}
			if aux.Tutor.Carnet < carnet {
				aux = aux.Siguiente
			} else {
				nuevoNodo.Anterior = aux.Anterior
				aux.Anterior.Siguiente = nuevoNodo
				nuevoNodo.Siguiente = aux
				aux.Anterior = nuevoNodo
				l.Longitud++
				return
			}
			contador++
		}
		if aux.Tutor.Carnet > carnet {
			nuevoNodo.Siguiente = aux
			nuevoNodo.Anterior = aux.Anterior
			aux.Anterior.Siguiente = nuevoNodo
			aux.Anterior = nuevoNodo
			l.Longitud++
			return
		}
		nuevoNodo.Anterior = aux
		nuevoNodo.Siguiente = l.Inicio
		aux.Siguiente = nuevoNodo
		l.Longitud++
	}
}

func (l *ListaDobleCircular) Mostrar() {
	aux := l.Inicio
	contador := 1
	for contador <= l.Longitud {
		fmt.Println("Código del curso:", "[", aux.Tutor.Curso, "]", "-> Nombre del tutor:", "[", aux.Tutor.Nombre, "]")
		aux = aux.Siguiente
		contador++
	}
}

func (l *ListaDobleCircular) Buscar(curso string) bool {
	if l.Longitud == 0 {
		return false
	} else {
		aux := l.Inicio
		contador := 1
		for l.Longitud >= contador {
			if aux.Tutor.Curso == curso {
				return true
			}
			aux = aux.Siguiente
			contador++
		}
	}
	return false
}

func (l *ListaDobleCircular) BuscarTutor(curso string) *NodoListaCircular {
	aux := l.Inicio
	contador := 1
	for l.Longitud >= contador {
		if aux.Tutor.Curso == curso {
			return aux
		}
		aux = aux.Siguiente
		contador++
	}
	return nil
}

func (l *ListaDobleCircular) Eliminar(carnet int) {
	if l.Longitud == 0 {
		fmt.Println("La lista está vacía.")
		return
	}

	aux := l.Inicio
	for i := 1; i <= l.Longitud; i++ {
		if aux.Tutor.Carnet == carnet {
			if l.Longitud == 1 {
				l.Inicio = nil
			} else {
				aux.Anterior.Siguiente = aux.Siguiente
				aux.Siguiente.Anterior = aux.Anterior

				if aux == l.Inicio {
					l.Inicio = aux.Siguiente
				}
			}

			fmt.Println("Tutor sustituido con éxito.")
			l.Longitud--
			return
		}
		aux = aux.Siguiente
	}

	fmt.Println("No se encontró al tutor con carnet:", carnet)
}
