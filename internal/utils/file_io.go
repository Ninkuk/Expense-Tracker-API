package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ninkuk/expense-tracker-api/models"
)

func GetCategories() models.Categories {
	bytes, err := os.ReadFile("./data/categories.json")

	if err != nil {
		fmt.Println("Error loading file: ", err)
	}

	var cat models.Categories
	err = json.Unmarshal(bytes, &cat)

	if err != nil {
		fmt.Println("Error parsing file: ", err)
	}

	return cat
}

func GetUsers() models.Users {
	bytes, err := os.ReadFile("./data/users.json")

	if err != nil {
		fmt.Println("Error loading file: ", err)
	}

	var users models.Users
	err = json.Unmarshal(bytes, &users)

	if err != nil {
		fmt.Println("Error parsing file: ", err)
	}

	return users
}

func SaveUsers(users models.Users) bool {
	bytes, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	err = os.WriteFile("./data/users.json", bytes, 0666)

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	return true
}
