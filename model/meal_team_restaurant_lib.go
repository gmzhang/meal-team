package model

import "time"

type MealTeamRestaurantLib struct {
	ID           int       `db:"id" json:"id"`
	TeamId       int       `db:"team_id" json:"-"`
	Name         string    `db:"name" json:"name"`
	CreateOpenid string    `db:"create_openid" json:"-"`
	CreateAt     time.Time `db:"create_at" json:"create_at"`
}
