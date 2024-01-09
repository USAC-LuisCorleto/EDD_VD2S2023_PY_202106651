# MANUAL TÉCNICO
## Luis Carlos Corleto Marroquín
#### Proyecto único - Tutorías ECYS
#### OBJETIVOS
**General**
* Aplicar los conocimientos del curso de Estructuras de Datos en el desarrollo de   las diferentes estructuras de datos y los diferentes algoritmos de manipulación de la información de ellas.

**Específicos**
* Utilizar el lenguaje de Go para implementar estructuras de datos lineales.
* Utilizar la herramienta de Graphviz para graficar las estructuras de datos.
* Definir e implementar algoritmos de ordenamiento, búsqueda e inserción en las diferentes estructuras a implementar.

### Definición del problema
La escuela de ciencias y sistemas de la facultad de ingeniería de la universidad de
San Carlos de Guatemala ha visto a muchos estudiantes que han sido Tutores
académicos han demostrado una calidad de estudiantes, dando confianza a los
estudiantes a participar en clases, muchos de ellos son evaluados mediantes
estudiantes tomando en cuenta, su presentación, su formalidad, sus conocimientos,
sus habilidades en los diferentes en los cuales ha sido tutor, por lo cual la escuela
tomó la decisión de darles un incentivo con la ayuda de la universidad, para que
puedan ofrecer tutorías a estudiantes que estén en una etapa inicial, intermedia o
final de su carrera. Por lo cual desean que usted como estudiante de estructuras de
datos, cree una aplicación para llevar control de todos los estudiantes tutores como
los que desean llevar las tutorías.

### Resumen del problema
La escuela de ciencias y sistemas quedo contento con el prototipo de la aplicación en
consola, sin embargo noto algunos inconvenientes con eficiencia, por lo que solicita que
usted, como estudiante del curso de Estructura de datos, se le pide que en base a sus
conocimientos sobre eficiencia haga cambio de algunas estructuras para que pueda
mejorarse el uso de la plataforma aparte de crear una interfaz gráfica interactiva que pueda
ser usado para el estudiante y tutores. Por lo que se le da algunas opciones con Go.

### Estructuras implentadas
#### Arbol B

