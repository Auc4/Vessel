package entities

type Usuario struct {
	NombreUsuario   string `json:"Nombre_usuario"`
	PasswordUsuario string `json:"Password_usuario"`
	Email           string `json:"Email"`
}
