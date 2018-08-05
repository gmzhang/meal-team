package repository

import "github.com/gmzhang/meal-team/model"

type MealTeamRepository interface {
	FindALLMealTeamName() (names []string, err error)
	CreateMealTeam(name, openid, time string) (mealTeamId int, err error)
	FindMealTeamById(id int) (mealTeam *model.MealTeam, err error)
	FindMealTeamMemberByMealTeamId(mealTeamId int) (members []model.MealTeamMember, err error)
	FindMealTeamRestaurantById(id int) (restaurant model.MealTeamRestaurantLib, err error)
	FindMealTeamRestaurantByMealTeamId(mealTeamId int) (restaurants []model.MealTeamRestaurantLib, err error)
	UpdateMealTeamMemberRestaurantByMealTeamId(mealTeamId, restaurantId int) (err error)
	CreateMealTeamBroadcast(mealTeamId,isBroadcast int) (err error)
	FindMealTeamManagerRunningByMealTeamId(mealTeamId int)(running []model.MealTeamManagerRunning, err error)

	UpdateMealTeamMemberIsManager(memberId ,isManager int) (err error)
	CreateMealTeamMember(mealTeamId int, openid,nick,avatar string, isManager int) (err error)

	CreateMealTeamNotify(mealTeamId int, openid string, result int) (err error)

	GetMealTeamIdsByOpenid(openid string)(mealTeamIds []int, err error)
}