```
package arbolB

import (
	"EDD_VD2S2023_PY_202106651/estructuras/Peticiones"
	"fmt"
	"strconv"
)

type ArbolB struct {
	Raiz  *RamaB
	Orden int
}

func (a *ArbolB) insertar_rama(nodo *NodoB, rama *RamaB) *NodoB { // 20,
	if rama.Hoja {
		rama.Insertar(nodo)
		if rama.Contador == a.Orden {
			return a.dividir(rama)
		} else {
			return nil
		}
	} else {
		temp := rama.Primero
		for temp != nil {
			if nodo.Valor.Curso == temp.Valor.Curso {
				return nil
			} else if nodo.Valor.Curso < temp.Valor.Curso {
				obj := a.insertar_rama(nodo, temp.Izquierdo)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			} else if temp.Siguiente == nil {
				obj := a.insertar_rama(nodo, temp.Derecho)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			}
			temp = temp.Siguiente
		}
	}
	return nil
}

func (a *ArbolB) dividir(rama *RamaB) *NodoB {
	tutor := &Tutores{Carnet: 0, Nombre: "", Curso: "", Password: ""}
	val := &NodoB{Valor: tutor}
	aux := rama.Primero
	rderecha := &RamaB{Primero: nil, Contador: 0, Hoja: true}
	rizquierda := &RamaB{Primero: nil, Contador: 0, Hoja: true}
	contador := 0
	for aux != nil {
		contador++
		if contador < 2 {
			temp := &NodoB{Valor: aux.Valor}
			temp.Izquierdo = aux.Izquierdo
			if contador == 1 {
				temp.Derecho = aux.Siguiente.Izquierdo
			}
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rizquierda.Hoja = false
			}
			rizquierda.Insertar(temp)
		} else if contador == 2 {
			val.Valor = aux.Valor
		} else {
			temp := &NodoB{Valor: aux.Valor}
			temp.Izquierdo = aux.Izquierdo
			temp.Derecho = aux.Derecho
			if temp.Derecho != nil && temp.Izquierdo != nil {
				rderecha.Hoja = false
			}
			rderecha.Insertar(temp)
		}
		aux = aux.Siguiente
	}
	nuevo := &NodoB{Valor: val.Valor}
	nuevo.Derecho = rderecha
	nuevo.Izquierdo = rizquierda
	return nuevo
}

func (a *ArbolB) Insertar(carnet int, nombre string, curso string, password string) { //15
	tutor := &Tutores{Carnet: carnet, Nombre: nombre, Curso: curso, Password: password}
	nuevoNodo := &NodoB{Valor: tutor}
	if a.Raiz == nil {
		a.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
		a.Raiz.Insertar(nuevoNodo)
	} else {
		obj := a.insertar_rama(nuevoNodo, a.Raiz)
		if obj != nil {
			a.Raiz = &RamaB{Primero: nil, Hoja: true, Contador: 0}
			a.Raiz.Insertar(obj)
			a.Raiz.Hoja = false
		}
	}
}

/***************************************/
func (a *ArbolB) Graficar(nombre string) {
	cadena := ""
	nombre_archivo := "./Reporte/" + nombre + ".dot"
	nombre_imagen := "./Reporte/" + nombre + ".jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol { \nnode[shape=record]\n"
		cadena += a.grafo(a.Raiz.Primero)
		cadena += a.conexionRamas(a.Raiz.Primero)
		cadena += "}"
	}
	Peticiones.CrearArchivo(nombre_archivo)
	Peticiones.EscribirArchivo(cadena, nombre_archivo)
	Peticiones.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolB) grafo(rama *NodoB) string {
	dot := ""
	if rama != nil {
		dot += a.grafoRamas(rama)
		aux := rama
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += a.grafo(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += a.grafo(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

func (a *ArbolB) grafoRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		dot = dot + "R" + rama.Valor.Curso + "[label=\""
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot = dot + "<C" + strconv.Itoa(r) + ">|"
				r++
			}
			if aux.Siguiente != nil {
				dot = dot + aux.Valor.Curso + "|"
			} else {
				dot = dot + aux.Valor.Curso
				if aux.Derecho != nil {
					dot = dot + "|<C" + strconv.Itoa(r) + ">"
				}
			}
			aux = aux.Siguiente
		}
		dot = dot + "\"];\n"
	}
	return dot
}

func (a *ArbolB) conexionRamas(rama *NodoB) string {
	dot := ""
	if rama != nil {
		aux := rama
		actual := "R" + rama.Valor.Curso
		r := 1
		for aux != nil {
			if aux.Izquierdo != nil {
				dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + aux.Izquierdo.Primero.Valor.Curso + ";\n"
				r++
				dot += a.conexionRamas(aux.Izquierdo.Primero)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + aux.Derecho.Primero.Valor.Curso + ";\n"
					r++
					dot += a.conexionRamas(aux.Derecho.Primero)
				}
			}
			aux = aux.Siguiente
		}
	}
	return dot
}

/********************/
func (a *ArbolB) Buscar(numero string, listaSimple *ListaSimple) {
	valTemp, _ := strconv.Atoi(numero)
	a.buscarArbol(a.Raiz.Primero, valTemp, listaSimple)
	if listaSimple.Longitud > 0 {
		fmt.Println("Se encontro el elemento", listaSimple.Longitud)
	} else {
		fmt.Println("No se encontro")
	}
}

func (a *ArbolB) buscarArbol(raiz *NodoB, numero int, listaSimple *ListaSimple) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.buscarArbol(aux.Izquierdo.Primero, numero, listaSimple)
			}
			if aux.Valor.Carnet == numero {
				listaSimple.Insertar(aux)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.buscarArbol(aux.Derecho.Primero, numero, listaSimple)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func (a *ArbolB) GuardarLibro(raiz *NodoB, nombre string, contenido string, carnet int) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.GuardarLibro(aux.Izquierdo.Primero, nombre, contenido, carnet)
			}
			if aux.Valor.Carnet == carnet {
				raiz.Valor.Libros = append(raiz.Valor.Libros, &Libro{Nombre: nombre, Contenido: contenido, Estado: 1, Curso: raiz.Valor.Curso, Tutor: raiz.Valor.Carnet})
				fmt.Println("Registre el libro")
				return
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.GuardarLibro(aux.Derecho.Primero, nombre, contenido, carnet)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func (a *ArbolB) GuardarPublicacion(raiz *NodoB, contenido string, carnet int) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.GuardarPublicacion(aux.Izquierdo.Primero, contenido, carnet)
			}
			if aux.Valor.Carnet == carnet {
				raiz.Valor.Publicaciones = append(raiz.Valor.Publicaciones, &Publicacion{Contenido: contenido, Curso: raiz.Valor.Curso})
				fmt.Println("Registre publicacion")
				return
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.GuardarPublicacion(aux.Derecho.Primero, contenido, carnet)
				}
			}
			aux = aux.Siguiente
		}
	}
}

/*
Visitar Tabla hash, si coincide el alumnos, jalan el atributo Cursos
Buscan en Arbol B, los cursos
*/

/********* NUEVO */

func (a *ArbolB) VerLibroAdmin(raiz *NodoB, listaSimple *ListaSimple) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.VerLibroAdmin(aux.Izquierdo.Primero, listaSimple)
			}
			if len(aux.Valor.Libros) > 0 {
				listaSimple.Insertar(aux)
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.VerLibroAdmin(aux.Derecho.Primero, listaSimple)
				}
			}
			aux = aux.Siguiente
		}
	}
}

func (a *ArbolB) ActualizarLibro(raiz *NodoB, nombre string, curso string, tipo int) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Izquierdo != nil {
				a.ActualizarLibro(aux.Izquierdo.Primero, nombre, curso, tipo)
			}
			if aux.Valor.Curso == curso {
				for i := 0; i < len(aux.Valor.Libros); i++ {
					if aux.Valor.Libros[i].Nombre == nombre {
						aux.Valor.Libros[i].Estado = tipo
					}
				}
				fmt.Println("Actualizado libro")
				return
			}
			if aux.Siguiente == nil {
				if aux.Derecho != nil {
					a.ActualizarLibro(aux.Derecho.Primero, nombre, curso, tipo)
				}
			}
			aux = aux.Siguiente
		}
	}
}
```
**Estructuras de Datos:**

