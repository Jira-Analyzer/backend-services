package domain

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `db:"avatar_url" json:"avatar_url"`
	Type        string `json:"type"`
	Archived    bool   `json:"archived"`
}
