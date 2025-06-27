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
)

// SanitizePhone make sure phone number starts with 62 and remove space, dots, and dashes
// Returns phone number and error.
func SanitizePhone(phone string) (string, error) {
	if len(phone) < 1 {
		return "", fmt.Errorf("Phone number cannot be empty")
	}
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, ".", "")
	if phone[0:1] == "0" {
		phone = "62" + phone[1:]
	} else if phone[0:1] == "8" {
		phone = "62" + phone
	} else if phone[0:3] == "+62" || phone[0:3] == " 62" {
		phone = phone[1:]
	}
	if phone[0:2] != "62" {
		return "", fmt.Errorf("Invalid phone %v", phone)
	}
	return phone, nil
}
