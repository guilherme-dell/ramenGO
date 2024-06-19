package main

import (
	"os"
	"ramenGO/cmd/api/middleware"
	"ramenGO/cmd/api/routes"
	inmemory "ramenGO/internal/repositories/in-memory"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	middlewareKEY := os.Getenv("MIDDLEWARE_KEY")
	orderMiddlewareKey := os.Getenv("ORDER_KEY")

	listenPort := "0.0.0.0:" + port

	rep := inmemory.NewDataBaseInMemoryRepository(&orderMiddlewareKey)
	rep.PopulateInMemoryRepository()

	r := gin.Default()

	r.Use(middleware.ConfigCORS())
	r.Use(middleware.ApiKeyAuth(middlewareKEY))

	routes.InitRoutes(&r.RouterGroup, rep)

	r.Run(listenPort)
}