ArbolB:
Es la estructura principal que representa el árbol B.
Contiene un puntero a la raíz del árbol (Raiz) y el orden del árbol (Orden).

NodoB:
Representa un nodo del árbol B que contiene valores y referencias a otros nodos.
Tiene campos para izquierda, derecha y siguiente, así como un puntero a un valor (Tutores).

RamaB:
Representa una rama del árbol B que contiene referencias a nodos.
Tiene un puntero al primer nodo (Primero), un indicador si es una hoja (Hoja), y un contador de nodos (Contador).

Tutores:
Estructura que contiene información sobre los tutores, como el carné, nombre, curso y contraseña.

**Métodos:**

* insertar_rama:
Método para insertar un nodo en una rama del árbol B.
Divide la rama si alcanza el límite definido por el orden del árbol.

* dividir:
Método para dividir una rama cuando alcanza su límite.
Crea dos nuevas ramas y redistribuye los nodos.

* Insertar:
Método para insertar un nuevo tutor en el árbol B.
Llama a insertar_rama y ajusta la raíz si es necesario.

* Graficar:
Método para generar un archivo DOT y una imagen JPG que representa gráficamente el árbol B.

```
package arbolMerkle

import (
	"EDD_VD2S2023_PY_202106651/estructuras/Peticiones"
	"encoding/hex"
	"math"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"
)

type ArbolMerkle struct {
	RaizMerkle      *NodoMerkle
	BloqueDeDatos   *NodoBloqueDatos
	CantidadBloques int
}

func fechaActual() string {
	now := time.Now()
	formato := "02-01-2006::15:04:05"
	fechahoraFormato := now.Format(formato) // 27-12-2023::12:02:40
	return fechahoraFormato
}

func (a *ArbolMerkle) AgregarBloque(estado string, nombreLibro string, carnet int) {
	nuevoRegistro := &InformacionBloque{Fecha: fechaActual(), Accion: estado, Nombre: nombreLibro, Tutor: carnet}
	nuevoBloque := &NodoBloqueDatos{Valor: nuevoRegistro}
	if a.BloqueDeDatos == nil {
		a.BloqueDeDatos = nuevoBloque
		a.CantidadBloques++
	} else {
		aux := a.BloqueDeDatos
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoBloque.Anterior = aux
		aux.Siguiente = nuevoBloque
		a.CantidadBloques++
	}
}

func (a *ArbolMerkle) GenerarArbol() {
	nivel := 1
	for int(math.Pow(2, float64(nivel))) < a.CantidadBloques {
		nivel++
	}
	for i := a.CantidadBloques; i < int(math.Pow(2, float64(nivel))); i++ {
		a.AgregarBloque(strconv.Itoa(i), "nulo", 0)
	}
	/*
		♫ -> ☼ -> ☼ -> ☼ -> ☼ -> nulo -> nulo -> nulo
	*/
	a.generarHash()
}

func (a *ArbolMerkle) generarHash() {
	var arrayNodos []*NodoMerkle
	aux := a.BloqueDeDatos
	for aux != nil {
		contanetacion := aux.Valor.Fecha + aux.Valor.Accion + aux.Valor.Nombre + strconv.Itoa(aux.Valor.Tutor)
		encriptado := a.encriptarSha3(contanetacion)
		nodoHoja := &NodoMerkle{Valor: encriptado, Bloque: aux}
		arrayNodos = append(arrayNodos, nodoHoja)
		aux = aux.Siguiente
	}
	a.RaizMerkle = a.crearArbol(arrayNodos)
}

func (a *ArbolMerkle) crearArbol(arrayNodos []*NodoMerkle) *NodoMerkle {
	var auxNodos []*NodoMerkle
	var raiz *NodoMerkle
	if len(arrayNodos) == 2 {
		encriptado := a.encriptarSha3(arrayNodos[0].Valor + arrayNodos[1].Valor)
		raiz = &NodoMerkle{Valor: encriptado}
		raiz.Izquierda = arrayNodos[0]
		raiz.Derecha = arrayNodos[1]
		return raiz
	} else {
		for i := 0; i < len(arrayNodos); i += 2 {
			encriptado := a.encriptarSha3(arrayNodos[i].Valor + arrayNodos[i+1].Valor)
			nodoRaiz := &NodoMerkle{Valor: encriptado}
			nodoRaiz.Izquierda = arrayNodos[i]
			nodoRaiz.Derecha = arrayNodos[i+1]
			auxNodos = append(auxNodos, nodoRaiz)
		}
		return a.crearArbol(auxNodos)
	}
}

func (a *ArbolMerkle) encriptarSha3(cadena string) string {
	hash := sha3.New256()
	hash.Write([]byte(cadena))
	encriptacion := hex.EncodeToString(hash.Sum(nil))
	return encriptacion
}

func (a *ArbolMerkle) Graficar() {
	cadena := ""
	nombre_archivo := "./Reporte/arbolMerkle.dot"
	nombre_imagen := "./Reporte/arbolMerkle.jpg"
	if a.RaizMerkle != nil {
		cadena += "digraph arbol { node [shape=box];"
		cadena += a.retornarValoresArbol(a.RaizMerkle, 0)
		cadena += "}"
	}
	Peticiones.CrearArchivo(nombre_archivo)
	Peticiones.EscribirArchivo(cadena, nombre_archivo)
	Peticiones.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolMerkle) retornarValoresArbol(raiz *NodoMerkle, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" [dir=back];\n"
		if raiz.Izquierda != nil && raiz.Derecha != nil {
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierda.Valor[:20]) + "\"" + " -> " + "\"" + (raiz.Derecha.Valor[:20]) + "\"" + " [style=invis]}; \n"
		}
	}
	if raiz.Bloque != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" -> "
		cadena += "\""
		cadena += raiz.Bloque.Valor.Fecha + "\n" + raiz.Bloque.Valor.Accion + "\n" + raiz.Bloque.Valor.Nombre + "\n" + strconv.Itoa(raiz.Bloque.Valor.Tutor)
		cadena += "\" [dir=back];\n "
	}
	return cadena
}
```
**Estructuras de Datos:**

