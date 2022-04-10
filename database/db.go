package database

import (
	"encoding/json"
	"os"
	"strings"
)

type User struct {
	Username     string        `json:"username"`
	Balance      int64         `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Amount    int64  `json:"amount"`
	Type      string `json:"type"`
	Narration string `json:"narration"`
}

func getUsers() ([]User, error) {
	data, err := os.ReadFile("database/db.json")

	var users []User

	if err == nil {
		json.Unmarshal(data, &users)
	}

	return users, err
}

func updateDB(data []User) {
	bytes, err := json.Marshal(data)

	if err == nil {
		os.WriteFile("database/db.json", bytes, 0644)
	} else {
		panic(err.(any))
	}
}

func FindUser(username string) (*User, error) {
	users, err := getUsers()

	if err == nil {
		for _, user := range users {
			if strings.EqualFold(user.Username, username) {
				return &user, nil
			}
		}
	}

	return nil, err
}

func FindOrCreateUser(username string) (*User, error) {
	user, err := FindUser(username)

	if user == nil {
		newUser := User{
			Username:     strings.ToLower(username),
			Balance:      0,
			Transactions: []Transaction{},
		}
		users, err := getUsers()

		if err == nil {
			users = append(users, newUser)
			updateDB(users)
		}
		return &newUser, err
	}

	return user, err
}

func UpdateUser(user *User) {
	users, err := getUsers()
	if err == nil {
		for index := 0; index < len(users); index++ {
			users[index] = *user
		}
	}
	updateDB(users)
}
