/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package library

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteInit struct {
	Engine    *gin.Engine
	SQLMaster *gorm.DB
	SQLSlave  *gorm.DB
	Env       config.Config
}
