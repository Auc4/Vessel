package entities

type Libro struct {
	ID               string `json:"id"`
	IDUsuario        string `json:"id_usuario" binding:"required"`
	TituloLibro      string `json:"titulo_libro" binding:"required"`
	AutorLibro       string `json:"autor_libro" binding:"required"`
	FechaPublicacion string `json:"fecha_publicacion" binding:"required"`
	Favorito         bool   `json:"favorito"`
}

type CrearLibro struct {
	IDUsuario        string `json:"id_usuario" binding:"required"`
	TituloLibro      string `json:"titulo_libro" binding:"required"`
	AutorLibro       string `json:"autor_libro" binding:"required"`
	FechaPublicacion string `json:"fecha_publicacion" binding:"required"`
	Favorito         bool   `json:"favorito"`
}
