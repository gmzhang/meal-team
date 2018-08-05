package model

import "time"

type MealTeamNotify struct {
	ID       int       `db:"id" json:"id"`
	TeamId   int       `db:"team_id" json:"team_id"`
	Openid   string    `db:"openid" json:"openid"`
	Result   string    `db:"result" json:"result"`
	CreateAt time.Time `db:"create_at" json:"create_at"`
}
