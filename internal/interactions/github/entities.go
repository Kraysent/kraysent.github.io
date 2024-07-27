package github

type Repository struct {
	Name          string
	Author        string
	Description   string
	StarNumber    int
	WatcherNumber int
	ForksNumber   int
	Website       string
	Tags          []string
}