ArbolMerkle:
Representa el árbol Merkle y contiene un puntero a la raíz del árbol (RaizMerkle), un puntero al bloque de datos (BloqueDeDatos), y la cantidad de bloques en el árbol (CantidadBloques).

NodoBloqueDatos:
Representa un nodo en la lista de bloques de datos enlazados.
Contiene un puntero al siguiente bloque (Siguiente), un puntero al bloque anterior (Anterior), y un puntero al valor (Valor) que contiene información sobre el bloque.

InformacionBloque:
Estructura que almacena información sobre un bloque, como la fecha, la acción, el nombre del libro y el carnet del tutor.

NodoMerkle:
Representa un nodo en el árbol Merkle.
Contiene un valor hash (Valor), punteros a nodos hijos izquierdo y derecho (Izquierda y Derecha), y un puntero al bloque de datos asociado (Bloque).

**Métodos:**

* fechaActual:
Función auxiliar que devuelve la fecha y hora actual en un formato específico.

* AgregarBloque:
Método para agregar un bloque de datos al árbol Merkle.
Crea un nuevo nodo de bloque con la información proporcionada y lo agrega al final de la lista.

* GenerarArbol:
Método para generar el árbol Merkle.
Completa la cantidad de bloques hasta la potencia de 2 más cercana, y luego llama a generarHash.

