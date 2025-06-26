/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// ValidateRequestBody performs binding and validation of JSON request data.
// It returns an error if binding or validation fails.
func ValidateRequestBody(ctx *gin.Context, request interface{}) error {
	if err := ctx.ShouldBindJSON(request); err != nil {
		return fmt.Errorf("failed binding data: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	return nil
}

// ValidateURLParams validates the parameters in the URL.
// This function returns an error if the validation fails.
func ValidateURLParams(ctx *gin.Context, params interface{}) error {
	if err := ctx.ShouldBindUri(params); err != nil {
		return fmt.Errorf("failed binding URI parameters: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return fmt.Errorf("parameter validation error: %v", err)
	}

	return nil
}

// ValidateQueryParams validates the query parameters in the URL.
// This function returns an error if the validation fails.
func ValidateQueryParams(ctx *gin.Context, params interface{}) error {
	if err := ctx.ShouldBindQuery(params); err != nil {
		return fmt.Errorf("failed binding Query parameters: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return fmt.Errorf("parameter validation error: %v", err)
	}

	return nil
}
