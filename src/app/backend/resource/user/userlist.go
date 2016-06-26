package user

import (
	"log"

	"github.com/kubernetes/dashboard/client"
	"github.com/kubernetes/dashboard/resource/common"
	httpdbuser "github.com/wzzlYwzzl/httpdatabase/resource/user"
)

//GetUserList return a list of all users in the cluster.
func GetUserListFromChannels(channels *common.ResourceChannels, httpdbClient *client.HttpDBClient) (
	*httpdbuser.UserList, error) {

	users := <-channels.UserList.List
	err := <-channels.UserList.Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserList(httpdbClient *client.HttpDBClient) (*httpdbuser.UserList, error) {
	log.Println("Getting list of users from the httpdb server")

	channels := &common.ResourceChannels{
		UserList: common.GetUserListChannel(httpdbClient, 1),
	}

	return GetUserListFromChannels(channels, httpdbClient)
}
