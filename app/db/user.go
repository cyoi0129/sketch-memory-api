package db

import (
	"fmt"
	"sketch/graph/model"
)

func FetchUserList() ([]*model.User, error) {
	var users []*model.User
	rows, err := DB.Query("SELECT id, name, password FROM \"sketch_users\"")
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	for rows.Next() {
		var user model.User
		rows.Scan(&user.ID, &user.Name, &user.Password)
		users = append(users, &user)
	}

	return users, nil
}

func FetchUserById(user_id string) (*model.User, error) {
	var user model.User
	row := DB.QueryRow("SELECT id, name, password FROM \"sketch_users\" WHERE id = $1", user_id)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		fmt.Println(err)
		return &user, err
	}
	return &user, nil
}
