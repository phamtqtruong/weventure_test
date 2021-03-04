package services

import (
	"weventure_test/api/models"
	"weventure_test/common/db"
)

// UsersService ...
type UsersService struct {
}

// GetAll ...
func (s *UsersService) GetAll() []*models.User {
	var (
		uID      string
		uPwd     string
		uMaxTodo int
		sql      = "SELECT * FROM users"
		rows, _  = db.DB().Query(sql)
		data     = make([]*models.User, 0)
	)
	for rows.Next() {
		err := rows.Scan(&uID, &uMaxTodo, &uPwd)
		if err != nil {
			return data
		}
		data = append(data, &models.User{ID: uID, MaxTodo: uMaxTodo, Pwd: uPwd})
	}

	return data
}

// FindByID ...
func (s *UsersService) FindByID(id string) *models.User {
	var (
		uID      string
		uPwd     string
		uMaxTodo int
		sql      = "SELECT * FROM users id=$1 AND password=$2 LIMIT 1"
		rows     = db.DB().QueryRow(sql, id)
	)
	if err := rows.Scan(&uID, &uMaxTodo, &uPwd); err != nil {
		return nil
	}

	return &models.User{ID: uID, MaxTodo: uMaxTodo, Pwd: uPwd}
}

func (s *UsersService) RefreshMaxTodo(id string) {
	var (
		sql = "UPDATE `users` SET max_todo=$2 WHERE id=$1"
	)
	err := db.DB().QueryRow(sql, id, 4)
	if err != nil {
		return

	}
}
