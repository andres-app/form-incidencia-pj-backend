package modelo

type Data struct {
	NombreSistema  string   `json:"nombreSistema,omitempty"`
	VersionSistema string   `json:"versionSistema,omitempty"`
	UsuarioNombre  string   `json:"usuarioNombre,omitempty"`
	CelularAxeso   string   `json:"celularAxeso,omitempty"`
	Sede           string   `json:"sede,omitempty"`
	Cargo          string   `json:"cargo,omitempty"`
	PersonaReporte string   `json:"personaReporte,omitempty"`
	CelularPersona string   `json:"celularPersona,omitempty"`
	Fecha          string   `json:"fecha,omitempty"`
	Descripcion    string   `json:"descripcion,omitempty"`
	PreguntaUno    string   `json:"preguntaUno,omitempty"`
	PreguntaDos    string   `json:"preguntaDos,omitempty"`
	PreguntaTres   string   `json:"preguntaTres,omitempty"`
	PreguntaCuatro string   `json:"preguntaCuatro,omitempty"`
	PreguntaCinco  string   `json:"preguntaCinco,omitempty"`
	Imagenes       []Imagen `json:"imagenes"`
	IdCorteCsj     string   `json:"idCorteCsj,omitempty"`
}
