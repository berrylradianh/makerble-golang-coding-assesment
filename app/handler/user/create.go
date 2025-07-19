/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package user

import (
	"net/http"

	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/middleware/auth"
	"github.com/gin-gonic/gin"
)

func (h *handler) Create(ctx *gin.Context) {
	var req request.CreateUserRequest

	authenticatedUserID, err := auth.GetAuthenticatedUser(ctx.Request)
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	req.AuthenticatedUserID = authenticatedUserID

	if err := helper.ValidateRequestBody(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.uc.Create(req)
	if err == helper.ErrForbidden {
		helper.NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helper.NewSuccessResponse(ctx, "you have successfully created a new user", nil)
}
