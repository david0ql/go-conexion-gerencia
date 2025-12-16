# conexion-gerencia-go

Libreria en Go para desencriptar datos de conexion a base de datos usando AES-256-CBC.

## Instalacion

```bash
go get github.com/amovil/conexion-gerencia-go
```

## Uso

```go
package main

import (
    "fmt"
    "log"
    
    conexiongerencia "github.com/amovil/conexion-gerencia-go"
)

func main() {
    nbAgenteComercial := "tu-empresa"
    baseURL := "https://api.example.com/endpoint"
    key := "tu-clave-de-encriptacion"
    
    gerencia := conexiongerencia.NewGerenciaDecrypt(nbAgenteComercial, baseURL, key)
    
    connection, err := gerencia.Do()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Host: %s\n", connection.Host)
    fmt.Printf("Port: %d\n", connection.Port)
    fmt.Printf("BD: %s\n", connection.BD)
}
```

## API

### NewGerenciaDecrypt

Crea una nueva instancia de GerenciaDecrypt.

```go
func NewGerenciaDecrypt(nbAgenteComercial, baseURL, key string) *GerenciaDecrypt
```

**Parametros:**
- `nbAgenteComercial`: Nombre de la empresa/agente comercial
- `baseURL`: URL base del endpoint para obtener los datos encriptados
- `key`: Clave de encriptacion AES-256

### Do

Ejecuta el proceso completo: obtiene los datos encriptados del servidor y los desencripta.

```go
func (g *GerenciaDecrypt) Do() (*DatabaseConnectionGerencia, error)
```

**Retorna:**
- `*DatabaseConnectionGerencia`: Estructura con los datos de conexion (host, port, bd)
- `error`: Error si ocurre algun problema durante el proceso

## Estructuras

### DatabaseConnectionGerencia

```go
type DatabaseConnectionGerencia struct {
    Host string `json:"host"`
    Port int    `json:"port"`
    BD   string `json:"bd"`
}
```

## Requisitos

- Go 1.21 o superior
