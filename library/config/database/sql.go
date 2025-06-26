/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package config

import (
	"fmt"

	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDBSQL initializes PostgreSQL database connection
func InitDBSQL(env config.Config, parentKey string) (db *gorm.DB, err error) {
	dbUser := env.GetString(parentKey + `.user`)
	dbPass := env.GetString(parentKey + `.pass`)
	dbName := env.GetString(parentKey + `.database`)
	dbHost := env.GetString(parentKey + `.address`)
	dbPort := env.GetString(parentKey + `.port`)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db, err
}
