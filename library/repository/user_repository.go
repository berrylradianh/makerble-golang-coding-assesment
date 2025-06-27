/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

 
package repository

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(model *model.User, tx *gorm.DB) (*model.User, error)
	Update(model *model.User, tx *gorm.DB) error
	Delete(model *model.User, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.User, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.User, error)
	Count(criteria map[string]interface{}) int
}
