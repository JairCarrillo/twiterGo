package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/JairCarrillo/twiterGo/bd"
	"github.com/JairCarrillo/twiterGo/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JTWSsign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := string.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error))  {
		return miClave nil
	})
	if err == nil {
		// rutina que chequea contra la BD
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Tokenn Invalido")
	}

	return &claims, false, string(""), err
}