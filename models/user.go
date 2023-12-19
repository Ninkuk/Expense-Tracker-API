package models

type Users struct {
	User []struct {
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		BirthDate string `json:"birthDate"`
	} `json:"users"`
}
