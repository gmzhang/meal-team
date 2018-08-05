package usecase

import "github.com/gmzhang/meal-team/model"

type MealTeamUsecase interface {
	GetALLMealTeamName() (names []string, err error)

	Create(name, openid string) (mealTeam *model.MealTeam, err error)
	GetMealTeam(mealTeamId int) (mealTeam *model.MealTeam, err error)

	GetMealTeamMember(mealTeamId int) (members []model.MealTeamMember, err error)

	GetMealTeamALLRestaurant(mealTeamId int) (restaurants []model.MealTeamRestaurantLib, err error)

	UpdateMealTeamRestaurantId(mealTeamId, restaurantId int) (err error)

	BroadcastMealTeamMember(mealTeamId int) (err error)

	TurnOverMealTeamer(mealTeamId int) (err error)

	JoinMealTeam(mealTeamId int, openid,nick,avatar string) (err error)

	NotifyMealTeamer(mealTeamId int, openid string) (err error)

	GetMealTeamManagerRunning(mealTeamId int) (running []model.MealTeamManagerRunning, err error)

	GetMealTeamByMemberOpenid(openid string)(mealTeams []model.MealTeam, err error)
}
