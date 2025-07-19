/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package request

type CreateUserRequest struct {
	AuthenticatedUserID       int      `json:"authenticated_user_id" form:"authenticated_user_id" validate:"required"`
	IdentityNumber            string   `json:"identity_number" form:"identity_number" validate:"required"`
	Email                     string   `json:"email" form:"email" validate:"required,email"`
	Password                  string   `json:"password" form:"password" validate:"required,min=8"`
	Name                      string   `json:"name" form:"name" validate:"required"`
	Phone                     string   `json:"phone" form:"phone" validate:"required"`
	DateOfBirth               string   `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Address                   string   `json:"address" form:"address" validate:"required"`
	Gender                    string   `json:"gender" form:"gender" validate:"required"`
	EmploymentStatus          string   `json:"employment_status" form:"employment_status" validate:"required"`
	Department                string   `json:"department" form:"department" validate:"required"`
	ProfessionalDegree        []string `json:"professional_degree" form:"professional_degree" validate:"required"`
	Citizenship               string   `json:"citizenship" form:"citizenship" validate:"required"`
	Language                  []string `json:"language" form:"language" validate:"required"`
	GMCNumber                 string   `json:"gmc_number" form:"gmc_number" validate:"required"`
	Specialization            *string  `json:"specialization" form:"specialization" validate:"required"`
	SubSpecialization         *string  `json:"sub_specialization" form:"sub_specialization" validate:"required"`
	LastEducation             string   `json:"last_education" form:"last_education" validate:"required"`
	YearGraduated             string   `json:"year_graduated" form:"year_graduated" validate:"required"`
	ProfessionalCertification []string `json:"professional_certification" form:"professional_certification" validate:"required"`
}
