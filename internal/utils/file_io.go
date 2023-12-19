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
