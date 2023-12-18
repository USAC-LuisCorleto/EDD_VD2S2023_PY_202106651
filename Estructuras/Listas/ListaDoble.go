package Listas

import (
	"EDD_VD2S2023_PY_202106651/Estructuras/GenerarArchivos"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type ListaDoble struct {
	Inicio   *NodoListaDoble
	Longitud int
}

func (l *ListaDoble) Agregar(carnet int, nombre string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodoListaDoble{Alumno: nuevoAlumno, Siguiente: nil, Anterior: nil}

	if l.Longitud == 0 {
		l.Inicio = nuevoNodo
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoNodo.Anterior = aux
		aux.Siguiente = nuevoNodo
		l.Longitud++
	}
}

func (l *ListaDoble) Buscar(carnet string, password string) bool {
	if l.Longitud == 0 {
		return false
	} else {
		aux := l.Inicio
		for aux != nil {
			if strconv.Itoa(aux.Alumno.Carnet) == carnet && carnet == password {
				return true
			}
			aux = aux.Siguiente
		}
	}
	return false
}

func (l *ListaDoble) LeerCSV(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo: ", "[", ruta, "]")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No se pudo leer la línea: ", "[", linea, "] del CSV")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		l.Agregar(valor, linea[1])
	}
}

func (l *ListaDoble) ReporteListaDoble() {
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record, style=\"filled\", fillcolor=\"#FFDDC1\", fontname=\"Arial\"];\n"
	texto += "nodonull1[label=\"null\", shape=plaintext];\n"
	texto += "nodonull2[label=\"null\", shape=plaintext];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + strconv.Itoa(aux.Alumno.Carnet) + "\", style=\"filled\", fillcolor=\"#87CEEB\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	GenerarArchivos.CrearArchivo(nombreArchivo)
	GenerarArchivos.EscribirArchivo(texto, nombreArchivo)
	GenerarArchivos.Ejecutar(nombreImagen, nombreArchivo)
}
