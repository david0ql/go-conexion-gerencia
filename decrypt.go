package conexiongerencia

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
)

// DecryptService maneja la desencriptacion de datos usando AES-256-CBC
type DecryptService struct {
	key []byte
}

// NewDecryptService crea una nueva instancia de DecryptService
func NewDecryptService(key string) *DecryptService {
	return &DecryptService{
		key: []byte(key),
	}
}

// Decrypt desencripta los datos usando AES-256-CBC
func (s *DecryptService) Decrypt(data *ResponseData) (*DatabaseConnectionGerencia, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(data.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 content: %w", err)
	}

	iv := []byte(data.IV)

	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("IV length must be %d bytes, got %d", aes.BlockSize, len(iv))
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// Remover padding PKCS7
	decrypted = s.removePKCS7Padding(decrypted)

	// Estructura intermedia para manejar port como string o int
	var temp struct {
		Host string      `json:"host"`
		Port interface{} `json:"port"` // Puede ser string o int
		BD   string      `json:"bd"`
	}

	if err := json.Unmarshal(decrypted, &temp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal decrypted data: %w", err)
	}

	// Convertir port a int
	var port int
	switch v := temp.Port.(type) {
	case float64:
		port = int(v)
	case int:
		port = v
	case string:
		var err error
		port, err = s.parsePort(v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse port: %w", err)
		}
	default:
		return nil, fmt.Errorf("unexpected port type: %T", temp.Port)
	}

	databaseConnection := DatabaseConnectionGerencia{
		Host: temp.Host,
		Port: port,
		BD:   temp.BD,
	}

	return &databaseConnection, nil
}

// removePKCS7Padding remueve el padding PKCS7
func (s *DecryptService) removePKCS7Padding(data []byte) []byte {
	if len(data) == 0 {
		return data
	}

	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		return data
	}

	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return data
		}
	}

	return data[:len(data)-padding]
}

// parsePort convierte un string a int para el puerto
func (s *DecryptService) parsePort(portStr string) (int, error) {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, fmt.Errorf("invalid port format: %s", portStr)
	}
	return port, nil
}
