//
//Created by xuzhuoxi
//on 2019-03-24.
//@author xuzhuoxi
//
package internel

import (
	"github.com/json-iterator/go"
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/netx"
)

const (
	RemoteAddress = "127.0.0.1:31000"
	Network       = netx.TcpNetwork
)

var ClientCreator = netx.NewTCPClient

func NewUserClient(uId string) *UserClient {
	client := ClientCreator()
	return &UserClient{SockClient: client, UserId: uId}
}

type UserClient struct {
	UserId     string
	SockClient netx.ISockClient
}

func (uc *UserClient) Open() error {
	return uc.SockClient.OpenClient(netx.SockParams{RemoteAddress: RemoteAddress, Network: Network})
}

func (uc *UserClient) TestLoginExtension() {
	buffToBlock := bytex.NewBuffToBlock(bytex.NewDefaultDataBlockHandler())
	buffToBlock.WriteData([]byte("Login"))   //ExtensionName
	buffToBlock.WriteData([]byte("LI"))      //ProtoId
	buffToBlock.WriteData([]byte(uc.UserId)) //Uid
	buffToBlock.WriteData([]byte(uc.UserId)) //Data(Password)
	uc.SockClient.SendPackTo(buffToBlock.ReadBytes())
}

func (uc *UserClient) TestReLoginExtension() {
	buffToBlock := bytex.NewBuffToBlock(bytex.NewDefaultDataBlockHandler())
	buffToBlock.WriteData([]byte("Login"))   //ExtensionName
	buffToBlock.WriteData([]byte("RLI"))     //ProtoId
	buffToBlock.WriteData([]byte(uc.UserId)) //Uid
	buffToBlock.WriteData([]byte(uc.UserId)) //Data(Password)
	uc.SockClient.SendPackTo(buffToBlock.ReadBytes())
}

func (uc *UserClient) TestDemoExtension() {
	bsName := []byte("ObjDemo")
	bsPid := []byte("Obj_0")
	bsUid := []byte("顶你个肺")
	data := testA{A: "A", B: 99, C: false}

	buffToBlock := bytex.NewBuffToBlock(bytex.NewDefaultDataBlockHandler())
	buffToBlock.WriteData(bsName)
	buffToBlock.WriteData(bsPid)
	buffToBlock.WriteData(bsUid)
	dataBs, _ := jsoniter.Marshal(data)
	buffToBlock.WriteData(dataBs)
	buffToBlock.WriteData(dataBs)
	uc.SockClient.SendPackTo(buffToBlock.ReadBytes())
}
