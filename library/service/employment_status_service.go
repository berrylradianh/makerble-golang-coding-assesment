/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */


package service

import (
	"context"
	"strconv"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/repository"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/struct/model"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type employmentStatusService struct {
	db     *gorm.DB
	client *elastic.Client
}

func NewEmploymentStatusService(db *gorm.DB, client *elastic.Client) repository.EmploymentStatusRepository {
	return &employmentStatusService{db, client}
}

func (srv *employmentStatusService) FindOneBy(criteria map[string]interface{}) (*model.EmploymentStatus, error) {
	m := new(model.EmploymentStatus)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *employmentStatusService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.EmploymentStatus, error) {
	var data []*model.EmploymentStatus

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (srv *employmentStatusService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.EmploymentStatus{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}

func (srv *employmentStatusService) Create(model *model.EmploymentStatus, tx *gorm.DB) (*model.EmploymentStatus, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *employmentStatusService) Update(model *model.EmploymentStatus, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

func (srv *employmentStatusService) Delete(model *model.EmploymentStatus, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}

func (srv *employmentStatusService) CreateIndex(model *model.EmploymentStatus) error {
	exists, err := srv.client.IndexExists(model.TableName()).Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		_, err := srv.client.CreateIndex(model.TableName()).Do(context.Background())
		if err != nil {
			return err
		}
	}

	_, err = srv.client.Index().
		Index(model.TableName()).
		Id(strconv.Itoa(model.ID)).
		BodyJson(&model).
		Do(context.Background())
	return err
}

func (srv *employmentStatusService) CreateOrUpdateIndex(model *model.EmploymentStatus) error {
	ctx := context.Background()
	// try to find index
	exists, err := srv.client.IndexExists(model.TableName()).Do(ctx)
	if err != nil {
		return err
	}

	// index not exist so create first
	if !exists {
		_, err = srv.client.CreateIndex(model.TableName()).Do(ctx)
		if err != nil {
			return err
		}
	}

	// try to delete old entry
	_, err = srv.client.Delete().
		Index(model.TableName()).
		Id(strconv.Itoa(model.ID)).
		Refresh("true").
		Do(ctx)

	// if error is not null and is not data not found
	if err != nil && !elastic.IsNotFound(err) {
		return err
	}

	// create entry
	_, err = srv.client.Index().
		Index(model.TableName()).
		Id(strconv.Itoa(model.ID)).
		BodyJson(&model).
		Do(context.Background())
	return err
}
