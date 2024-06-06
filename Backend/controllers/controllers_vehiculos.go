package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/JonnierMB/DW_S5/Backend/models"
	"github.com/JonnierMB/DW_S5/Backend/repository"
)

var (
	listQuery   = "SELECT id, typ, color, model, brand, photo, price, disponibility FROM vehiculos limit $1 offset $2" //Servirá para listar todos los vehiculos
	readQuery   = "SELECT id, typ, color, model, brand, photo, price, disponibility FROM vehiculos WHERE id = $1"      //placeholder $1 cierto valor
	updateQuery = "UPDATE vehiculos SET %s WHERE id=:id"                                                               // time=:time, comment=:comment. reactions=:reactions
)

//Capa adicional, los controladores poseen toda la logica que comunica al repositorio con los handlers
//Posee la logica de los servicios

type Controller struct { //El controlador debe estar instanciado al repositorio sql que ya esté conectado a una base de datos SQL
	repo repository.Repository[models.Vehiculo] //interactuar con una tabla en la cual se almacenan los vehiculos
} //*******La estructura tendrá metodos que nos permitirá hacer operaciones sobre el repositorio****************

// (repo = Lo que recibo objeto de tipo repositorio) -> Se retorna (puntero al controlador o un error)
func NewController(repo repository.Repository[models.Vehiculo]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es necesario un repositorio valido")
	}
	return &Controller{
		repo: repo,
	}, nil
}

// Metodo de estructura Controller llamado ListarVehiculos recibe una estructura de byte y lo devuelve en formato JSON
func (c *Controller) ListarVehiculos(limit, offset int) ([]byte, error) {
	vehiculos, _, err := c.repo.List(context.Background(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer Vehiculos, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer Vehiculos, con error: %s", err.Error())
	}
	jsonVehiculos, err := json.Marshal(vehiculos)
	if err != nil {
		log.Printf("fallo al codificar Json: %s", err.Error())
		return nil, fmt.Errorf("fallo al codificar Json: %s", err.Error())
	}
	return jsonVehiculos, nil
}

// Metodo de estructura Controller llamado traer Vehiculo
func (c *Controller) TraerVehiculos(id string) ([]byte, error) {
	vehiculo, err := c.repo.Read(context.Background(), readQuery, id)
	if err != nil {
		log.Printf("fallo al traer Vehiculo, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al traer Vehiculo, con error: %s", err.Error())
	}
	Vehiculojson, err := json.Marshal(vehiculo)
	if err != nil {
		log.Printf("fallo al codificar Json: %s", err.Error())
		return nil, fmt.Errorf("fallo al codificar Json: %s", err.Error())
	}
	return Vehiculojson, nil
}

//Metodo de estructura controller llamado actualizar Actualizar Vehiculo

func (c *Controller) ActualizarVehiculo(body []byte, id string) error {
	//convertir el body que viene de la operación de tipo PUT del handler
	//Desempacar el JSON en un objeto de tipo mapa
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un Vehiculo, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un Vehiculo, con error: %s", err.Error())
	}
	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un Vehiculo, no hay datos en el body")
		return fmt.Errorf("fallo al actualizar un Vehiculo, no hay datos en el body")
	}
	//actualizar el query
	updtQuery := fmt.Sprintf(updateQuery, buildUpdateQuery(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	err = c.repo.Update(context.Background(), updtQuery, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un Vehiculo, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un Vehiculo, con error: %s", err.Error())
	}
	return nil
}

func buildUpdateQuery(columnasActualizar map[string]any) string {
	//recorrer el mapa y agregar los keys
	columnas := []string{}
	for key := range columnasActualizar {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	return strings.Join(columnas, ",")
}
