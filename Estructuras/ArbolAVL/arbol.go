package ArbolAVL

import (
	"EDD_VD2S2023_PY_202106651/Estructuras/GenerarArchivos"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type ArbolAVL struct {
	Raiz *NodoArbol
}

type Curso struct {
	Codigo string `json:"Codigo"`
	Nombre string `json:"Nombre"`
}

type DatosCursos struct {
	Cursos []Curso `json:"Cursos"`
}

func (a *ArbolAVL) altura(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (a *ArbolAVL) equilibrio(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return (a.altura(raiz.Derecho) - a.altura(raiz.Izquierdo))
}

func (a *ArbolAVL) rotacionI(raiz *NodoArbol) *NodoArbol { //Raiz = 10
	raiz_derecho := raiz.Derecho             // 10.derecho = 15
	hijo_izquierdo := raiz_derecho.Izquierdo // 10.derecho.izquierdo = null
	raiz_derecho.Izquierdo = raiz            // 15.izquierdo = 10
	raiz.Derecho = hijo_izquierdo            // 10.derecho = null
	/*Calcular nuevamente alturas de raiz*/
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	/*Calcular nuevamente alturas de raiz.derecho*/
	numeroMax = math.Max(float64(a.altura(raiz_derecho.Izquierdo)), float64(a.altura(raiz_derecho.Derecho)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz_derecho.Factor_Equilibrio = a.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (a *ArbolAVL) rotacionD(raiz *NodoArbol) *NodoArbol { //Raiz = 20
	raiz_izquierdo := raiz.Izquierdo       // 20.izquierdo = 15
	hijo_derecho := raiz_izquierdo.Derecho // 20.izquierdo.derecho = null
	raiz_izquierdo.Derecho = raiz          //15.derecho = 20
	raiz.Izquierdo = hijo_derecho          // 20.izquierdo = null
	/*Calcular nuevamente alturas de raiz.derecho*/
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	numeroMax = math.Max(float64(a.altura(raiz_izquierdo.Izquierdo)), float64(a.altura(raiz_izquierdo.Derecho)))
	raiz_izquierdo.Altura = 1 + int(numeroMax)
	raiz_izquierdo.Factor_Equilibrio = a.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (a *ArbolAVL) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor > nuevoNodo.Valor {
			raiz.Izquierdo = a.insertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.insertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	balanceo := a.equilibrio(raiz)
	raiz.Factor_Equilibrio = balanceo
	if balanceo > 1 && nuevoNodo.Valor > raiz.Derecho.Valor {
		return a.rotacionI(raiz)
	} else if balanceo < -1 && nuevoNodo.Valor < raiz.Izquierdo.Valor {
		return a.rotacionD(raiz)
	} else if balanceo > 1 && nuevoNodo.Valor < raiz.Derecho.Valor {
		raiz.Derecho = a.rotacionD(raiz.Derecho)
		return a.rotacionI(raiz)
	} else if balanceo < -1 && nuevoNodo.Valor > raiz.Izquierdo.Valor {
		raiz.Izquierdo = a.rotacionI(raiz.Izquierdo)
		return a.rotacionD(raiz)
	}
	return raiz
}

func (a *ArbolAVL) InsertarElemento(valor string) {
	nuevoNodo := &NodoArbol{Valor: valor}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *ArbolAVL) busqueda_arbol(valor string, raiz *NodoArbol) *NodoArbol {
	var valorEncontro *NodoArbol
	if raiz != nil {
		if raiz.Valor == valor {
			valorEncontro = raiz
		} else {
			if raiz.Valor > valor {
				valorEncontro = a.busqueda_arbol(valor, raiz.Izquierdo)
			} else {
				valorEncontro = a.busqueda_arbol(valor, raiz.Derecho)
			}
		}
	}
	return valorEncontro
}

func (a *ArbolAVL) Busqueda(valor string) bool {
	return a.busqueda_arbol(valor, a.Raiz) != nil
}

func (a *ArbolAVL) LeerJson(ruta string) {
	data, err := os.ReadFile(ruta)
	if err != nil {
		log.Fatal("Error al leer el archivo:", err)
	}

	var datos DatosCursos
	err = json.Unmarshal(data, &datos)
	if err != nil {
		log.Fatal("Error al decodificar el JSON:", err)
	}

	for _, curso := range datos.Cursos {
		a.InsertarElemento(curso.Codigo)
	}
}

func (a *ArbolAVL) Graficar() {
	cadena := ""
	nombre_archivo := "./ArbolAVL.dot"
	nombre_imagen := "ArbolAVL.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol{ "
		cadena += a.retornarValoresArbol(a.Raiz, 0)
		cadena += "}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolAVL) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor
		cadena += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "\"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}

func (a *ArbolAVL) MostrarCursos() {
	fmt.Println("Cursos disponibles:")
	a.mostrarCursosRecursivo(a.Raiz)
}

func (a *ArbolAVL) mostrarCursosRecursivo(raiz *NodoArbol) {
	if raiz != nil {
		a.mostrarCursosRecursivo(raiz.Izquierdo)
		fmt.Println(raiz.Valor)
		a.mostrarCursosRecursivo(raiz.Derecho)
	}
}
