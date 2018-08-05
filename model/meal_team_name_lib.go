package model

import "time"

type MealTeamNameLib struct {
	ID       int       `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	CreateAt time.Time `db:"create_at" json:"create_at"`
}
