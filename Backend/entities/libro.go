package entities

type Libro struct {
	IDUsuario        string `json:"id_usuario"`
	TituloLibro      string `json:"titulo_libro"`
	AutorLibro       string `json:"autor_libro"`
	FechaPublicacion string `json:"fecha_publicacion"`
	Favorito         bool   `json:"favorito"`
}
