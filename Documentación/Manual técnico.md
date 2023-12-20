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

### Estructuras implentadas
#### Lista doble
```
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
	nombreArchivo := "./ReporteEstudiantes.dot"
	nombreImagen := "./ReporteEstudiantes.jpg"
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
```
**Estructura**:
* ListaDoble: Una estructura que representa la lista doblemente enlazada. Contiene un puntero al inicio de la lista (Inicio) y la longitud de la lista (Longitud).

* NodoListaDoble: Cada nodo de la lista contiene un puntero al siguiente nodo (Siguiente), un puntero al nodo anterior (Anterior), y un objeto de tipo Alumno que almacena información sobre el estudiante.

* Alumno: Estructura que contiene información sobre un estudiante, como el carnet y el nombre.

**Métodos**:
1. Agregar: Agrega un nuevo estudiante a la lista. Si la lista está vacía, el nuevo nodo se convierte en el inicio. En caso contrario, se recorre la lista hasta el último nodo y se agrega el nuevo nodo al final.

2. Buscar: Busca un estudiante en la lista según el carnet y la contraseña proporcionados. Retorna true si se encuentra una coincidencia y false en caso contrario.

3. LeerCSV: Lee un archivo CSV que contiene información de estudiantes y agrega cada estudiante a la lista mediante el método Agregar. Ignora la primera línea del CSV (encabezado).

4. ReporteListaDoble: Genera un archivo DOT (Gráfico o reporte) para representar la lista doblemente enlazada. Utiliza el formato DOT para describir nodos, conexiones y atributos. Después, utiliza un paquete llamado GenerarArchivos para crear un archivo de imagen (JPG) a partir del archivo DOT.

#### Lista doble circular
```
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
	fmt.Print("\033[H\033[2J")
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                               Lista de Tutores Disponibles                             ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════════════╣")

	aux := l.Inicio
	contador := 1

	for contador <= l.Longitud {
		fmt.Println("║ Código del curso: [", aux.Tutor.Curso, "] -> Nombre del tutor: [", aux.Tutor.Nombre, "] ║")
		aux = aux.Siguiente
		contador++
	}

	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════╝")
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

func (l *ListaDoble) ReporteListaDobleCircular() {
	nombreArchivo := "./ReporteTutores.dot"
	nombreImagen := "./ReporteTutores.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Inicio
	contador := 0
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + strconv.Itoa(aux.Alumno.Carnet) + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0 \n"
	texto += "nodo0 -> " + "nodo" + strconv.Itoa(contador) + "\n"
	texto += "}"
	GenerarArchivos.CrearArchivo(nombreArchivo)
	GenerarArchivos.EscribirArchivo(texto, nombreArchivo)
	GenerarArchivos.Ejecutar(nombreImagen, nombreArchivo)
}
```
**Estructura**:
* ListaDobleCircular: Una estructura que representa la lista doblemente circular. Contiene un puntero al inicio de la lista (Inicio) y la longitud de la lista (Longitud).

* NodoListaCircular: Cada nodo de la lista contiene un puntero al siguiente nodo (Siguiente), un puntero al nodo anterior (Anterior), y un objeto de tipo Tutores que almacena información sobre el tutor.

* Tutores: Estructura que contiene información sobre un tutor, como el carnet, nombre, curso y nota.

**Métodos**:
1. Agregar: Agrega un nuevo tutor a la lista en orden ascendente según el carnet. Si la lista está vacía, el nuevo nodo se convierte en el único nodo de la lista.

2. Mostrar: Imprime en la consola la información de todos los tutores en la lista.

3. Buscar: Busca un tutor en la lista según el curso proporcionado. Retorna true si se encuentra un tutor con el curso especificado y false en caso contrario.

4. BuscarTutor: Busca un tutor en la lista según el curso proporcionado y retorna el nodo correspondiente si se encuentra, o nil si no se encuentra.

5. Eliminar: Elimina un tutor de la lista según el carnet proporcionado. Si el tutor es el único en la lista, se establece el inicio como nil. Si hay más tutores, se ajustan los punteros de los nodos adyacentes para eliminar el nodo.

