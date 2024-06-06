package models

/**************************************ESTRUCTURA PARA ALMACENAR LOS VEHICULOS**************************************************/
//Definición de estructura de los vehiculos
type Vehiculo struct {
	Id            uint   `json:"id" db:"id"`                       //Identificar vehiculo de otro
	Type          string `json:"typ" db:"typ"`                     //Tipo del vehiculo
	Color         string `json:"color" db:"color"`                 //Color del vehiculo
	Model         string `json:"model" db:"model"`                 //Modelo del vehiculo
	Brand         string `json:"brand" db:"brand"`                 //Marca del vehiculo
	Photo         string `json:"photo" db:"photo"`                 //Foto del Vehiculo
	Price         uint   `json:"price" db:"price"`                 //Precio del vehiculo
	Disponibility bool   `json:"disponibility" db:"disponibility"` //Disponibilidad del vehiculo
}
