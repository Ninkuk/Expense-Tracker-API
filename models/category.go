package models

type Categories struct {
	Category []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"categories"`
}
