package main

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/snail"
	"github.com/xuzhuoxi/snail/engine/mmo"
	"github.com/xuzhuoxi/snail/engine/mmo/basis"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var sleep = make(chan struct{})

func main() {
	//initDemoMMO()
	snail.Run(false)
}

func initDemoMMO() {
	mmoMgr := mmo.NewMMOManager()
	mmoMgr.InitManager()
	mmoMgr.SetLogger(logx.DefaultLogger())

	em := mmoMgr.GetEntityManager()
	em.InitWorld("world", "HelloWorld")
	go setVar(em.World(), time.Minute)
	i := 0
	for i < 10 {
		zoneId := "zone" + strconv.Itoa(i)
		z, _ := em.CreateZone(zoneId, zoneId)
		fmt.Println(z)
		go setVar(z, 30*time.Second)
		j := 0
		for j < 10 {
			roomId := "room" + strconv.Itoa(i) + "_" + strconv.Itoa(j)
			r, _ := em.CreateRoomAt(roomId, roomId, zoneId)
			go setVar(r, 5*time.Second)
			j++
		}
		i++
	}

	em.World().ForEachChild(func(child basis.IEntity) (interruptCurrent bool, interruptRecurse bool) {
		logx.Traceln(child.UID())
		return
	})
}

func setVar(entity basis.IEntity, interval time.Duration) {
	if v, ok := entity.(basis.IVariableSupport); ok {
		ran := rand.Intn(10)
		if ran > 5 {
			v.SetVar("temp", rand.Intn(1000))
		} else {
			vs := basis.NewVarSet()
			vs.Set("temp1", rand.Intn(300))
			vs.Set("temp2", rand.Intn(300))
			v.SetVars(vs)
		}
	}
	time.Sleep(interval + time.Duration(rand.Int63n(int64(time.Second)*5)))
	setVar(entity, interval)
}
