/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package user

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/app/context/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	uc user.Usecase
}

type Handler interface {
	Create(ctx *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		uc: user.NewUsecase(db),
	}
}
