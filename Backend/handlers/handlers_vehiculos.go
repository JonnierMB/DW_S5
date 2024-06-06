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

/***************************************************************HANDLER PARA LOS VEHICULOS*************************************************************/
type HandlerVehiculos struct {
	controller *controllers.Controller
}

func NewHandlerVehiculo(controller *controllers.Controller) (*HandlerVehiculos, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &HandlerVehiculos{
		controller: controller,
	}, nil
}

// Traer todos los Vehiculos almacenados en la base de datos
func (hc *HandlerVehiculos) ListarVehiculos(w http.ResponseWriter, r *http.Request) { //Handler para peticiones GET
	vehiculos, err := hc.controller.ListarVehiculos(100, 0)
	if err != nil {
		log.Printf("fallo al leer Vehiculos, con error: %s", err.Error())
		http.Error(w, "fallo al leer Vehiculos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") //DEBE IR ANTES QUE WRITE HEADER
	w.WriteHeader(http.StatusOK)
	w.Write(vehiculos)
}

//Función para obtener un comentario especifico

func (hc *HandlerVehiculos) TraerVehiculos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	vehiculo, err := hc.controller.TraerVehiculos(id)
	if err != nil {
		log.Printf("falla al leer un comentario, con error %s", err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("el comentario con id %s no se pudo encontrar", id)))
		return
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(vehiculo)
}

// Función para actualizar comentario
func (hc *HandlerVehiculos) ActualizarVehiculo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al actualizar comentario, con error: %s", err.Error())
		http.Error(w, "fallo al actualizar comentario", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = hc.controller.ActualizarVehiculo(body, id)
	if err != nil {
		log.Printf("fallo al actualizar comentario, con error: %s", err.Error())
		http.Error(w, "fallo al actualizar comentario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
