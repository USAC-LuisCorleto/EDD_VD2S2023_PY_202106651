import React from "react";

function Principal_Admin() {
  const archivos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/archivo", "_self");
  };

  const alumnos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/alumnos", "_self");
  };

  const libros = (e) => {
    e.preventDefault();
    window.open("/principal/admin/libros", "_self");
  };

  const reporte = (e) => {
    e.preventDefault();
    window.open("/principal/admin/reporte", "_self");
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body bg-light">
          <h1 className="h3 mb-3 fw-normal">BIENVENIDO ADMINISTRADOR</h1>
          <br />
          <center>
            <button className="w-75 btn btn-primary" onClick={archivos}>
              Carga Archivos
            </button>
          </center>
          <br />
          <center>
            <button className="w-75 btn btn-primary" onClick={alumnos}>
              Ver Alumnos
            </button>
          </center>
          <br />
          <center>
            <button className="w-75 btn btn-primary" onClick={libros}>
              Ver Libros
            </button>
          </center>
          <br />
          <center>
            <button className="w-75 btn btn-primary" onClick={reporte}>
              Reportes
            </button>
          </center>
          <br />
          <center>
            <button className="w-75 btn btn-primary" onClick={salir}>
              Salir
            </button>
          </center>
        </form>
      </div>
    </div>
  );
}

export default Principal_Admin;