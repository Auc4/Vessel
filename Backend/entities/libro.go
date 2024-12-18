package entities

type Libro struct {
	LibroID          string `json:"libro_ID"`
	Usuario_ID       int    `json:"usuario_ID"`
	TituloLibro      string `json:"titulo"`
	AutorLibro       string `json:"autor"`
	FechaPublicacion int16  `json:"año_publicacion"`
	Favorito         bool   `json:"favorito"`
}

type FocusedLibro struct {
	TituloLibro      string   `json:"titulo"`
	AutorLibro       string   `json:"autor"`
	FechaPublicacion int16    `json:"año_publicacion"`
	Favorito         bool     `json:"favorito"`
	Etiquetas        []string `json:"etiquetas"`
}
