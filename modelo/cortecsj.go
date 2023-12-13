package modelo

type CorteCsj struct {
	Id                 string `json:"id"`
	NombreCorte     string `json:"nombre_corte"`
	NombreServicio    string `json:"nombre_servicio"`
	PasswordDb         string `json:"password_db"`
	NombreDb           string `json:"nombre_db"`
	Ip                 string `json:"ip"`
	UserId             string `json:"user_id"`
	Puerto             string `json:"puerto"`
	NombreServicioTest string `json:"nombre_servicio_test,omitempty"`
	PasswordDbTest     string `json:"password_db_test,omitempty"`
	NombreDbTest       string `json:"nombre_db_test,omitempty"`
	IpTest             string `json:"ip_test,omitempty"`
	UserIdTest         string `json:"user_id_test,omitempty"`
	PuertoTest         string `json:"puerto_test,omitempty"`
}
