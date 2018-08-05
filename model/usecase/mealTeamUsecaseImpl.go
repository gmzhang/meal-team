package usecase

import (
	"github.com/gmzhang/meal-team/model/repository"
	"github.com/gmzhang/meal-team/model"
	"time"
)

type mealTeamUsecase struct {
	repo repository.MealTeamRepository
}

func NewMealTeamUsecase(repo repository.MealTeamRepository) MealTeamUsecase {
	return &mealTeamUsecase{repo: repo}
}

func (m *mealTeamUsecase) GetALLMealTeamName() (names []string, err error) {
	names, err = m.repo.FindALLMealTeamName()
	if err != nil {
		return names, model.ErrGetALLMealTeamName
	}
	return
}

func (m *mealTeamUsecase) Create(name, openid string) (mealTeam *model.MealTeam, err error) {
	if name == "" || openid == "" {
		return nil, model.ErrInvalidParam
	}

	mealTeam = &model.MealTeam{}
	id, err := m.repo.CreateMealTeam(name, openid, time.Now().Format(time.RFC3339))

	if err != nil {
		return nil, model.ErrCreateMealTeam
	}
	mealTeam.ID = id
	return
}

func (m *mealTeamUsecase) GetMealTeam(mealTeamId int) (mealTeam *model.MealTeam, err error) {
	if mealTeamId <= 0 {
		return nil, model.ErrInvalidParam
	}
	mealTeam, err = m.repo.FindMealTeamById(mealTeamId)
	if err != nil {
		return nil, model.ErrGetMealTeam
	}
	return
}

func (m *mealTeamUsecase) GetMealTeamMember(mealTeamId int) (members []model.MealTeamMember, err error) {
	if mealTeamId <= 0 {
		return nil, model.ErrInvalidParam
	}
	members, err = m.repo.FindMealTeamMemberByMealTeamId(mealTeamId)
	if err != nil {
		return nil, model.ErrGetMealTeamMember
	}

	for k, v := range members {

		restaurant, err := m.repo.FindMealTeamRestaurantById(v.RestaurantId)
		if err != nil {
			return nil, model.ErrGetMealTeamRestaurant
		}
		members[k].Restaurant = restaurant
	}

	return
}

func (m *mealTeamUsecase) GetMealTeamALLRestaurant(mealTeamId int) (restaurants []model.MealTeamRestaurantLib, err error) {
	if mealTeamId <= 0 {
		return nil, model.ErrInvalidParam
	}

	restaurants, err = m.repo.FindMealTeamRestaurantByMealTeamId(mealTeamId)
	if err != nil {
		return nil, model.ErrGetMealTeamALLRestaurant
	}
	return
}

func (m *mealTeamUsecase) UpMealTeamRestaurantId(mealTeamId, restaurantId int) (err error) {
	if mealTeamId <= 0 || restaurantId <= 0 {
		return model.ErrInvalidParam
	}
	err = m.repo.UpdateMealTeamMemberRestaurantByMealTeamId(mealTeamId, restaurantId)
	if err != nil {
		return model.ErrUpMealTeamRestaurantId
	}
	return
}

func (m *mealTeamUsecase) BroadcastMealTeamMember(mealTeamId int) (err error) {
	if mealTeamId <= 0 {
		return model.ErrInvalidParam
	}
	err = m.repo.CreateMealTeamBroadcast(mealTeamId, 0)
	//这里创建通知任务，另外走协程去真正执行通知任务
	if err != nil {
		return model.ErrCreateMealTeamBroadcast
	}
	return
}

func (m *mealTeamUsecase) TurnOverMealTeamer(mealTeamId int) (err error) {
	//按照加入时间正序，依次移交团长
	members, err := m.repo.FindMealTeamMemberByMealTeamId(mealTeamId)
	if err != nil {
		return model.ErrGetMealTeamMember
	}

	nextTeamerKey := 0
	currentTeamKey := 0
	for k, v := range members {
		if v.IsManager != 1 {
			continue
		}
		currentTeamKey = k
		nextTeamerKey = k + 1
	}

	if (nextTeamerKey + 1) == len(members) {
		nextTeamerKey = 0
	}

	//TODO::应该是事物
	err = m.repo.UpdateMealTeamMemberIsManager(members[currentTeamKey].ID, 0)
	if err != nil {
		return model.ErrUpdateMealTeamMemberIsManager
	}
	err = m.repo.UpdateMealTeamMemberIsManager(members[nextTeamerKey].ID, 1)
	if err != nil {
		return model.ErrUpdateMealTeamMemberIsManager
	}
	return
}

func (m *mealTeamUsecase) JoinMealTeam(mealTeamId int, openid, nick, avatar string) (err error) {
	if mealTeamId <= 0 || openid == "" || nick == "" || avatar == "" {
		return model.ErrInvalidParam
	}
	err = m.repo.CreateMealTeamMember(mealTeamId, openid, nick, avatar, 0)
	if err != nil {
		return model.ErrCreateMealTeamMember
	}
	return
}

func (m *mealTeamUsecase) NotifyMealTeamer(mealTeamId int, openid string) (err error) {
	if mealTeamId <= 0 || openid == "" {
		return model.ErrInvalidParam
	}
	//TODO::调用模板消息通知团长
	err = m.repo.CreateMealTeamNotify(mealTeamId, openid, 1)
	if err != nil {
		return model.ErrCreateMealTeamNotify
	}
	return
}

func (m *mealTeamUsecase) GetMealTeamManagerRunning(mealTeamId int) (running []model.MealTeamManagerRunning, err error) {
	if mealTeamId <= 0 {
		return nil, model.ErrInvalidParam
	}
	running, err = m.repo.FindMealTeamManagerRunningByMealTeamId(mealTeamId)
	if err != nil {
		return nil, model.ErrGetMealTeamManagerRunning
	}

	for k, v := range running {
		restaurant, err := m.repo.FindMealTeamRestaurantById(v.RestaurantId)
		if err != nil {
			return nil, model.ErrGetMealTeamRestaurant
		}
		running[k].Restaurant = restaurant
	}
	return
}
