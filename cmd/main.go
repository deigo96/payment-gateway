package main

import (
	"account-service/api/route"
	"account-service/bootstrap"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	gin.SetMode(app.Env.GinMode)

	db := app.Db
	defer app.CloseDBConnection()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	handler := gin.Default()

	handler.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	handler.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	routerV1 := handler.Group("")

	route.Setup(app.Env, timeout, *db, routerV1)

	host := app.Env.ServerHost
	port := app.Env.ServerPort

	handler.Run(host+ ":" + port)
}
