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

type EmergencyContactRepository interface {
	Create(model *model.EmergencyContact, tx *gorm.DB) (*model.EmergencyContact, error)
	Update(model *model.EmergencyContact, tx *gorm.DB) error
	Delete(model *model.EmergencyContact, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.EmergencyContact, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.EmergencyContact, error)
	Count(criteria map[string]interface{}) int
	CreateIndex(model *model.EmergencyContact) error
}
