package conexiongerencia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetDatabaseService maneja las peticiones HTTP para obtener los datos encriptados
type GetDatabaseService struct {
	baseURL string
	client  *http.Client
}

// NewGetDatabaseService crea una nueva instancia de GetDatabaseService
func NewGetDatabaseService(baseURL string) *GetDatabaseService {
	return &GetDatabaseService{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// GetDatabase realiza una peticion POST para obtener los datos encriptados
func (s *GetDatabaseService) GetDatabase(nbAgenteComercial string) (*ResponseData, error) {
	requestData := RequestData{
		Empresa: nbAgenteComercial,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %w", err)
	}

	req, err := http.NewRequest("POST", s.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &responseData, nil
}
