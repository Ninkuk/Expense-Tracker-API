package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ninkuk/expense-tracker-api/models"
)

func GetDocTXT() []byte {
	bytes, err := os.ReadFile("./data/docs.txt")

	if err != nil {
		fmt.Println("Error loading file: ", err)
		return nil
	}

	return bytes
}

func GetDocHTML() []byte {
	bytes, err := os.ReadFile("./data/index.html")

	if err != nil {
		fmt.Println("Error loading file: ", err)
		return nil
	}

	return bytes
}

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

func GetExpenses() models.Expenses {
	bytes, err := os.ReadFile("./data/expenses.json")

	if err != nil {
		fmt.Println("Error loading file: ", err)
	}

	var expenses models.Expenses
	err = json.Unmarshal(bytes, &expenses)

	if err != nil {
		fmt.Println("Error parsing file: ", err)
	}

	return expenses
}

func SaveExpenses(expenses models.Expenses) bool {
	bytes, err := json.MarshalIndent(expenses, "", "\t")

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	err = os.WriteFile("./data/expenses.json", bytes, 0666)

	if err != nil {
		fmt.Println("Error saving file: ", err)
		return false
	}

	return true
}
