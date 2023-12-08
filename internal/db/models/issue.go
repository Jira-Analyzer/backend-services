package models

import "time"

type Issue struct {
	Id          int
	ProjectId   int `db:"project_id"`
	AuthorId    int `db:"author_id"`
	ReporterId  int `db:"reporter_id"`
	Key         string
	Summary     string
	Type        string
	Priority    string
	Status      string
	CreatedTime time.Time     `db:"created_time"`
	ClosedTime  time.Time     `db:"closed_time"`
	UpdatedTime time.Time     `db:"updated_time"`
	TimeSpent   time.Duration `db:"time_spent"`
}
