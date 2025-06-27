/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package auth

import (
	"errors"
	"fmt"

	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/middleware/auth"
	"gorm.io/gorm"
)

func (uc *usecase) Login(req request.LoginRequest) (*auth.TokenResponse, error) {
	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	if user == nil {
		return nil, helper.ErrInvalidEmailOrPassword
	}

	role, err := uc.roleRepo.FindOneBy(map[string]interface{}{
		"id": user.RoleID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find role: %v", err)
	}

	if !helper.ComparePasswords(user.Password, req.Password) {
		return nil, helper.ErrInvalidEmailOrPassword
	}

	tx := uc.db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Commit().Error; err != nil {
		err = fmt.Errorf("failed to commit transaction: %w", err)
	}

	tokenStruct := auth.TokenStructure{
		UserID: user.ID,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   role.Slug,
	}

	conf := config.NewConfig()
	cswAuth := auth.NewClinicPortalAuth([]byte(conf.GetString("app.signature")))
	res, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return res, nil
}
