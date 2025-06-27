/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */


package service

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/repository"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/model"

	"gorm.io/gorm"
)

type emergencyContactService struct {
	db     *gorm.DB
}

func NewEmergencyContactService(db *gorm.DB) repository.EmergencyContactRepository {
	return &emergencyContactService{db}
}

func (srv *emergencyContactService) FindOneBy(criteria map[string]interface{}) (*model.EmergencyContact, error) {
	m := new(model.EmergencyContact)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *emergencyContactService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.EmergencyContact, error) {
	var data []*model.EmergencyContact

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (srv *emergencyContactService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.EmergencyContact{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}

func (srv *emergencyContactService) Create(model *model.EmergencyContact, tx *gorm.DB) (*model.EmergencyContact, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *emergencyContactService) Update(model *model.EmergencyContact, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

func (srv *emergencyContactService) Delete(model *model.EmergencyContact, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}
