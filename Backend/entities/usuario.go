package entities

type Usuario struct {
	ID              string `json:"id"`
	NombreUsuario   string `json:"nombre_usuario"`
	PasswordUsuario string `json:"password_usuario"`
}
type CrearUsuario struct {
	NombreUsuario   string `json:"nombre_usuario" binding:"required"`
	PasswordUsuario string `json:"password_usuario" binding:"required"`
}
