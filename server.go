/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/berrylradianh/makerble-golang-coding-assesment/library"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	database "github.com/berrylradianh/makerble-golang-coding-assesment/library/config/database"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/middleware/auth"
	appRoutes "github.com/berrylradianh/makerble-golang-coding-assesment/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
)

var env config.Config

func startApp() {
	err := auth.NewMiddlewareConfig(env)
	if err != nil {
		fmt.Println(err)
	}

	SQLMasterConn, err := database.InitDBSQL(env, "postgresql")
	if err != nil {
		fmt.Println(err)
	}

	SQLSlaveConn := SQLMasterConn
	if len(env.GetString(`postgresql_slave.user`)) > 0 {
		SQLSlaveConn, err = database.InitDBSQL(env, "postgresql_slave")
		if err != nil {
			fmt.Println(err)
		}
	}

	//start gin engine
	engine := gin.New()

	//add swagger documentation
	swagger(engine)

	//add route endpoint and healthcheck
	healthCheck(engine)

	//call route per module
	req := library.RouteInit{Engine: engine, SQLMaster: SQLMasterConn, SQLSlave: SQLSlaveConn, Env: env}
	appRoutes.NewRouteInit(req)

	//run server
	serverPort := env.GetString("server.address")
	log.Println("run on port ", serverPort)
	fmt.Printf("App running on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, engine); err != nil {
		log.Fatal(err)
	}
}

func healthCheck(engine *gin.Engine) {
	// root endpoint
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "CSW API Service")
	})

	// Healthcheck endpoint
	engine.GET("ping",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong!")
		})
}

func swagger(engine *gin.Engine) {
	if env.GetString("server.env") == "stage" {
		docs.SwaggerInfo.Host = "stage.api.csw.id"
	}

	if env.GetString("server.env") == "local" {
		docs.SwaggerInfo.Host = "127.0.0.1:5000"
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
