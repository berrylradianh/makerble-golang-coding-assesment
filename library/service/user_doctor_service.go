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

type userDoctorService struct {
	db *gorm.DB
}

func NewUserDoctorService(db *gorm.DB) repository.UserDoctorRepository {
	return &userDoctorService{db}
}

func (srv *userDoctorService) FindOneBy(criteria map[string]interface{}) (*model.UserDoctor, error) {
	m := new(model.UserDoctor)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *userDoctorService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.UserDoctor, error) {
	var data []*model.UserDoctor

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (srv *userDoctorService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.UserDoctor{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}

func (srv *userDoctorService) Create(model *model.UserDoctor, tx *gorm.DB) (*model.UserDoctor, error) {
	db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *userDoctorService) Update(model *model.UserDoctor, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

func (srv *userDoctorService) Delete(model *model.UserDoctor, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}
