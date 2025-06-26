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

type MedicalRecordRepository interface {
	Create(model *model.MedicalRecord, tx *gorm.DB) (*model.MedicalRecord, error)
	Update(model *model.MedicalRecord, tx *gorm.DB) error
	Delete(model *model.MedicalRecord, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.MedicalRecord, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.MedicalRecord, error)
	Count(criteria map[string]interface{}) int
	CreateIndex(model *model.MedicalRecord) error
}
