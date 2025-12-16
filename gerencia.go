package conexiongerencia

// GerenciaDecrypt es la clase principal que obtiene y desencripta los datos de conexion
type GerenciaDecrypt struct {
	nbAgenteComercial string
	getDatabaseService *GetDatabaseService
	decryptService     *DecryptService
}

// NewGerenciaDecrypt crea una nueva instancia de GerenciaDecrypt
func NewGerenciaDecrypt(nbAgenteComercial, baseURL, key string) *GerenciaDecrypt {
	return &GerenciaDecrypt{
		nbAgenteComercial:  nbAgenteComercial,
		getDatabaseService: NewGetDatabaseService(baseURL),
		decryptService:      NewDecryptService(key),
	}
}

// Do ejecuta el proceso completo: obtiene los datos encriptados y los desencripta
func (g *GerenciaDecrypt) Do() (*DatabaseConnectionGerencia, error) {
	responseData, err := g.getDatabaseService.GetDatabase(g.nbAgenteComercial)
	if err != nil {
		return nil, err
	}

	return g.decryptService.Decrypt(responseData)
}
