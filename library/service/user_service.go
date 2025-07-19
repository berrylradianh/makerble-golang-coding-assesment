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

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) repository.UserRepository {
	return &userService{db}
}

func (srv *userService) FindOneBy(criteria map[string]interface{}) (*model.User, error) {
	m := new(model.User)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *userService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.User, error) {
	var data []*model.User

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (srv *userService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.User{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}

func (srv *userService) Create(model *model.User, tx *gorm.DB) (*model.User, error) {
	db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *userService) Update(model *model.User, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

func (srv *userService) Delete(model *model.User, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}