* generarHash:
Método para generar el hash de cada bloque de datos y construir el árbol Merkle.

* crearArbol:
Método recursivo para construir el árbol Merkle a partir de los nodos hoja.

* encriptarSha3:
Función para realizar la encriptación SHA3 de una cadena.

* Graficar:
Método para generar un archivo DOT y una imagen JPG que representa gráficamente el árbol Merkle.

* retornarValoresArbol:
Método recursivo para obtener los valores del árbol Merkle para su representación gráfica en formato DOT.
```
package grafo

import "EDD_VD2S2023_PY_202106651/estructuras/Peticiones"

type Grafo struct {
	Principal *NodoListaAdyacencia
}

func (g *Grafo) insertarColumna(curso string, post string) {
	nuevoNodo := &NodoListaAdyacencia{Valor: post}
	if g.Principal != nil && curso == g.Principal.Valor {
		g.insertarFila(post)
		aux := g.Principal
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
	} else {
		g.insertarFila(curso)
		aux := g.Principal
		for aux != nil {
			if aux.Valor == curso {
				break
			}
			aux = aux.Abajo
		}
		if aux != nil {
			for aux.Siguiente != nil {
				aux = aux.Siguiente
			}
			aux.Siguiente = nuevoNodo
		}
	}
}

func (g *Grafo) insertarFila(curso string) {
	nuevoNodo := &NodoListaAdyacencia{Valor: curso}
	if g.Principal == nil {
		g.Principal = nuevoNodo
	} else {
		aux := g.Principal
		for aux.Abajo != nil {
			if aux.Valor == curso {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nuevoNodo
	}
}

func (g *Grafo) InsertarValores(curso string, post string) {
	if g.Principal == nil {
		//insertar Fila
		g.insertarFila(curso)
		//insertar Columna
		g.insertarColumna(curso, post)
	} else {
		g.insertarColumna(curso, post)
	}
}

func (g *Grafo) Reporte(nombre string) {
	cadena := ""
	nombre_archivo := "./" + nombre + ".dot"
	nombre_imagen := nombre + ".jpg"
	if g.Principal != nil {
		cadena += "digraph grafoDirigido{ \n rankdir=LR; \n node [shape=box]; layout=neato; \n nodo" + g.Principal.Valor + "[label=\"" + g.Principal.Valor + "\"]; \n"
		cadena += "node [shape = ellipse]; \n"
		cadena += g.retornarValoresMatriz()
		cadena += "\n}"
	}
	Peticiones.CrearArchivo(nombre_archivo)
	Peticiones.EscribirArchivo(cadena, nombre_archivo)
	Peticiones.Ejecutar(nombre_imagen, nombre_archivo)
}

func (g *Grafo) retornarValoresMatriz() string {
	cadena := ""
	/*CREACION DE NODOS*/
	aux := g.Principal.Abajo //Filas
	aux1 := aux              //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux1.Valor + "[label=\"" + aux1.Valor + "\" ]; \n"
			aux1 = aux1.Siguiente
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
	}
	/*CONEXION DE NODOS*/
	aux = g.Principal    //Filas
	aux1 = aux.Siguiente //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux.Valor + " -> "
			cadena += "nodo" + aux1.Valor + "[len=1.00]; \n"
			aux1 = aux1.Siguiente
		}
		if aux.Abajo != nil {
			aux = aux.Abajo
			aux1 = aux.Siguiente
		} else {
			aux = aux.Abajo
		}
	}

	return cadena
}
```
**Estructuras de Datos:**

