package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wzzlYwzzl/httpdatabase/resource/user"
)

type HttpDBClient struct {
	Host string
}

func (c HttpDBClient) JudgeName(name string, password string) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/" + name + "/" + password
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func (c HttpDBClient) CreateNS(name, namespace string) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/" + name + "/" + namespace
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusCreated {
		return true, nil
	}

	return false, nil
}

func (c HttpDBClient) GetNS(name string) ([]string, error) {
	ns := make([]string, 0, 10)
	url := "http://" + c.Host + "/api/v1/user/ns/" + name
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http get error", err)
		return ns, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll error :", err)
		return ns, err
	}

	err = json.Unmarshal(body, &ns)
	if err != nil {
		log.Println("json Unmarshal  error :", err)
		return ns, err
	}
	return ns, nil
}

func (c HttpDBClient) GetNSAll(name string) ([]string, error) {
	ns := make([]string, 0, 10)
	url := "http://" + c.Host + "/api/v1/user/ns/all/" + name
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http get error", err)
		return ns, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll error :", err)
		return ns, err
	}
	err = json.Unmarshal(body, &ns)
	if err != nil {
		log.Println("json Unmarshal  error :", err)
		return ns, err
	}
	return ns, nil
}

func (c HttpDBClient) GetAllInfo(name string) (*user.User, error) {
	user := new(user.User)
	url := "http://" + c.Host + "/api/v1/user/allinfo/" + name
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return user, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll in func GetAllInfo err: ", err)
		return user, err
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		log.Println("json Unmarshal err: ", err)
		return user, err
	}

	return user, nil
}

func (c HttpDBClient) DeleteUser(name string) (bool, error) {
	client := &http.Client{}
	url := "http://" + c.Host + "/api/v1/user/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
		return false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func (c HttpDBClient) CreateUser(user *user.User) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/"
	bs, err := json.Marshal(*user)
	if err != nil {
		log.Println(err)
		return false, err
	}

	reader := bytes.NewReader(bs)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusCreated {
		return true, nil
	}

	return false, nil
}

func (c HttpDBClient) UpdateResource(user *user.User) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/resource"
	bs, err := json.Marshal(*user)
	if err != nil {
		log.Println(err)
		return false, err
	}

	reader := bytes.NewReader(bs)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusCreated {
		return true, nil
	}

	return false, nil
}

func (c HttpDBClient) GetAllUserInfo() (*user.UserList, error) {
	userlist := new(user.UserList)
	url := "http://" + c.Host + "/api/v1/user/all"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll in func GetAllInfo err: ", err)
		return nil, err
	}

	err = json.Unmarshal(body, userlist)
	if err != nil {
		log.Println("json Unmarshal err: ", err)
		return nil, err
	}

	return userlist, nil
}

func (c HttpDBClient) AddApp(deploy *user.UserDeploy) (bool, error) {
	url := "http://" + c.Host + "/api/v1/user/app"

	bs, err := json.Marshal(*deploy)
	if err != nil {
		log.Println(err)
		return false, err
	}

	reader := bytes.NewReader(bs)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if resp.StatusCode == http.StatusCreated {
		return true, nil
	}

	return false, nil
}

func (c HttpDBClient) DeleteApp(appName string) (*user.UserDeploy, error) {
	client := &http.Client{}
	url := "http://" + c.Host + "/api/v1/user/app" + appName
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ReadAll in func GetAllInfo err: ", err)
			return nil, err
		}

		deploy := new(user.UserDeploy)

		err = json.Unmarshal(body, deploy)
		if err != nil {
			log.Println("json Unmarshal err: ", err)
			return nil, err
		}

		return deploy, nil
	}

	return nil, nil
}
