# MANUAL DE USUARIO
## Luis Carlos Corleto Marroquín
### Proyecto único - Tutorías ECYS
#### Menú sistema de inicio
Lo primero en mostrarse al usuario es un menú como el que se muestra a continuación.
![Imágen 1](Images/Imágen%201.jpeg)
En donde él decide si inciar sesión o simplemente salirse del programa.
#### Menú inicio de sesión administrador
Luego, tanto el administrador como el usuario pueden iniciar sesión, en este caso, las credenciales para el usuario administrador será, la palabra ADMIN_#Carnet y su contraseña "admin". Si el usuario ingresa correctamente sus credenciales entrará al menú del administrador.
![Imágen 2](Images/Imágen%202.jpeg)
#### Menú de administrador
![Imágen 3](Images/Imágen%203.jpeg)
El administrador cuenta con 6 opciones dentro de su menú.
* ###### Carga masiva (Estudiantes tutores, estudiantes y cursos):
  Al usuario administrador se le mostrará un mensaje como el siguiente:
  ![Imágen 4](Images/Imágen%204.jpeg)
  Aquí deberá ingresar el nombre del archivo csv para los tutores y estudiantes y un archivo json para los cursos dentro de su directorio, si la carga es exitosa se mostrará el mensaje de éxito y volverá al menú del administrador
  ![Imágen 5](Images/Imágen%205.jpeg)
* ###### Control de estudiantes:
  El administrador tiene la tarea de gestionar (Aceptar o rechazar) a los tutores que quieran inscribirse para dar las tutorías, cada estudiante presenta un curso al que quiere dar dicha tutoría, su nota, nombre y carnet.
  ![Imágen 6](Images/Imágen%206.jpeg)
  El administrador verá estos estudiantes en orden de prioridad (número de nota) de la 1 a la 4, donde la prioridad es la más alta.
  Si el administrador decide aceptar a un tutor se mostrará un mensaje como el siguiente:
  ![Imágen 7](Images/Imágen%207.jpeg)
  Y volverá a mostrar al siguiente tutor en la cola.
  En este punto existe una prioridad adicional, el sistema se encarga de proporcionar la mejor calidad de clase para los cursos, dicho esto, se encarga de verificar que si 2 o más tutores están aplicando al mismo curso priorizar al que tenga la nota más alta, si algún tutor con una nota menor al tutor que ya está ocupando el curso se presenta, se mostrará un mensaje como el siguiente:
  ![Imágen 8](Images/Imágen%208.jpeg)
  En este caso el administrador deberá rechazarlo y no podrá aplicar al curso, si el curso no tiene ningún tutor puede ser aceptado cualquiera con la nota que sea, siempre y cuando no se presente uno con mejor nota, de lo contrario será reemplazado como en el siguiente caso:
  ![Imágen 10](Images/Imágen%2010.jpeg)
  Aquí se muestra un mensaje comunicándole al administrador que el tutor que estaba anteriormente ocupando el curso fue sustituido por uno con mejor nota, indicando el nombre del tutor sustituido, el que sustituye y el curso.
* ##### Menú de reportes:
  El administrador cuenta con un menú de reportes como el siguiente:
  ![Imágen 9](Images/Imágen%209.jpeg)
  Cada reporte muestra como se desarrolló la estructura de datos, para determinada tarea, como indican las opciones, Tutores, Estudiantes, Cursos y Asignaciones.
#### Menú de sesión usuario
Una vez el administrador haya realizado la carga masiva de los estudiantes, cursos y tutores, el usuario podrá iniciar sesión, utilizando como usuario y contraseña su número de carnet, como se muestra en la siguiente imágen:
![Imágen 11](Images/Imágen%2011.jpeg)
Si las credenciales del usuario son correctas, pasará el menú del usuario:
![Imágen 12](Images/Imágen%2012.jpeg)
Aquí se le da la bienvenida al usuario loggeado y se le presentan 3 opciones:
* ##### Tutores disponibles:
  El usuario puede verificar los tutores que fueron cargados y aceptados por el administrador, como se muestra en la siguiente imágen:
  ![Imágen 13](Images/Imágen%2013.jpeg)
  Aquí el usuario visualiza los tutores, con su nombre y curso.
  Es importante aclarar, que estos tutores fueron aceptados, pero puede que en los cursos cargados, no se encuentre alguno por los que aplicaron algunos tutores.
* ##### Asignarse tutores:
  Aquí el usuario al seleccionar la opción, se muestra un mensaje en donde debe ingresar el código del curso para (Asignarse, encontrar el curso sin tutor o no encontrar el curso).
  Si el usuario encuentra el curso con un tutor disponible, se mostrará el siguiente mensaje:
  ![Imágen 14](Images/Imágen%2014.jpeg)
  Si el usuario no encuentra algún curso:
  ![Imágen 15](Images/Imágen%2015.jpeg)
  Y si el usuario encontró el curso sin tutor disponible:
  ![Imágen 16](Images/Imágen%2016.jpeg)
Al finalizar el proceso de asignaciones, el administrador puede proceder a realizar los reportes en su respectivo menú, para ver como quedaron los cursos asignados, los cursos, tutores y estudiantes.
