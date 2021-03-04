package services

import (
	"fmt"
	"net/http"
	"weventure_test/api/dto"
	"weventure_test/api/models"
	"weventure_test/common/db"
	"weventure_test/common/rest"
)

// TasksService ...
type TasksService struct {
}

// GetAll ...
func (t *TasksService) GetAll(assignee string, dueDate string) rest.Response {
	var (
		res          rest.Response
		tID          int
		tContent     string
		tUserID      string
		tCreatedDate string
		tAssigner    string
		tAssignee    string
		tDueDate     string
		tStatus      string
		sql          = "SELECT * FROM tasks"
		data         = make([]models.Task, 0)
	)
	if assignee != "" || dueDate != "" {
		var condition = " WHERE "
		if assignee != "" {
			var assigneeCond = "assignee='" + assignee + "'"
			if condition != " WHERE " {
				assigneeCond = " AND " + assigneeCond
			}
			condition += assigneeCond
		}
		if dueDate != "" {
			var dueDateCond = "due_date='" + dueDate + "'"
			if condition != " WHERE " {
				dueDateCond = " AND " + dueDateCond
			}
			condition += dueDateCond
		}
		sql += condition
	}
	var rows, _ = db.DB().Query(sql)
	for rows.Next() {
		err := rows.Scan(&tID, &tContent, &tUserID, &tCreatedDate, &tStatus, &tAssigner, &tAssignee, &tDueDate)
		if err != nil {
			var exception = &rest.Error{
				Code:    http.StatusUnauthorized,
				Message: "No tasks found",
			}
			res = rest.Response{
				Status: 0,
				Error:  exception,
			}
			return res
		}
		data = append(data, models.Task{ID: tID, Content: tContent, UserID: tUserID, CreatedDate: tCreatedDate, Assigner: tAssigner, Assignee: tAssignee, DueDate: tDueDate, Status: tStatus})
	}
	res = rest.Response{
		Code:   http.StatusOK,
		Status: 1,
		Data:   data,
	}

	return res
}

func (t *TasksService) New(data dto.TaskDTO) *models.Task {
	var (
		ret         models.Task
		userService = UsersService{}
		user        = userService.FindByID(data.UserID)
	)

	if user == nil || user.MaxTodo == 0 {
		return nil
	}

	var (
		sql     = "INSERT INTO `tasks` (`content`, `user_id`, `created_date`, `assigner`, `assignee`, `progress`, `due_date`) VALUES (?, ?, ?, ?, ?, ?, ?)"
		stmt, _ = db.DB().Prepare(sql)
		_, err  = stmt.Exec(data.Content, data.UserID, data.CreatedDate, data.Assigner, data.Assignee, data.Status, data.DueDate)
	)
	if err != nil {
		fmt.Println("db.DB()", data, err.Error())
		return nil
	}

	userService.RefreshMaxTodo(data.UserID)
	ret = models.Task{Content: data.Content, UserID: data.UserID, CreatedDate: data.CreatedDate, Assigner: data.Assigner, Assignee: data.Assignee, Status: data.Status, DueDate: data.DueDate}

	return &ret
}
