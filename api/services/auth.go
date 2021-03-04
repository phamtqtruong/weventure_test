package services

import (
	"net/http"
	"weventure_test/common/rest"
)

// AuthService ...
type AuthService struct {
}

// Login ...
func (t *AuthService) Login(id string, password string) rest.Response {
	var (
		res         rest.Response
		data        = make(map[string]interface{})
		userService = UsersService{}
		user        = userService.FindByID(id)
	)
	if user == nil || user.Pwd != password {
		var exception = &rest.Error{
			Code:    http.StatusUnauthorized,
			Message: "Id or password was wrong",
		}
		res = rest.Response{
			Status: 0,
			Error:  exception,
		}
		return res
	}
	data["token"] = "new_generated_token"
	res = rest.Response{
		Code:   http.StatusOK,
		Status: 1,
		Data:   data,
	}

	return res
}
