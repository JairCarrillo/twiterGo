package routers

import (
	"context"

	"github.com/JairCarrillo/twiterGo/bd"
	"github.com/JairCarrillo/twiterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		r.Message = "El parametro ID es obligatorio" + err.Error()
		return r
	}

	if !status {
		r.Message = "no se ha logrado insertar la relación "
		return r
	}

	r.Status = 200
	r.Message = "Alta de Relación OK"
	return r
}
