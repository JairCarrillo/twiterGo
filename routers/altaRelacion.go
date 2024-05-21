package routers

import (
	"context"

	"github.com/JairCarrillo/twiterGo/bd"
	"github.com/JairCarrillo/twiterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

// AltaRelacion maneja la solicitud para crear una relación entre usuarios.
// Retorna una respuesta con el estado de la operación.
func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	// Obtener el ID del usuario con el que se establecerá la relación.
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	// Crear una estructura Relacion con los IDs de los usuarios.
	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	// Insertar la relación en la base de datos.
	status, err := bd.InsertoRelacion(t)
	if err != nil {
		r.Message = "Error al insertar la relación: " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar la relación"
		return r
	}

	// Configurar respuesta exitosa.
	r.Status = 200
	r.Message = "Alta de Relación OK"
	return r
}
