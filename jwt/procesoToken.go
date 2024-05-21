package jwt

import (
	"errors"
	"strings"

	"github.com/JairCarrillo/twiterGo/bd"
	"github.com/JairCarrillo/twiterGo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

// Variables globales para almacenar el email y el ID del usuario extraídos del token.
var Email string
var IDUsuario string

// ProcesoToken valida y procesa un token JWT.
func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	// Convertir la clave de firma JWT a un byte array.
	miClave := []byte(JWTSign)
	var claims models.Claim

	// Separar el token en partes usando "Bearer" como delimitador.
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("formato de token inválido")
	}

	// Limpiar espacios en blanco del token.
	tk = strings.TrimSpace(splitToken[1])

	// Parsear el token con las reclamaciones.
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	// Si no hay errores al parsear y el token es válido, se verifica contra la base de datos.
	if err == nil {
		// Verificar si el usuario existe en la base de datos.
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return &claims, encontrado, IDUsuario, nil
	}

	// Si el token no es válido, retornar un error.
	if !tkn.Valid {
		return &claims, false, "", errors.New("token inválido")
	}

	// Retornar cualquier otro error ocurrido durante el proceso.
	return &claims, false, "", err
}
