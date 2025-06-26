/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseSuccessPaged struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewSuccessResponse(c *gin.Context, message string, data interface{}) {
	response := ResponseSuccess{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func NewSuccessPagedResponse(c *gin.Context, message string, data interface{}, page, size, total int) {
	response := ResponseSuccessPaged{
		Status:  "success",
		Message: message,
		Data:    data,
		Page:    page,
		Size:    size,
		Total:   total,
	}
	c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c *gin.Context, code int, message string) {
	response := ResponseError{
		Status:  "failed",
		Message: message,
	}
	c.JSON(code, response)
}
