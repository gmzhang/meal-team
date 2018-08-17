package repository

import (
	"github.com/jmoiron/sqlx"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/gmzhang/meal-team/model"
	"time"
)

type mealTeamRepository struct {
	coon *sqlx.DB
}

func NewMealTeamRepository(coon *sqlx.DB) MealTeamRepository {
	repo := mealTeamRepository{
		coon: coon,
	}
	return &repo
}

func (m *mealTeamRepository) FindALLMealTeamName() (names []string, err error) {

	names = []string{}

	sqlName := "select name from meal_team_name_lib"

	err = m.coon.Select(&names, sqlName)

	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		logrus.WithError(err).Error("select name from meal_team_name_lib error")
	}
	return
}

func (m *mealTeamRepository) CreateMealTeam(name, openid, time string) (mealTeamId int, err error) {
	sqlCreate := "INSERT INTO meal_team(openid,name,create_at) VALUE (?,?,?)"

	result, err := m.coon.Exec(sqlCreate, openid, name, time)

	if err != nil {
		logrus.WithError(err).WithField("name", name).WithField("openid", openid).Error("insert meal_team sql exec error")
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		logrus.WithError(err).WithField("name", name).WithField("openid", openid).WithField("time", time).Error("create meal team get lastInsertId error")
		return 0, err
	}
	return int(id), nil
}

func (m *mealTeamRepository) FindMealTeamById(id int) (mealTeam *model.MealTeam, err error) {
	sqlStr := "select id, name,openid,create_at from meal_team where id=?"
	mealTeam = &model.MealTeam{}
	err = m.coon.Get(mealTeam, sqlStr, id)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("select meal team by id error")
	}
	return
}

func (m *mealTeamRepository) FindMealTeamMemberByMealTeamId(mealTeamId int) (members []model.MealTeamMember, err error) {
	members = []model.MealTeamMember{}
	sqlStr := "select id,team_id,restaurant_id,openid,nick,avatar,is_manager,up,down,create_at FROM meal_team_member WHERE team_id=? order by id asc"

	err = m.coon.Select(&members, sqlStr, mealTeamId)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).Error("select meal team member by team_id error")
	}
	return

}

func (m *mealTeamRepository) FindMealTeamRestaurantById(id int) (restaurant model.MealTeamRestaurantLib, err error) {
	restaurant = model.MealTeamRestaurantLib{}
	sqlStr := "select id,name,create_at from meal_team_restaurant_lib where id=?"

	err = m.coon.Get(&restaurant, sqlStr, id)
	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("get meal team restaurant error")
	}
	return

}

func (m *mealTeamRepository) FindMealTeamRestaurantByMealTeamId(mealTeamId int) (restaurants []model.MealTeamRestaurantLib, err error) {
	restaurants = []model.MealTeamRestaurantLib{}

	sqlStr := "select id, name, create_at from meal_team_restaurant_lib where team_id=?"

	err = m.coon.Select(&restaurants, sqlStr, mealTeamId)
	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).Error("select meal team restaurant error")
	}
	return
}

func (m *mealTeamRepository) UpdateMealTeamMemberRestaurantByMealTeamId(mealTeamId, restaurantId int) (err error) {
	sqlStr := "update meal_team_member set restaurant_id = ? where team_id=?"

	result, err := m.coon.Exec(sqlStr, restaurantId, mealTeamId)
	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("restaurantId", restaurantId).Error("update meal_team_member restaurant error")
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("restaurantId", restaurantId).Error("update meal_team_member restaurant rowsAffects error")
	}
	return
}

func (m *mealTeamRepository) CreateMealTeamBroadcast(mealTeamId, isBroadcast int) (err error) {
	sqlCreate := "INSERT INTO meal_team_broadcast(team_id,is_broadcast) VALUE (?,?)"

	result, err := m.coon.Exec(sqlCreate, mealTeamId, isBroadcast)

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("isBroadcast", isBroadcast).Error("insert meal_team_broadcast sql exec error")
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("isBroadcast", isBroadcast).Error("create meal_team_broadcast get lastInsertId error")
		return err
	}
	return nil
}

func (m *mealTeamRepository) FindMealTeamManagerRunningByMealTeamId(mealTeamId int) (running []model.MealTeamManagerRunning, err error) {

	running = []model.MealTeamManagerRunning{}

	sqlStr := "select id, team_id, restaurant_id,openid,nick,avatar,up,down,create_at from meal_team_manager_running where team_id=?"

	err = m.coon.Select(&running, sqlStr, mealTeamId)
	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).Error("select meal_team_manager_running error")
	}
	return
}

func (m *mealTeamRepository) UpdateMealTeamMemberIsManager(memberId, isManager int) (err error) {

	updateSql := "update meal_team_member set is_manager = ? where id=?"

	result, err := m.coon.Exec(updateSql, isManager, memberId)
	if err != nil {
		logrus.WithError(err).WithField("id", memberId).WithField("isManager", isManager).Error("update meal_team_member is_manager sql error")
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logrus.WithError(err).WithField("id", memberId).WithField("isManager", isManager).Error("update meal_team_member is_manager rowsAffects error")
	}
	return
}

func (m *mealTeamRepository) CreateMealTeamMember(mealTeamId int, openid, nick, avatar string, isManager int) (err error) {
	sqlCreate := "INSERT INTO meal_team_member(team_id,openid,nick,avatar,is_manager,create_at) VALUE (?,?,?,?,?,?)"

	result, err := m.coon.Exec(sqlCreate, mealTeamId, openid, nick, avatar, isManager, time.Now().Format(time.RFC3339))

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("openid", openid).WithField("nick", nick).WithField("avatar", avatar).WithField("isManager", isManager).Error("insert meal_team_member sql exec error")
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("openid", openid).WithField("nick", nick).WithField("avatar", avatar).WithField("isManager", isManager).Error("insert meal_team_member get lastInsertId error")
		return err
	}
	return nil
}

func (m *mealTeamRepository) CreateMealTeamNotify(mealTeamId int, openid string, notifyResult int) (err error) {
	sqlCreate := "INSERT INTO meal_team_notify(team_id,openid,result,create_at) VALUE (?,?,?,?)"

	result, err := m.coon.Exec(sqlCreate, mealTeamId, openid, notifyResult, time.Now().Format(time.RFC3339))

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("openid", openid).WithField("notifyResult", notifyResult).Error("insert meal_team_notify sql exec error")
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		logrus.WithError(err).WithField("mealTeamId", mealTeamId).WithField("openid", openid).WithField("notifyResult", notifyResult).Error("insert meal_team_notify get lastInsertId error")
		return err
	}
	return
}

func (m *mealTeamRepository) GetMealTeamManagersByOpenid(openid string)(mealTeamManagers []model.MealTeamMember, err error) {

	sqlStr := "select id,team_id,is_manager,openid,nick,avatar,up,down FROM meal_team_member WHERE openid=? and is_manager = 1 order by id desc"

	mealTeamManagers = []model.MealTeamMember{}
	err = m.coon.Select(&mealTeamManagers, sqlStr, openid)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		logrus.WithError(err).WithField("openid", openid).Error("select team_id from meal_team_member by openid error")
	}
	return

}
