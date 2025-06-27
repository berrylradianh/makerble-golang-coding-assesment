/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package auth

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/middleware/auth"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/repository"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	Login(req request.LoginRequest) (*auth.TokenResponse, error)
}

type usecase struct {
	db       *gorm.DB
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:       db,
		userRepo: service.NewUserService(db),
		roleRepo: service.NewRoleService(db),
	}
}
