package helper

import (
	"encoding/json"
	"web/models"
)

func ConvertMapToUsersStruct(userMap interface{}) models.Users {
	var user models.Users
	jsonData, _ := json.Marshal(userMap)
	json.Unmarshal(jsonData, &user)

	return user
}
