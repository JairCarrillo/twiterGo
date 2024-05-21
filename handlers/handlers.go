package handlers

import (
	"context"
	"fmt"

	"github.com/JairCarrillo/twiterGo/jwt"
	"github.com/JairCarrillo/twiterGo/models"
	"github.com/JairCarrillo/twiterGo/routers"
	"github.com/aws/aws-lambda-go/events"
)

// Manejadores es la función principal que maneja las solicitudes entrantes.
func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	r.Status = 400

	// Verifica la autorización de la solicitud.
	isOK, statusCode, msg, claim := validoAuthorization(ctx, request)
	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	// Maneja las solicitudes POST.
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)
		case "login":
			return routers.Login(ctx)
		case "tweet":
			return routers.GraboTweet(ctx, claim)
		case "altaRelacion":
			return routers.AltaRelacion(ctx, request, claim)
		case "subirAvatar":
			return routers.UploadImage(ctx, "A", request, claim)
		case "subirBanner":
			return routers.UploadImage(ctx, "B", request, claim)
		}
	case "GET":
		// Maneja las solicitudes GET.
		switch ctx.Value(models.Key("path")).(string) {
		case "verperfil":
			return routers.VerPerfil(request)
		case "leoTweets":
			return routers.LeoTweets(request)
		case "obtenerAvatar":
			return routers.ObtenerImagen(ctx, "A", request, claim)
		case "obtenerBanner":
			return routers.ObtenerImagen(ctx, "B", request, claim)
		case "ConsultaRelacion":
			return routers.ConsultaRelacion(request, claim)
		case "listaUsuarios":
			return routers.ListaUsuarios(request, claim)
		case "leoTweetsSeguidores":
			return routers.LeoTweetsSeguidores(request, claim)
		}
	case "PUT":
		// Maneja las solicitudes PUT.
		switch ctx.Value(models.Key("path")).(string) {
		case "modificarPerfil":
			return routers.ModificarPerfil(ctx, claim)
		}
		//
	case "DELETE":
		// Maneja las solicitudes DELETE.
		switch ctx.Value(models.Key("path")).(string) {
		case "eliminarTweet":
			return routers.EliminarTweet(request, claim)
		case "bajaRelacion":
			return routers.BajaRelacion(request, claim)
		}
	}

	r.Message = "Method Invalid"
	return r
}

// validoAuthorization verifica si la solicitud tiene una autorización válida.
func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}
