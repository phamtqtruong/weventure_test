package services

import (
	"fmt"
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
		sql      = "SELECT * FROM `users` WHERE id=?"
		rows     = db.DB().QueryRow(sql, id)
	)
	if err := rows.Scan(&uID, &uPwd, &uMaxTodo); err != nil {
		fmt.Println("UsersService.FindById", err.Error())
		return nil
	}

	return &models.User{ID: uID, MaxTodo: uMaxTodo, Pwd: uPwd}
}

func (s *UsersService) RefreshMaxTodo(id string) {
	var (
		sql  = "UPDATE `users` SET max_todo=? WHERE id=?"
		user = s.FindByID(id)
	)

	if user == nil {
		fmt.Println("UsersService.RefreshMaxTodo", user)
		return
	}

	var stmt, _ = db.DB().Prepare(sql)
	var _, err = stmt.Exec(user.MaxTodo-1, id)
	if err != nil {
		return
	}
}
