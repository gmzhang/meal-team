package model

import (
	"errors"
)

const ErrCodeOK = 1000
const ErrCodeSegmentBase = 2500

var (
	ErrInvalidParam                  = errors.New("invalid param")
	ErrGetALLMealTeamName            = errors.New("get all meal team name error")
	ErrCreateMealTeam                = errors.New("create meal team error")
	ErrGetMealTeam                   = errors.New("get meal team error")
	ErrGetMealTeamMember             = errors.New("get meal team member error")
	ErrGetMealTeamRestaurant         = errors.New("get meal team restaurant error")
	ErrGetMealTeamALLRestaurant      = errors.New("get meal team all restaurant error")
	ErrUpMealTeamRestaurantId        = errors.New("update meal team restaurant error")
	ErrCreateMealTeamBroadcast       = errors.New("create meal team broadcast error")
	ErrGetMealTeamManagerRunning     = errors.New("get meal team manager running error")
	ErrUpdateMealTeamMemberIsManager = errors.New("update meal team member is manager error")
	ErrCreateMealTeamMember          = errors.New("create meal team member error")
	ErrCreateMealTeamNotify          = errors.New("create meal team notify error")
	ErrGetMealTeamIdsByOpenid         = errors.New("get meal team ids by openid error")
)

func GetErrorCode(err error) int32 {
	switch err {
	case ErrInvalidParam:
		return ErrCodeSegmentBase + 1
	case ErrGetALLMealTeamName:
		return ErrCodeSegmentBase + 2
	case ErrCreateMealTeam:
		return ErrCodeSegmentBase + 3
	case ErrGetMealTeam:
		return ErrCodeSegmentBase + 4
	case ErrGetMealTeamMember:
		return ErrCodeSegmentBase + 5
	case ErrGetMealTeamRestaurant:
		return ErrCodeSegmentBase + 6
	case ErrGetMealTeamALLRestaurant:
		return ErrCodeSegmentBase + 7
	case ErrUpMealTeamRestaurantId:
		return ErrCodeSegmentBase + 8
	case ErrCreateMealTeamBroadcast:
		return ErrCodeSegmentBase + 9
	case ErrGetMealTeamManagerRunning:
		return ErrCodeSegmentBase + 10
	case ErrUpdateMealTeamMemberIsManager:
		return ErrCodeSegmentBase + 11
	case ErrCreateMealTeamMember:
		return ErrCodeSegmentBase + 12
	case ErrCreateMealTeamNotify:
		return ErrCodeSegmentBase + 13
	case ErrGetMealTeamIdsByOpenid:
		return ErrCodeSegmentBase + 14

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
