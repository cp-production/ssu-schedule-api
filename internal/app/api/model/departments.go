package model

type Departments struct {
	Id        int    `json:"id"`
	FullName  string `json:"full_name"`
	ShortName string `json:"short_name"`
	Url       string `json:"url"`
}