Grafo:
Representa el grafo dirigido y contiene un puntero al nodo principal (Principal), que representa la primera fila de la matriz de adyacencia.

NodoListaAdyacencia:
Representa un nodo en la lista de adyacencia de un grafo.
Contiene un valor (Valor), que puede ser el nombre del curso o el post, y punteros a nodos adyacentes (Abajo y Siguiente).

**Métodos:**

* insertarColumna:
Método para insertar una columna en la matriz de adyacencia.
Añade un nuevo nodo a la fila correspondiente.

* insertarFila:
Método para insertar una fila en la matriz de adyacencia.
Añade un nuevo nodo principal si no existe, o agrega un nuevo nodo a la última fila existente.

* InsertarValores:
Método para insertar valores en la matriz de adyacencia.
Llama a los métodos insertarFila y insertarColumna según sea necesario.

* Reporte:
Método para generar un archivo DOT y una imagen JPG que representan gráficamente el grafo dirigido.

* retornarValoresMatriz:
Método para obtener los valores de la matriz de adyacencia para su representación gráfica en formato DOT.
```
package tablaHash

import "strconv"

type TablaHash struct {
	Tabla       map[int]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) calculoIndice(carnet int) int {
	var numeros []int
	for {
		if carnet > 0 {
			digito := carnet % 10
			numeros = append([]int{digito}, numeros...)
			carnet = carnet / 10
		} else {
			break
		}
	}

	var numeros_ascii []rune
	for _, numero := range numeros {
		valor := rune(numero + 48)
		numeros_ascii = append(numeros_ascii, valor)
	}

	final := 0
	for _, numero_ascii := range numeros_ascii {
		final += int(numero_ascii)
	}

	indice := final % t.Capacidad
	return indice
}

func (t *TablaHash) capacidadTabla() {
	auxCap := float64(t.Capacidad) * 0.6
	if t.Utilizacion > int(auxCap) {
		auxAnterior := t.Capacidad
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar(auxAnterior)
	}
}
func (t *TablaHash) nuevaCapacidad() int {
	contador := 0
	a, b := 0, 1
	for contador < 100 {
		contador++
		if a > t.Capacidad {
			return a
		}
		a, b = b, a+b
	}
	return a
}

func (t *TablaHash) reInsertar(capacidadAnterior int) {
	auxTabla := t.Tabla
	t.Tabla = make(map[int]NodoHash)
	for i := 0; i < capacidadAnterior; i++ {
		if usuario, existe := auxTabla[i]; existe {
			t.Insertar(usuario.Persona.Carnet, usuario.Persona.Nombre, usuario.Persona.Password, usuario.Persona.Cursos)
		}
	}
}

func (t *TablaHash) reCalculoIndice(carnet int, contador int) int {
	nuevoIndice := t.calculoIndice(carnet) + (contador * contador) //5+4=9
	return t.nuevoIndice(nuevoIndice)
}

func (t *TablaHash) nuevoIndice(nuevoIndice int) int {
	nuevoPosicion := 0
	if nuevoIndice < t.Capacidad {
		nuevoPosicion = nuevoIndice
	} else {
		nuevoPosicion = nuevoIndice - t.Capacidad
		nuevoPosicion = t.nuevoIndice(nuevoPosicion)
	}
	return nuevoPosicion
}

func (t *TablaHash) Insertar(carnet int, nombre string, password string, cursos []string) { // cursos []string
	indice := t.calculoIndice(carnet)
	nuevoNodo := &NodoHash{Llave: indice, Persona: &Persona{Carnet: carnet, Nombre: nombre, Password: password, Cursos: cursos}}
	if indice < t.Capacidad {
		if _, existe := t.Tabla[indice]; !existe {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidadTabla()
		} else {
			contador := 1
			indice = t.reCalculoIndice(carnet, contador)
			for {
				if _, existe := t.Tabla[indice]; existe {
					contador++
					indice = t.reCalculoIndice(carnet, contador)
				} else {
					nuevoNodo.Llave = indice
					t.Tabla[indice] = *nuevoNodo
					t.Utilizacion++
					t.capacidadTabla()
					break
				}
			}
		}
	}
}

func (t *TablaHash) Buscar(carnet string, password string) bool {
	valTemp, err := strconv.Atoi(carnet)
	if err != nil {
		return false
	}
	indice := t.calculoIndice(valTemp)
	if indice < t.Capacidad {
		if usuario, existe := t.Tabla[indice]; existe {
			if usuario.Persona.Carnet == valTemp {
				if usuario.Persona.Password == password {
					return true
				}
			} else {
				contador := 1
				indice = t.reCalculoIndice(valTemp, contador)
				for {
					if usuario, existe := t.Tabla[indice]; existe {
						if usuario.Persona.Carnet == valTemp {
							if usuario.Persona.Password == password {
								return true
							} else {
								return false
							}
						} else {
							contador++
							indice = t.reCalculoIndice(valTemp, contador)
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}

func (t *TablaHash) ConvertirArreglo() []NodoHash {
	var arrays []NodoHash
	if t.Utilizacion > 0 {
		for i := 0; i < t.Capacidad; i++ {
			if usuario, existe := t.Tabla[i]; existe {
				arrays = append(arrays, usuario)
			}
		}
	}
	return arrays
}

func (t *TablaHash) BuscarSesion(carnet string) *Persona {
	valTemp, err := strconv.Atoi(carnet)
	if err != nil {
		return nil
	}
	indice := t.calculoIndice(valTemp)
	if indice < t.Capacidad {
		if usuario, existe := t.Tabla[indice]; existe {
			if usuario.Persona.Carnet == valTemp {
				return usuario.Persona
			} else {
				contador := 1
				indice = t.reCalculoIndice(valTemp, contador)
				for {
					if usuario, existe := t.Tabla[indice]; existe {
						if usuario.Persona.Carnet == valTemp {
							return usuario.Persona
						} else {
							contador++
							indice = t.reCalculoIndice(valTemp, contador)
						}
					} else {
						return nil
					}
				}
			}
		}
	}
	return nil
}
```
**Estructuras de Datos:**

