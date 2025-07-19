/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package helper

import "errors"

var (
	ErrDuplicatePhone         = errors.New("your phone number is already registered")
	ErrDuplicateEmail         = errors.New("your email is already registered")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrEmailNotFound          = errors.New("email not found")
	ErrTokenExpired           = errors.New("Your token has expired, please login again")
	ErrForbidden              = errors.New("you don't have permission to access this feature")
)

const (
	_ = iota
	RoleAdminID
	RoleReceptionistID
	RoleDoctorID
	RolePatientID
	RoleSuperAdminID
)
