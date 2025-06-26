/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package main

import (
	"os"

	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	"github.com/sirupsen/logrus"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	env = config.NewConfig()

	logrus.Info("success loading .env")
	logrus.Info("Starting server... \n")

	startApp()
}
