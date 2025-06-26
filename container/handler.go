/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package container

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	"gorm.io/gorm"
)

type HandlerContainer struct {
}

// NewHandlerContainer{ initial value dependency injection for every handler
func NewHandlerContainer(SQLMaster *gorm.DB, sQLSlave *gorm.DB, conf config.Config) HandlerContainer {
	return HandlerContainer{}
}
