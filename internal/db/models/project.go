package models

type Project struct {
	Id          int
	Name        string
	Description string
	AvatarUrl   string `db:"avatar_url"`
	Type        string
	Archived    bool
}
