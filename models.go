package conexiongerencia

// RequestData representa los datos de la peticion HTTP
type RequestData struct {
	Empresa string `json:"empresa"`
}

// ResponseData representa la respuesta del servidor
type ResponseData struct {
	Content string `json:"content"`
	IV      string `json:"iv"`
}

// DatabaseConnectionGerencia representa la conexion a la base de datos desencriptada
type DatabaseConnectionGerencia struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	BD   string `json:"bd"`
}
