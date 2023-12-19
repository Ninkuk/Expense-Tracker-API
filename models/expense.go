package models

type Expenses struct {
	Expense []struct {
		ID          string `json:"id"`
		Amount      string `json:"amount"`
		Date        string `json:"date"`
		Description string `json:"description"`
		CategoryID  string `json:"category"`
		UserID      string `json:"user"`
	} `json:"users"`
}
