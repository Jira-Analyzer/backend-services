package domain

import "time"

type Issue struct {
	Id          int           `json:"id"`
	ProjectId   int           `db:"project_id"`
	Author      string        `json:"author"`
	Reporter    string        `json:"reporter"`
	Key         string        `json:"key"`
	Summary     string        `json:"summary"`
	Type        string        `json:"type"`
	Priority    string        `json:"priority"`
	Status      string        `json:"status"`
	CreatedTime time.Time     `db:"created_time" json:"created_time"`
	ClosedTime  time.Time     `db:"closed_time" json:"closed_time"`
	UpdatedTime time.Time     `db:"updated_time" json:"updated_time"`
	TimeSpent   time.Duration `db:"time_spent"  json:"time_spent" swaggertype:"string"`
}
