package main

import (
	//"fmt"

	"github.com/wzzlYwzzl/httpdatabase/client"
	"github.com/wzzlYwzzl/httpdatabase/resource/user"
)

type UserListChannel struct {
	List  chan *user.UserList
	Error chan error
}

func GetUserListChannel(client client.Client, numReads int) UserListChannel {
	channel := UserListChannel{
		List:  make(chan *user.UserList, numReads),
		Error: make(chan error, numReads),
	}

	go func() {
		rcs, err := client.GetAllUserInfo()
		for i := 0; i < numReads; i++ {
			channel.List <- rcs
			channel.Error <- err
		}
	}()

	return channel
}

func main() {
	// 	client := client.Client{Host: "localhost:9080"}
	// 	channel := GetUserListChannel(client, 2)

	// 	res := <-channel.List
	// 	err := <-channel.Error
	// 	if err != nil {
	// 		fmt.Println("func err", err)
	// 	}

	// 	fmt.Println(res)
	// }
	TestBasic()
}
