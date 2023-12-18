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

	if nota >= 90 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 75 && nota <= 89 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 65 && nota <= 74 {
		nuevoNodo.Prioridad = 3
	} else if nota >= 61 && nota <= 64 {
		nuevoNodo.Prioridad = 4
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

func (c *Cola) OrdenarPorPrioridad() {
	if c.Longitud <= 1 {
		return
	}
	nodosOrdenados := make([]*NodoCola, c.Longitud)
	aux := c.Primero
	for i := 0; i < c.Longitud; i++ {
		nodosOrdenados[i] = aux
		aux = aux.Siguiente
	}
	for i := 0; i < c.Longitud-1; i++ {
		for j := 0; j < c.Longitud-i-1; j++ {
			if nodosOrdenados[j].Prioridad > nodosOrdenados[j+1].Prioridad {
				nodosOrdenados[j], nodosOrdenados[j+1] = nodosOrdenados[j+1], nodosOrdenados[j]
			}
		}
	}
	c.Primero = nodosOrdenados[0]
	aux = c.Primero
	for i := 1; i < c.Longitud; i++ {
		aux.Siguiente = nodosOrdenados[i]
		aux = aux.Siguiente
	}
	aux.Siguiente = nil
}

func (c *Cola) OrdenarPorNota() {
	if c.Longitud <= 1 {
		return
	}
	nodosOrdenados := make([]*NodoCola, c.Longitud)
	aux := c.Primero
	for i := 0; i < c.Longitud; i++ {
		nodosOrdenados[i] = aux
		aux = aux.Siguiente
	}
	for i := 0; i < c.Longitud-1; i++ {
		for j := 0; j < c.Longitud-i-1; j++ {
			if nodosOrdenados[j].Tutor.Nota < nodosOrdenados[j+1].Tutor.Nota {
				nodosOrdenados[j], nodosOrdenados[j+1] = nodosOrdenados[j+1], nodosOrdenados[j]
			}
		}
	}
	c.Primero = nodosOrdenados[0]
	aux = c.Primero
	for i := 1; i < c.Longitud; i++ {
		aux.Siguiente = nodosOrdenados[i]
		aux = aux.Siguiente
	}
	aux.Siguiente = nil
}
