package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/JairCarrillo/twiterGo/bd"
	"github.com/JairCarrillo/twiterGo/jwt"
	"github.com/JairCarrillo/twiterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

// Login maneja la solicitud de inicio de sesión de un usuario.
// Retorna una respuesta con el estado de la operación y, en caso de éxito, un token JWT.
func Login(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	// Obtener el cuerpo de la solicitud.
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario y/o Contraseña Inválidos: " + err.Error()
		return r
	}

	// Validar que el email no esté vacío.
	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido"
		return r
	}

	// Intentar iniciar sesión con los datos proporcionados.
	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario y/o Contraseña Inválidos"
		return r
	}

	// Generar el token JWT.
	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Error al generar el token: " + err.Error()
		return r
	}

	// Crear la respuesta con el token.
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	// Convertir la respuesta a JSON.
	token, err := json.Marshal(resp)
	if err != nil {
		r.Message = "Error al formatear el token a JSON: " + err.Error()
		return r
	}

	// Crear una cookie con el token.
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookie.String()

	// Configurar la respuesta de la API.
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	// Configurar la respuesta de la función.
	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
