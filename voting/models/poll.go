package models

type Poll struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Choices     []string `json:"choice"`
}
