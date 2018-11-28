package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonzalonegre/sandbox1/lib/controllers"
)

func init() {
	router := gin.New()
	v1 := router.Group("api/v1")

	dias := v1.Group("/clima")
	dias.POST("/", controllers.CreateDias)
	dias.GET("/:dia", controllers.GetDia)
	dias.GET("/", controllers.GetDias)
	dias.PUT("/:dia", controllers.UpdateDia)
	dias.DELETE("/", controllers.DeleteDias)

	http.Handle("/", router)
}