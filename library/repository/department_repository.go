/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

 
package repository

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/struct/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(model *model.Department, tx *gorm.DB) (*model.Department, error)
	Update(model *model.Department, tx *gorm.DB) error
	Delete(model *model.Department, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Department, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Department, error)
	Count(criteria map[string]interface{}) int
	CreateIndex(model *model.Department) error
}
