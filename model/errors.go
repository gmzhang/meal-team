package model

import (
	"errors"
)

const ErrCodeOK = 1000
const ErrCodeSegmentBase = 2500

var (
	ErrInvalidParam       = errors.New("invalid param")
	ErrGetALLMealTeamName = errors.New("get all meal team name error")
	ErrCreateMealTeam = errors.New("create meal team error")
	ErrGetMealTeam = errors.New("get meal team error")
	ErrGetMealTeamMember = errors.New("get meal team member error")
	ErrGetMealTeamRestaurant = errors.New("get meal team restaurant error")
	ErrGetMealTeamALLRestaurant = errors.New("get meal team all restaurant error")
	ErrUpMealTeamRestaurantId = errors.New("update meal team restaurant error")
	ErrCreateMealTeamBroadcast = errors.New("create meal team broadcast error")
	ErrGetMealTeamManagerRunning = errors.New("get meal team manager running error")
	ErrUpdateMealTeamMemberIsManager = errors.New("update meal team member is manager error")
	ErrCreateMealTeamMember = errors.New("create meal team member error")
	ErrCreateMealTeamNotify = errors.New("create meal team notify error")

)

func GetErrorCode(err error) int32 {
	switch err {
	case ErrInvalidParam:
		return ErrCodeSegmentBase + 1
	case ErrGetALLMealTeamName:
		return ErrCodeSegmentBase + 2

	case nil:
		return ErrCodeOK
	default:
		return 0
	}
}

func GetErrorMap(err error) map[string]interface{} {
	var msg = "OK"
	if err != nil {
		msg = err.Error()
	}

	return map[string]interface{}{
		"errcode": GetErrorCode(err),
		"msg":     msg,
	}
}
