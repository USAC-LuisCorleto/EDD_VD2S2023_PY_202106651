package ColaPrioridad

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

func (c *Cola) Encolar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutores{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if nota >= 85 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 70 && nota <= 84 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 61 && nota <= 69 {
		nuevoNodo.Prioridad = 3
	} else {
		return
	}

	if c.Longitud == 0 {
		c.Primero = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Primero
		for aux.Siguiente != nil {
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad == nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad < nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else {
				aux = aux.Siguiente
			}
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay tutores en la cola")
	} else {
		c.Primero = c.Primero.Siguiente
		c.Longitud--
	}
}

func (c *Cola) LeerCSV(ruta string) {
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
			fmt.Println("No se pudo leer la linea: ", "[", linea, "]")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}

		carnet, _ := strconv.Atoi(linea[0])
		nota, _ := strconv.Atoi(linea[3])
		c.Encolar(carnet, linea[1], linea[2], nota)
	}
}

func (c *Cola) Primero_Cola() {
	if c.Longitud == 0 {
		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║         No hay más Tutores en la Cola      ║")
		fmt.Println("╚════════════════════════════════════════════╝")
	} else {
		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║            Información del Tutor           ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║ Actual: ", c.Primero.Tutor.Carnet)
		fmt.Println("║ Nombre: ", c.Primero.Tutor.Nombre)
		fmt.Println("║ Curso:  ", c.Primero.Tutor.Curso)
		fmt.Println("║ Nota:   ", c.Primero.Tutor.Nota)
		fmt.Println("║ Prioridad: ", c.Primero.Prioridad)
		if c.Primero.Siguiente != nil {
			fmt.Println("║ Siguiente: ", c.Primero.Siguiente.Tutor.Carnet)
		} else {
			fmt.Println("║ Siguiente: No hay más tutores por evaluar")
		}
		fmt.Println("╚════════════════════════════════════════════╝")
	}
}
