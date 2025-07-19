/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package helper

import (
	"fmt"
	"strings"
	"time"
)

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}

func GenerateEmployeeNumber(name, dateOfBirth string) (string, error) {
	const layout = "2 January 2006"
	parsedDate, err := time.Parse(layout, dateOfBirth)
	if err != nil {
		return "", fmt.Errorf("failed to parse date_of_birth: %v", err)
	}

	nameParts := strings.Split(strings.TrimSpace(name), " ")
	var initial string
	for _, part := range nameParts {
		if !strings.HasPrefix(strings.ToLower(part), "dr.") && len(part) > 0 {
			initial = strings.ToUpper(string(part[0]))
			break
		}
	}
	if initial == "" {
		return "", fmt.Errorf("could not determine initial from name: %s", name)
	}

	dateStr := parsedDate.Format("02012006")

	employeeNumber := fmt.Sprintf("EMP-%s-%s", initial, dateStr)
	return employeeNumber, nil
}
