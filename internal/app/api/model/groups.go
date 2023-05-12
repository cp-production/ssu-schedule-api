package model

type Groups struct {
	Id       int `json:"id"`
	EdForm   string `json:"education_form"`
	GroupNum string `json:"group_num"`
	DepId    int `json:"department_id"`
}
