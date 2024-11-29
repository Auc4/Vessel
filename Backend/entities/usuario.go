package entities

type Usuario struct {
	NombreUsuario   string `json:"nombre_usuario"`
	PasswordUsuario string `json:"password_usuario"`
	Email           string `json:"email"`
}
