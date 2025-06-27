/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package auth

import (
	"net/http"

	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) Login(ctx *gin.Context) {
	var req request.LoginRequest

	if err := helper.ValidateRequestBody(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.uc.Login(req)
	if err == helper.ErrInvalidEmailOrPassword {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, helper.ErrInvalidEmailOrPassword.Error())
		return
	}
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helper.NewSuccessResponse(ctx, "you have successfully logged in", res)
}
