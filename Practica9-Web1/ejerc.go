package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	var err error
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.POST("/saludo", func(ctx *gin.Context) {
		var persona Persona
		err := ctx.BindJSON(&persona)
		if err != nil {
			ctx.String(400, err.Error())
			return
		}
		ctx.String(200, "Hola "+persona.Nombre+" "+persona.Apellido)
	})

	if err = router.Run(":8080"); err != nil {
		panic(err)
	}
}
