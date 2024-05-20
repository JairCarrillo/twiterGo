package bd

import (
	"context"

	"github.com/JairCarrillo/twiterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) bool {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	//recortado por sugerencia de go
	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	return err == nil
}
