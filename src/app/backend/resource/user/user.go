package user

import (
	"github.com/kubernetes/dashboard/client"
	httpdbuser "github.com/wzzlYwzzl/httpdatabase/resource/user"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserCreate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Cpus     int    `json:"cpus"`
	Memory   int    `json:"memory"`
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

//delete user
func DeleteUser(httpdbClient *client.HttpDBClient, username string) error {
	_, err := httpdbClient.DeleteUser(username)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(httpdbClient *client.HttpDBClient, userCreate *UserCreate) (*httpdbuser.User, error) {
	httpdbUser := new(httpdbuser.User)
	httpdbUser.Name = userCreate.Name
	httpdbUser.Password = userCreate.Password
	httpdbUser.Cpus = userCreate.Cpus
	httpdbUser.Memory = userCreate.Memory

	_, err := httpdbClient.CreateUser(httpdbUser)
	return httpdbUser, err
}

func UpdateResource(httpdbClient *client.HttpDBClient, userinfo *httpdbuser.User) error {
	_, err := httpdbClient.UpdateResource(userinfo)
	return err
}
