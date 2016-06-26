package main

import (
	"fmt"
	"log"

	"github.com/wzzlYwzzl/httpdatabase/client"
	"github.com/wzzlYwzzl/httpdatabase/resource/user"
)

func TestBasic() {
	clientConf := client.Client{Host: "localhost:9080"}
	user := &user.User{Name: "test-zjw", Password: "test-zjw", Cpus: 10, Memory: 500, CpusUse: 8, MemoryUse: 123}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//test 1
	b, err := clientConf.JudgeName("admin", "admin123")
	if err != nil {
		fmt.Println("judegeName Error")
		return
	}
	if b == true {
		fmt.Println("JudgeName work well")
	} else {
		fmt.Println("JudgeName doesn't work")
	}

	//test2
	b, err = clientConf.CreateNS("admin", "admin_namespace_test")
	if err != nil {
		fmt.Println("CreateNS Error")
		return
	}
	if b == true {
		fmt.Println("CreateNS work well")
	} else {
		fmt.Println("CreateNS doesn't work")
	}

	//test3
	ns, err := clientConf.GetNS("admin")
	if err != nil {
		fmt.Println("GetNS Error")
		return
	}
	if len(ns) != 0 {
		fmt.Println("GetNS work well")
		fmt.Println(ns)
	} else {
		fmt.Println("GetNS doesn't work")
	}

	//test4
	ns, err = clientConf.GetNSAll("admin")
	if err != nil {
		fmt.Println("GetNSAll Error")
		return
	}
	if len(ns) != 0 {
		fmt.Println("GetNSAll work well")
		fmt.Println(ns)
	} else {
		fmt.Println("GetNSAll doesn't work")
	}

	//test5
	us, err := clientConf.GetAllInfo("admin")
	if err != nil {
		fmt.Println("GetAllInfo Error")
		return
	}
	if len(ns) != 0 {
		fmt.Println("GetAllInfo work well")
		fmt.Println(us)
	} else {
		fmt.Println("GetAllInfo doesn't work")
	}

	//test6
	b, err = clientConf.CreateUser(user)
	if err != nil {
		fmt.Println("CreateUser Error")
		return
	}
	if b == true {
		fmt.Println("CreateUser work well")
	} else {
		fmt.Println("CreateUser doesn't work")
	}

	// //test7
	// b, err = clientConf.DeleteUser("test-zjw")
	// if err != nil {
	// 	fmt.Println("DeleteUser Error")
	// 	return
	// }
	// if b == true {
	// 	fmt.Println("DeleteUser work well")
	// } else {
	// 	fmt.Println("DeleteUser doesn't work")
	// }

	//test8
	b, err = clientConf.UpdateResource(user)
	if err != nil {
		fmt.Println("UpdateResource Error")
		return
	}
	if b == true {
		fmt.Println("UpdateResource work well")
	} else {
		fmt.Println("UpdateResource doesn't work")
	}

	//test9
	userlist, err := clientConf.GetAllUserInfo()
	if err != nil {
		fmt.Println("GetAllUserInfo Error")
		return
	} else {
		fmt.Println(userlist)
	}

}
