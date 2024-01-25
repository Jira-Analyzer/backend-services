package dto

type JiraProject struct {
	Name        string `json:"name"`
	Link        string `json:"self"`
	Id          string `json:"id"`
	Description string `json:"description"`
	Type        string `json:"projectTypeKey"`
	Archived    bool   `json:"archived"`
	AvatarUrls  struct {
		Url string `json:"16x16"`
	} `json:"avatarUrls"`
}

type ProjectsResponse struct {
	Projects []Project `json:"projects"`
	PageInfo PageInfo  `json:"pageInfo"`
}

type Project struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type PageInfo struct {
	PageCount     int `json:"pageCount"`
	CurrentPage   int `json:"currentPage"`
	ProjectsCount int `json:"projectsCount"`
}
