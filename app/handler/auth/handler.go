/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package auth

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/app/context/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	uc auth.Usecase
}

type Handler interface {
	Login(ctx *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		uc: auth.NewUsecase(db),
	}
}
