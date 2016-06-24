package user

import (
	"github.com/kubernetes/dashboard/client"
	"github.com/kubernetes/dashboard/resource/common"
	httpdbuser "github.com/wzzlYwzzl/httpdatabase/resource/user"
)

func GetUserListFromChannels(channels *common.ResourceChannels, httpdbClient *client.HttpDBClient) (
	*httpdbuser.UserList, error) {

	users := <-channels.UserList.List
	err := <-channels.UserList.Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
