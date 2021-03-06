package model

import "time"

type MealTeamManagerRunning struct {
	ID           int       `db:"id" json:"id"`
	TeamId       int       `db:"team_id" json:"team_id"`
	RestaurantId int       `db:"restaurant_id" json:"restaurant_id"`
	Openid       string    `db:"openid" json:"openid"`
	Nick         string    `db:"nick" json:"nick"`
	Avatar       string    `db:"avatar" json:"avatar"`
	Up           int       `db:"up" json:"up"`
	Down         int       `db:"down" json:"down"`
	CreateAt     time.Time `db:"create_at" json:"create_at"`

	Restaurant MealTeamRestaurantLib `json:"restaurant"`
}
