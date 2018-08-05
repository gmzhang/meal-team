package model

import "time"

type MealTeamBroadcast struct {
	ID          int       `db:"id" json:"id"`
	TeamId      int       `db:"team_id" json:"team_id"`
	IsBroadcast int       `db:"is_broadcast" json:"is_broadcast"`
	StartAt     time.Time `db:"start_at" json:"start_at"`
	EndAt       time.Time `db:"end_at" json:"end_at"`
}
