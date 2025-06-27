/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package request

import "time"

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterRequest struct {
	Name                 string    `json:"name" form:"name" validate:"required"`
	Email                string    `json:"email" form:"email" validate:"required,email"`
	IdentityNumber       string    `json:"identity_number" form:"identity_number" validate:"required"`
	Phone                string    `json:"phone" form:"phone" validate:"required"`
	DateOfBirth          time.Time `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Address              string    `json:"address" form:"address" validate:"required"`
	Gender               string    `json:"gender" form:"gender" validate:"required"`
	Password             string    `json:"password" form:"password" validate:"required,min=8"`
	ConfirmationPassword string    `json:"confirmation_password" form:"confirmation_password" validate:"required,eqfield=Password"`
}
