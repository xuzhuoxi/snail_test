//
//Created by xuzhuoxi
//on 2019-02-24.
//@author xuzhuoxi
//
package client

import (
	"fmt"
	"github.com/xuzhuoxi/snail_test/src/client/internel"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const UserCount = 200 //服务器与测试程序运行在同一ip下,可用端口可能只有300个左右

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
	fmt.Println("TestClient")
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
	fmt.Println("TestLogin", len(UserClients))
}

func TestReLogin(t *testing.T) {
	for _, val := range UserClients {
		val.TestReLoginExtension()
	}
	fmt.Println("TestReLogin", len(UserClients))
}

func TestPressure(t *testing.T) {
	for _, val := range UserClients {
		go mgrAtRobot(val)
	}
	<-sleep
}

func mgrAtRobot(uc *internel.UserClient) {
	if err := uc.Open(); err != nil {
		fmt.Println(err)
		return
	}
	uc.TestLoginExtension()
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))
		uc.TestReLoginExtension()
	}
}
