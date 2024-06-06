package models

type Usuario struct {
	Id                 uint      `json:"id" db:"id"`                                 //Identificar usuario de otro
	Nombre_usuario     string    `json:"nombre_usuario" db:"nombre_usuario"`         //Nombre de usuario
	Contrasena         string    `json:"contrasena" db:"contrasena"`                 //Contraseña de usuario
	Correo_electronico string    `json:"correo_electronico" db:"correo_electronico"` //Correo electronico del usuario
	Reservas           []Reserva `json:"reservas" db:"reservas"`                     // Lista de reservas del usuario
}

type Reserva struct {
	Id          uint     `json:"id" db:"id"`                   // Identificador de la reserva
	Descripcion string   `json:"descripcion" db:"descripcion"` // Descripción de la reserva
	VehiculoId  uint     `json:"vehiculo_id" db:"vehiculo_id"` // Identificador del vehículo reservado
	Vehiculo    Vehiculo `json:"vehiculo" db:"vehiculo"`       // Detalles del vehículo reservado
}
