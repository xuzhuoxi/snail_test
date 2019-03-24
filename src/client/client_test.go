//
//Created by xuzhuoxi
//on 2019-02-24.
//@author xuzhuoxi
//
package client

import (
	"fmt"
	"github.com/xuzhuoxi/snail_test/src/client/internel"
	"sync"
	"testing"
	"time"
)

const UserCount = 300

var (
	UserIds     []string
	UserClients map[string]*internel.UserClient
	sleep       = make(chan struct{})
	once        sync.Once
)

func init() {
	for index := 0; index < UserCount; index++ {
		UserIds = append(UserIds, fmt.Sprintf("u_%d", 10000+index))
	}
	UserClients = make(map[string]*internel.UserClient)
	for _, userId := range UserIds {
		uc := internel.NewUserClient(userId)
		err := uc.Open()
		if nil != err {
			continue
		}
		UserClients[uc.UserId] = uc
	}
}

func TestClient(t *testing.T) {
	TestLogin(t)
	for {
		time.Sleep(time.Second)
		TestReLogin(t)
	}
	<-sleep
}

func TestLogin(t *testing.T) {
	for _, val := range UserClients {
		val.TestLoginExtension()
	}
}

func TestReLogin(t *testing.T) {
	for _, val := range UserClients {
		val.TestReLoginExtension()
	}
}

//
//func startClient(uc *userClient) {
//	go mgrAtRobot(uc)
//}
//
//func mgrAtRobot(uc *userClient) {
//	if err := uc.Open(); err != nil {
//		fmt.Println(err)
//		return
//	}
//	uc.TestLoginExtension()
//	for {
//		time.Sleep(time.Second)
//		uc.TestReLoginExtension()
//	}
//}
