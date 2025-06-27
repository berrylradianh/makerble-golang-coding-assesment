/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package service

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/model"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/repository"

	"gorm.io/gorm"
)

type medicalRecordService struct {
	db *gorm.DB
}

func NewMedicalRecordService(db *gorm.DB) repository.MedicalRecordRepository {
	return &medicalRecordService{db}
}

func (srv *medicalRecordService) FindOneBy(criteria map[string]interface{}) (*model.MedicalRecord, error) {
	m := new(model.MedicalRecord)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *medicalRecordService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.MedicalRecord, error) {
	var data []*model.MedicalRecord

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (srv *medicalRecordService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.MedicalRecord{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}

func (srv *medicalRecordService) Create(model *model.MedicalRecord, tx *gorm.DB) (*model.MedicalRecord, error) {
	db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *medicalRecordService) Update(model *model.MedicalRecord, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

func (srv *medicalRecordService) Delete(model *model.MedicalRecord, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}
