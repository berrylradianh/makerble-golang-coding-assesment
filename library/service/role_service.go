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

type roleService struct {
	db     *gorm.DB
}

func NewRoleService(db *gorm.DB) repository.RoleRepository {
	return &roleService{db}
}

func (srv *roleService) FindOneBy(criteria map[string]interface{}) (*model.Role, error) {
	m := new(model.Role)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *roleService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.Role, error) {
	var data []*model.Role

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (srv *roleService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.Role{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}

func (srv *roleService) Create(model *model.Role, tx *gorm.DB) (*model.Role, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *roleService) Update(model *model.Role, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

func (srv *roleService) Delete(model *model.Role, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}
