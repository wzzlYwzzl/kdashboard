package user

import (
	"github.com/kubernetes/dashboard/client"
	httpdbuser "github.com/wzzlYwzzl/httpdatabase/resource/user"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//Judge user whether exist, if exist, return all the info about the user.
func JudgeUser(httpdbClient *client.HttpDBClient, user *User) (*httpdbuser.User, error) {
	b, err := httpdbClient.JudgeName(user.Name, user.Password)
	if err != nil {
		return nil, err
	}
	if b != true {
		return nil, nil
	}

	userinfo, err := httpdbClient.GetAllInfo(user.Name)
	if err != nil {
		return nil, err
	}

	return userinfo, nil
}
