package model

type Bill struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}
