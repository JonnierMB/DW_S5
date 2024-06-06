package handlers

//Solo está la interfaz hacia la aplicación cliente (Donde se reciben solicitudes y se dan respuestas)
import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JonnierMB/DW_S5/Backend/controllers"
	"github.com/gorilla/mux"
)

/***************************************************************HANDLER PARA LOS USUARIOS*************************************************************/
type HandlerUsuarios struct {
	controller *controllers.ControllerUser
}

func NewHandlerUsuarios(controller *controllers.ControllerUser) (*HandlerUsuarios, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &HandlerUsuarios{
		controller: controller,
	}, nil
}

// Traer todos los usuarios almacenados en la base de datos
func (hc *HandlerUsuarios) ListarUsuarios(w http.ResponseWriter, r *http.Request) { //Handler para peticiones GET
	usuarios, err := hc.controller.ListarUsuarios(100, 0)
	if err != nil {
		log.Printf("fallo al leer usuarios, con error: %s", err.Error())
		http.Error(w, "fallo al leer usuarios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") //DEBE IR ANTES QUE WRITE HEADER
	w.WriteHeader(http.StatusOK)
	w.Write(usuarios)
}
func (hc *HandlerUsuarios) CrearUsuario(w http.ResponseWriter, r *http.Request) { //Handler para peticiones POST
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al crear nuevo usuario, con error: %s", err.Error())
		http.Error(w, "fallo al crear nuevo usuario", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	nuevoId, err := hc.controller.CrearUsuario(body)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		http.Error(w, "fallo al leer un usuario", http.StatusBadRequest)
		return
	}
	//avisar al cliente que su usuario quedó bien guardado
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id nuevo usuario: %d", nuevoId)))
}

//Función para obtener un usuario especifico

func (hc *HandlerUsuarios) TraerUsuarios(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	comentario, err := hc.controller.TraerUsuarios(id)
	if err != nil {
		log.Printf("falla al leer un usuario, con error %s", err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("el usuario con id %s no se pudo encontrar", id)))
		return
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(comentario)

}

// Función para actualizar comentario
func (hc *HandlerUsuarios) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al actualizar usuario, con error: %s", err.Error())
		http.Error(w, "fallo al actualizar usuario", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = hc.controller.ActualizarUsuario(body, id)
	if err != nil {
		log.Printf("fallo al actualizar usuario, con error: %s", err.Error())
		http.Error(w, "fallo al actualizar usuario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
