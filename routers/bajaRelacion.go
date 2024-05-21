package routers

import (
	"github.com/JairCarrillo/twiterGo/bd"
	"github.com/JairCarrillo/twiterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

// BajaRelacion maneja la solicitud para eliminar una relación entre usuarios.
// Retorna una respuesta con el estado de la operación.
func BajaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	// Obtener el ID del usuario con el que se eliminará la relación.
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	// Crear una estructura Relacion con los IDs de los usuarios.
	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	// Eliminar la relación en la base de datos.
	status, err := bd.BorroRelacion(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar borrar relación: " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado borrar relación"
		return r
	}

	// Configurar respuesta exitosa.
	r.Status = 200
	r.Message = "Baja Relación OK!"
	return r
}
