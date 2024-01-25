package domain

type Author struct {
	Id          int
	Name        string
	DisplayName string `db:"display_name"`
}