TablaHash:
Representa la tabla hash cerrada.
Contiene un mapa (Tabla) que almacena nodos hash.
Capacidad representa la capacidad máxima de la tabla.
Utilizacion cuenta la cantidad actual de elementos en la tabla.

NodoHash:
Representa un nodo en la tabla hash.
Contiene una clave (Llave) que es el índice en la tabla hash y un puntero a una estructura Persona.

Persona:
Representa la información de una persona con carnet, nombre, contraseña y cursos.

**Métodos Principales:**

* calculoIndice:
Calcula el índice de la tabla hash utilizando el método de suma de los dígitos del carnet.

* capacidadTabla:
Verifica si la utilización de la tabla supera el 60% de su capacidad, y si es así, realiza un redimensionamiento.

* nuevaCapacidad:
Calcula la nueva capacidad de la tabla en base a una serie de Fibonacci.

* reInsertar:
Reinserta todos los elementos en la tabla después de un redimensionamiento.

* reCalculoIndice:
Calcula un nuevo índice utilizando la técnica de cuadrados cuadrados en caso de colisión.

* nuevoIndice:
Ajusta el índice en caso de que sea necesario después de un redimensionamiento.

* Insertar:
Inserta un nuevo elemento en la tabla hash.

* Buscar:
Busca un elemento en la tabla hash por carnet y contraseña.

* ConvertirArreglo:
Convierte la tabla hash en un arreglo de nodos hash.

* BuscarSesion:
Busca un elemento en la tabla hash por carnet y devuelve la información de la persona.