/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package user

import (
	"fmt"
	"strings"

	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/model"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) Create(req request.CreateUserRequest) error {
	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"id": req.AuthenticatedUserID,
	})
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	role, err := uc.roleRepo.FindOneBy(map[string]interface{}{
		"id": user.RoleID,
	})
	if err != nil {
		return fmt.Errorf("failed to find role: %v", err)
	}

	if role.ID == helper.RolePatientID || role.ID == helper.RoleDoctorID {
		return helper.ErrForbidden
	}

	roleID := helper.RolePatientID
	if role.ID == helper.RoleSuperAdminID {
		roleID = helper.RoleAdminID
	}

	userPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed hashing password: %v", err)
	}

	dateOfBirth, err := helper.ConvertStringToTime(req.DateOfBirth)
	if err != nil {
		return fmt.Errorf("failed parsing date of birth: %v", err)
	}

	phone, err := helper.SanitizePhone(req.Phone)
	if err != nil {
		return fmt.Errorf("failed sanitizing phone number: %v", err)
	}

	userModel := &model.User{
		RoleID:         roleID,
		IdentityNumber: req.IdentityNumber,
		Email:          req.Email,
		Password:       string(userPassword),
		Name:           req.Name,
		Phone:          phone,
		DateOfBirth:    dateOfBirth,
		Address:        req.Address,
		Gender:         req.Gender,
	}

	tx := uc.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	newUser, err := uc.userRepo.Create(userModel, tx)
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}

	if role.ID == helper.RoleReceptionistID {
		employmentStatus, err := uc.employmentStatusRepo.FindOneBy(map[string]interface{}{
			"slug": req.EmploymentStatus,
		})
		if err != nil {
			return fmt.Errorf("failed to find employment status: %v", err)
		}

		department, err := uc.DepartmentRepo.FindOneBy(map[string]interface{}{
			"slug": req.Department,
		})
		if err != nil {
			return fmt.Errorf("failed to find department: %v", err)
		}

		employeeNumber, err := helper.GenerateEmployeeNumber(req.Name, req.DateOfBirth)
		if err != nil {
			return fmt.Errorf("failed generating employee number: %v", err)
		}

		professionalDegree := "{" + strings.Join(req.ProfessionalDegree, ",") + "}"
		language := "{" + strings.Join(req.Language, ",") + "}"
		professionalCertification := "{" + strings.Join(req.ProfessionalCertification, ",") + "}"
		userDoctorModel := &model.UserDoctor{
			UserID:                    newUser.ID,
			EmploymentStatusID:        employmentStatus.ID,
			DepartmentID:              department.ID,
			ProfessionalDegree:        professionalDegree,
			Citizenship:               req.Citizenship,
			Language:                  language,
			GmcNumber:                 req.GMCNumber,
			Specialization:            req.Specialization,
			SubSpecialization:         req.SubSpecialization,
			LastEducation:             req.LastEducation,
			YearGraduated:             req.YearGraduated,
			ProfessionalCertification: &professionalCertification,
			EmployeeNumber:            employeeNumber,
			IsActive:                  true,
		}

		if _, err := uc.userDoctorRepo.Create(userDoctorModel, tx); err != nil {
			return fmt.Errorf("failed creating user doctor: %v", err)
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed committing transaction: %w", err)
	}

	return nil
}
