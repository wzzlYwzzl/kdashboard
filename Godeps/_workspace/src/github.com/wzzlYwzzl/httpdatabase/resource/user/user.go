package user

import (
	"log"

	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

type User struct {
	Name       string   `json:"name"`
	Password   string   `json:"password"`
	Namespaces []string `json:"namespaces"`
	Cpus       int      `json:"cpus"`
	Memory     int      `json:"memory"`
	CpusUse    int      `json:"cpususe"`
	MemoryUse  int      `json:"memoryuse"`
}

type UserList struct {
	UserList []User `json:"userList"`
}

func (user User) JudgeExist(dbconf *sqlop.MysqlCon) (bool, error) {
	dbuser := new(sqlop.UserInfo)
	dbuser.Name = user.Name
	dbuser.Password = user.Password

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return false, err
	}

	defer db.Close()

	err = dbuser.QueryOne(db)
	if err != nil {
		log.Printf("return false")
		return false, err
	}

	if dbuser.Name == "" {
		return false, nil
	}

	return true, nil
}

func (user *User) GetNamespacesAll(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	res, err := dbuser.QueryAll(db)
	if err != nil || len(res) == 0 {
		return err
	}

	user.Namespaces = res
	return nil
}

func (user *User) GetNamespaces(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	res, err := dbuser.Query(db)
	if err != nil || len(res) == 0 {
		return err
	}

	user.Namespaces = res
	return nil
}

func (user *User) CreateNamespace(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name
	dbuser.Namespace = user.Namespaces[0]

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = dbuser.Insert(db)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) CreateUser(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)
	resource := new(sqlop.UserResource)
	userns := new(sqlop.User)

	userinfo.Name = user.Name
	userinfo.Password = user.Password
	userinfo.Cpus = user.Cpus
	userinfo.Mem = user.Memory
	resource.Name = user.Name
	userns.Name = user.Name
	userns.Namespace = user.Name + "-default"

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		log.Println("file : user.go, function")
		return err
	}

	defer db.Close()

	err = userinfo.Insert(db)
	if err != nil {
		return err
	}

	//At the same time, create one namespace entry in table userns
	err = userns.Insert(db)
	if err != nil {
		return err
	}

	//At the same time, create one resource entry
	err = resource.Insert(db)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) DeleteUser(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)
	userinfo.Name = user.Name

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = userinfo.Delete(db)
	if err != nil {
		log.Println("DeleteUser error :", err)
		return err
	}

	return nil
}

func (user *User) GetUser(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)
	userinfo.Name = user.Name

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = userinfo.Query(db)
	if err != nil {
		return err
	}

	user.Password = userinfo.Password
	user.Cpus = userinfo.Cpus
	user.Memory = userinfo.Mem

	return nil
}

func (user *User) UpdateResource(dbconf *sqlop.MysqlCon) error {
	userrs := new(sqlop.UserResource)
	userrs.Name = user.Name
	userrs.CpusUse = user.CpusUse
	userrs.MemUse = user.MemoryUse

	db, err := userrs.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = userrs.Update(db)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (user *User) GetAllInfo(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	userinfo := new(sqlop.UserInfo)
	rs := new(sqlop.UserResource)
	userinfo.Name = user.Name
	dbuser.Name = user.Name
	rs.Name = user.Name

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	user.Namespaces, err = dbuser.Query(db)
	if err != nil {
		log.Println(err)
		return err
	}

	err = userinfo.Query(db)
	if err != nil {
		log.Println(err)
		return err
	}

	err = rs.GetRS(db)
	if err != nil {
		log.Println(err)
		return err
	}

	user.Password = userinfo.Password
	user.Cpus = userinfo.Cpus
	user.Memory = userinfo.Mem
	user.CpusUse = rs.CpusUse
	user.MemoryUse = rs.MemUse

	return nil
}

func (userlist *UserList) GetAllUserInfo(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	users, err := userinfo.QueryUsers(db)
	if err != nil {
		log.Println(err)
		return err
	}

	userlist.UserList = make([]User, len(users))
	for i := 0; i < len(users); i++ {
		userlist.UserList[i].Name = users[i]
		userlist.UserList[i].GetAllInfo(dbconf)
	}

	return nil
}
