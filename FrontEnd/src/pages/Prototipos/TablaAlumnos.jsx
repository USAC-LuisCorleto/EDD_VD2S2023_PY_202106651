import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";

function TablaAlumnos() {
  const [alumnosRegistrados, SetAlumnosRegistrados] = useState([]);
  
  useEffect(() => {
    async function peticion() {
      const response = await fetch("http://localhost:4000/tabla-alumno");
      const result = await response.json();
      SetAlumnosRegistrados(result.Arreglo);
    }
    peticion();
  }, []);

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h1 mb-3 fw-normal">TABLA DE ALUMNOS</h1>
          <br />
          <div className="table-container">
            <table className="table table-dark table-striped">
              <thead>
                <tr>
                  <th scope="col">Carnet</th>
                  <th scope="col">Nombre</th>
                  <th scope="col">Password </th>
                  <th scope="col">Cursos </th>
                </tr>
              </thead>
              <tbody>
                {alumnosRegistrados.map((element, j) => {
                  if (element.Id_Cliente !== "") {
                    return (
                      <tr key={"alum" + j}>
                        <td>{element.Persona.Carnet}</td>
                        <td>{element.Persona.Nombre}</td>
                        <td>{element.Persona.Password}</td>
                        <td>{element.Persona.Cursos.join(" - ")}</td>
                      </tr>
                    );
                  }
                  return null;
                })}
              </tbody>
            </table>
          </div>
          <br />
          <br />
        </form>
      </div>
    </div>
  );
}

export default TablaAlumnos;