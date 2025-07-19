/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package user

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/repository"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	Create(req request.CreateUserRequest) error
}

type usecase struct {
	db                   *gorm.DB
	userRepo             repository.UserRepository
	userDoctorRepo       repository.UserDoctorRepository
	roleRepo             repository.RoleRepository
	employmentStatusRepo repository.EmploymentStatusRepository
	DepartmentRepo       repository.DepartmentRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                   db,
		userRepo:             service.NewUserService(db),
		roleRepo:             service.NewRoleService(db),
		employmentStatusRepo: service.NewEmploymentStatusService(db),
		DepartmentRepo:       service.NewDepartmentService(db),
		userDoctorRepo:       service.NewUserDoctorService(db),
	}
}
