/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package user

import (
	"github.com/berrylradianh/makerble-golang-coding-assesment/app/request"
	"github.com/berrylradianh/makerble-golang-coding-assesment/app/response"
)

type Usecase interface {
	AddUser(req request.AddUserRequest) (*response.AddUserResponse, error)
}
