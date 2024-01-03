import React, { useState, useEffect } from "react";

function Publicacion_Estudiante() {
  const [cursos, setCursos] = useState([]);
  const [publicacion, setPublicacion] = useState([]);

  useEffect(() => {
    async function PedirCursos() {
      const arregloCursos = localStorage.getItem("cursos");
      const cursosArreglo = JSON.parse(arregloCursos);
      const response = await fetch(
        "http://localhost:4000/obtener-publi-alumno"
      );
      const result = await response.json();
      console.log(result);
      setCursos(cursosArreglo);
      if (result.status === 200) {
        setPublicacion(result.Arreglo);
      }
    }
    PedirCursos();
  }, []);

  const Palabra = () => {
    return (
      <div className="row">
        <div className="row align-items-start">
          {cursos.map((item, i) => (
            <div className="form-signin1 col" key={"CursoEstudiante" + i}>
              <div className="text-center">
                <div className="card card-body">
                  <h1 className="text-left" key={"album" + i} value={i}>
                    {item}
                  </h1>
                  <div>
                    <span
                      className="input-group-text"
                      id="validationTooltipUsernamePrepend"
                    ></span>{" "}
                    <br />
                    {publicacion.map((ite, j) => {
                      if (ite.Curso === item) {
                        return (
                          <div key={"p-" + i}>
                            <label className="input-group-text" key={"lib" + j}>
                              {ite.Contenido}
                            </label>
                          </div>
                        );
                      }
                    })}
                  </div>
                </div>
                <br />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  const libro = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante/libro", "_self");
  };

  const principal = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Alumno - Publicaciones</h1>
          <br />
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-success"
              onClick={principal}
            >
              Principal
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={libro}>
              Ver Libros
            </button>
          </center>
          <br />
          {cursos.length > 0 ? <Palabra /> : null}
          <br />
          <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
          <br />
        </form>
      </div>
    </div>
  );
}

export default Publicacion_Estudiante;