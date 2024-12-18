package entities

type Usuario struct {
	UsuarioID       string `json:"usuario_ID"`
	NombreUsuario   string `json:"nombre_usuario"`
	PasswordUsuario string `json:"password_usuario"`
	Email           string `json:"email"`
}
