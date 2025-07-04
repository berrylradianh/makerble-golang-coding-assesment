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
	"time"

	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/middleware/auth"
	"github.com/form3tech-oss/jwt-go"
)

func (uc *usecase) Logout(req request.LogoutRequest) error {
	if auth.IsTokenBlacklisted(req.Token) {
		return helper.ErrTokenExpired
	}

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return auth.GetSigningKey(), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token: %v", err)
	}

	var expiry time.Time
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expiry = time.Unix(int64(exp), 0)
		} else {
			return helper.ErrTokenExpired
		}
	} else {
		return errors.New("invalid token claims")
	}

	if time.Now().After(expiry) {
		return helper.ErrTokenExpired
	}

	auth.AddToBlacklist(req.Token, expiry)
	return nil
}
