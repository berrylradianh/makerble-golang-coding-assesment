/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package auth

import (
	"net/http"
	"strings"

	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) Logout(ctx *gin.Context) {
	var req request.LogoutRequest

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		helper.NewErrorResponse(ctx, http.StatusUnauthorized, "You must be logged in")
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, "Invalid Authorization header")
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	req.Token = token
	err := h.uc.Logout(req)
	if err == helper.ErrTokenExpired {
		helper.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helper.NewSuccessResponse(ctx, "you have successfully logged out", nil)
}
