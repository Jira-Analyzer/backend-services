package jira

/*type FetchProjectResponseDTO struct {
	Projects []*domain.Project `json:"projects"`
}

type FetchIssueResponseDTO struct {
	Issues []*domain.Issue `json:"issues"`
}*/

type IssuesList struct {
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}

type Issue struct {
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`
	Type    struct {
		Name string `json:"name"`
	} `json:"issuetype"`
	Status struct {
		Name string `json:"name"`
	} `json:"status"`
	Priority struct {
		Name string `json:"name"`
	} `json:"priority"`
	Creator struct {
		Name string `json:"name"`
	} `json:"creator"`
	Project struct {
		Name string `json:"name"`
	} `json:"project"`
	Description  string `json:"description"`
	AssigneeName struct {
		Name string `json:"name"`
	} `json:"assignee"`
	CreatedTime string `json:"created"`
	UpdatedTime string `json:"updated"`
	ClosedTime  string `json:"resolutiondate"`
}

type IssueStatusChange struct {
	Changelog struct {
		Histories []History `json:"histories"`
	} `json:"changelog"`
}

type History struct {
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	CreatedTime string `json:"created"`
	Items       []Item
}

type Item struct {
	Field      string `json:"field"`
	FromString string `json:"fromString"`
	ToString   string `json:"toString"`
}

type JiraProject struct {
	Name string `json:"name"`
	Link string `json:"self"`
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
