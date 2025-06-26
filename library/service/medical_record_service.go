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

type medicalRecordService struct {
	db     *gorm.DB
	client *elastic.Client
}

func NewMedicalRecordService(db *gorm.DB, client *elastic.Client) repository.MedicalRecordRepository {
	return &medicalRecordService{db, client}
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

func (srv *medicalRecordService) CreateIndex(model *model.MedicalRecord) error {
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

func (srv *medicalRecordService) CreateOrUpdateIndex(model *model.MedicalRecord) error {
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
