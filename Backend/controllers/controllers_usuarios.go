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
	listQueryUser   = "SELECT id, nombre_usuario, correo_electronico, reservas FROM usuarios limit $1 offset $2" //Servirá para listar todos los usuarios
	readQueryUser   = "SELECT id, nombre_usuario, correo_electronico, reservas FROM usuarios WHERE id = $1"      //placeholder $1 cierto valor
	crearQueryUser  = "INSERT INTO usuarios (nombre_usuario, contrasena, correo_electronico, reservas) VALUES (:nombre_usuario, :contrasena, :correo_electronico, :reservas) RETURNING id"
	updateQueryUser = "UPDATE usuarios SET %s WHERE id=:id" // reserva=:reserva
)

//Capa adicional, los controladores poseen toda la logica que comunica al repositorio con los handlers
//Posee la logica de los servicios

type ControllerUser struct { //El controlador debe estar instanciado al repositorio sql que ya esté conectado a una base de datos SQL
	repo repository.Repository[models.Usuario] //interactuar con una tabla en la cual se almacenan usuarios
} //*******Los metodos de la estructura tendrá metodos que nos permitirá hacer operaciones sobre el repositorio****************

// (repo = Lo que recibo objeto de tipo repositorio) -> Se retorna (puntero al controlador o un error)
func NewControllerUser(repo repository.Repository[models.Usuario]) (*ControllerUser, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es necesario un repositorio valido")
	}
	return &ControllerUser{
		repo: repo,
	}, nil
}

// Metodo de estructura Controller llamado ListarUsuarios recibe una estructura de byte y lo devuelve en formato JSON
func (c *ControllerUser) ListarUsuarios(limit, offset int) ([]byte, error) {
	usuarios, _, err := c.repo.List(context.Background(), listQueryUser, limit, offset)
	if err != nil {
		log.Printf("fallo al leer los usuarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer usuarios, con error: %s", err.Error())
	}
	jsonUsuarios, err := json.Marshal(usuarios)
	if err != nil {
		log.Printf("fallo al codificar Json: %s", err.Error())
		return nil, fmt.Errorf("fallo al codificar Json: %s", err.Error())
	}
	return jsonUsuarios, nil
}

// Metodo de estructura Controller llamado traer usuario
func (c *ControllerUser) TraerUsuarios(id string) ([]byte, error) {
	usuario, err := c.repo.Read(context.Background(), readQueryUser, id)
	if err != nil {
		log.Printf("fallo al leer usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer usuario, con error: %s", err.Error())
	}
	Usuariojson, err := json.Marshal(usuario)
	if err != nil {
		log.Printf("fallo al codificar Json: %s", err.Error())
		return nil, fmt.Errorf("fallo al codificar Json: %s", err.Error())
	}
	return Usuariojson, nil
}

// Metodo de estructura Controller llamado crear usuario
func (c *ControllerUser) CrearUsuario(body []byte) (int64, error) { //El body debe convertirse de un objeto JSON a un objeto usuario
	nuevoUsuario := &models.Usuario{}
	err := json.Unmarshal(body, nuevoUsuario)
	if err != nil {
		log.Printf("fallo al crear usuario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear usuario, con error: %s", err.Error())
	}
	valoresColumnas := map[string]any{
		"nombre_usuario":     nuevoUsuario.Nombre_usuario,
		"contrasena":         nuevoUsuario.Contrasena,
		"correo_electronico": nuevoUsuario.Correo_electronico,
		"reservas":           nuevoUsuario.Reservas,
	}

	nuevoId, err := c.repo.Create(context.Background(), crearQueryUser, valoresColumnas)
	if err != nil {
		log.Printf("fallo al crear usuario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear usuario, con error: %s", err.Error())
	}
	return nuevoId, nil
}

//Metodo de estructura controller llamado actualizar usuario

func (c *ControllerUser) ActualizarUsuario(body []byte, id string) error {
	//convertir el body que viene de la operación de tipo PUT del handler
	//Desempacar el JSON en un objeto de tipo mapa
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}
	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un usuario, no hay datos en el body")
		return fmt.Errorf("fallo al actualizar un usuario, no hay datos en el body")
	}
	//actualizar el query
	updtQuery := fmt.Sprintf(updateQueryUser, buildUpdateQueryUser(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	err = c.repo.Update(context.Background(), updtQuery, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}
	return nil
}

func buildUpdateQueryUser(columnasActualizar map[string]any) string {
	//recorrer el mapa y agregar los keys
	columnas := []string{}
	for key := range columnasActualizar {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	return strings.Join(columnas, ",")
}
