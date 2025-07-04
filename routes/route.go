/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package routes

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/container"
	request "github.com/berrylradianh/makerble-golang-coding-assesment/library"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/middleware/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func NewRouteInit(req request.RouteInit) {
	moduleDI := container.NewHandlerContainer(req.SQLMaster, req.SQLSlave, req.Env)

	route := req.Engine.Group("api/v1")
	route.Use(apmgin.Middleware(req.Engine))
	route.Use(cors.CORSMiddleware())
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	route.OPTIONS("/*path", cors.CORSMiddleware())

	{
		authRoute := route.Group("auth")
		authRoute.POST("login", moduleDI.Auth.Login)
		authRoute.POST("logout", moduleDI.Auth.Logout)
	}
}
