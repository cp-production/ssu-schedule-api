package model

type Departments struct {
	Id        int    `json:"id"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	Url       string `json:"url"`
}