6. ReporteListaDobleCircular: Genera un archivo DOT (Gráfico o reporte) para representar la lista doblemente circular. Similar al método ReporteListaDoble, utiliza el paquete GenerarArchivos para crear un archivo de imagen (JPG) a partir del archivo DOT.

#### Cola de prioridad
```
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
```
**Estructura**:
* Cola: Representa la cola de tutores. Contiene un puntero al primer nodo (Primero) y la longitud de la cola (Longitud).

* NodoCola: Cada nodo de la cola contiene un puntero al siguiente nodo (Siguiente), una prioridad (Prioridad) y un objeto de tipo Tutores que almacena información sobre el tutor.

* Tutores: Estructura que contiene información sobre un tutor, como el carnet, nombre, curso y nota.

**Métodos**:
1. Encolar: Agrega un nuevo tutor a la cola con una prioridad basada en su nota. La prioridad se asigna según rangos de notas predefinidos (Puede visualizarse en el enunciado). La cola se mantiene ordenada por prioridad.

2. Descolar: Elimina el tutor que está al frente de la cola.

3. LeerCSV: Lee un archivo CSV que contiene información de tutores y los encola según el método Encolar. Ignora la primera línea del CSV (encabezado).

4. Primero_Cola: Imprime en la consola la información del tutor que está al frente de la cola, incluyendo su carnet, nombre, curso, nota, prioridad y el carnet del siguiente tutor en la cola.

5. OrdenarPorPrioridad: Ordena la cola de tutores por prioridad.

6. OrdenarPorNota: Ordena la cola de tutores por nota de mayor a menor.

#### Árbol AVL
```
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

func (a *ArbolAVL) rotacionI(raiz *NodoArbol) *NodoArbol {
	raiz_derecho := raiz.Derecho
	hijo_izquierdo := raiz_derecho.Izquierdo
	raiz_derecho.Izquierdo = raiz
	raiz.Derecho = hijo_izquierdo
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	numeroMax = math.Max(float64(a.altura(raiz_derecho.Izquierdo)), float64(a.altura(raiz_derecho.Derecho)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz_derecho.Factor_Equilibrio = a.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (a *ArbolAVL) rotacionD(raiz *NodoArbol) *NodoArbol {
	raiz_izquierdo := raiz.Izquierdo
	hijo_derecho := raiz_izquierdo.Derecho
	raiz_izquierdo.Derecho = raiz
	raiz.Izquierdo = hijo_derecho
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

func (a *ArbolAVL) ReporteArbolAVL() {
	cadena := ""
	nombre_archivo := "./ReporteCursos.dot"
	nombre_imagen := "ReporteCursos.jpg"
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
```
**Estructuras de datos**:
* ArbolAVL: Representa un árbol AVL y contiene un puntero a la raíz del árbol.

* Curso: Estructura que representa un curso con un código y un nombre.

* DatosCursos: Estructura para deserializar datos desde un archivo JSON que contiene una lista de cursos.

**Métodos**:
1. Altura: Devuelve la altura de un nodo en el árbol. Si el nodo es nulo, devuelve 0.

2. Equilibrio: Calcula el factor de equilibrio de un nodo en el árbol AVL.

3. RotacionI (rotación izquierda): Realiza una rotación izquierda en el árbol AVL.

4. RotacionD (rotación derecha): Realiza una rotación derecha en el árbol AVL.

5. InsertarNodo: Inserta un nuevo nodo en el árbol AVL, manteniendo su propiedad AVL mediante rotaciones.

6. InsertarElemento: Inserta un nuevo elemento (curso) en el árbol AVL.

7. Busqueda_arbol: Busca un valor en el árbol AVL y devuelve el nodo si se encuentra.

8. Busqueda: Verifica si un valor dado está presente en el árbol AVL.

9. LeerJson: Lee un archivo JSON que contiene datos de cursos y los inserta en el árbol AVL.

10. ReporteArbolAVL: Genera un archivo DOT para representar visualmente el árbol AVL y lo convierte en una imagen.

11. RetornarValoresArbol: Método auxiliar para generar el contenido del archivo DOT para representar visualmente el árbol AVL.

12. MostrarCursos: Imprime en consola la lista de cursos disponibles en el árbol AVL.

13. MostrarCursosRecursivo: Método auxiliar para mostrar los cursos en orden recursivamente.

