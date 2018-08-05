package model

import "time"

type MealTeam struct {
	ID       int       `db:"id" json:"id"`
	Openid   string    `db:"openid" json:"openid"`
	Name     string    `db:"name" json:"name"`
	CreateAt time.Time `db:"create_at" json:"create_at"`
}
