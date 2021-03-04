package models

type Task struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"userId"`
	CreatedDate string `json:"createdDate"`
	Status      string `json:"status"`
	Assigner    string `json:"assigner"`
	Assignee    string `json:"assignee"`
	DueDate     string `json:"dueDate"`
}
