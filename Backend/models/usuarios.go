package models

type Usuario struct {
	Id                 uint   `json:"id" db:"id"`                                 //Identificar usuario de otro
	Nombre_usuario     string `json:"nombre_usuario" db:"nombre_usuario"`         //Nombre de usuario
	Contrasena         string `json:"contrasena" db:"contrasena"`                 //Contraseña de usuario
	Correo_electronico string `json:"correo_electronico" db:"correo_electronico"` //Correo electronico del usuario
}
