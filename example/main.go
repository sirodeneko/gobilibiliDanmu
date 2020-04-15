package main

import (
	"log"
	"time"

	"github.com/sirodeneko/gobilibiliDanmu"
)

func main() {
	bili := gobilibili.NewBiliBiliClient()
	// bili.RegHandleFunc(gobilibili.CmdAll, gobilibili.DefaultHandler)
	// bili.RegHandleFunc(gobilibili.CmdAll, gobilibili.DebugHandler)

	// 以下一系列语句在把命令（cmd)对应的函数写入bili.handlerMap
	bili.RegHandleFunc(gobilibili.CmdDanmuMsg, func(c *gobilibili.Context) bool {
		dinfo := c.GetDanmuInfo()
		if dinfo.Uname != "" {
			log.Printf("[%d]%s(%d) 说: %s\r\n", c.RoomID, dinfo.Uname, dinfo.UID, dinfo.Text)
		} else {
			log.Printf("[%d]%d 说: %s\r\n", c.RoomID, dinfo.UID, dinfo.Text)
		}
		return false
	})
	bili.RegHandleFunc(gobilibili.CmdWelcome, func(c *gobilibili.Context) bool {
		winfo := c.GetWelcomeInfo()
		if winfo.Uname != "" {
			log.Printf("[%d]%s(%d) 进入了房间\r\n", c.RoomID, winfo.Uname, winfo.UID)
		} else {
			log.Printf("[%d]%d 进入了房间\r\n", c.RoomID, winfo.UID)
		}
		return false
	})

	bili.RegHandleFunc(gobilibili.CmdSendGift, func(c *gobilibili.Context) bool {
		gInfo := c.GetGiftInfo()
		log.Printf("[%d]%s(%d) %s 了 %s x %d (价值%.3f)\r\n", c.RoomID, gInfo.Uname, gInfo.UID, gInfo.Action, gInfo.GiftName, gInfo.Num, float32(gInfo.Price*gInfo.Num)/1000)
		return false
	})

	bili.RegHandleFunc(gobilibili.CmdOnlineChange, func(c *gobilibili.Context) bool {
		online := c.GetOnlineNumber()
		log.Printf("[%d]房间里当前在线：%d\r\n", c.RoomID, online)
		return false
	})

	bili.RegHandleFunc(gobilibili.CmdNoticeMsg, func(c *gobilibili.Context) bool {
		nMsg := c.GetNoticeMsg()
		log.Printf("[%d]系统消息通知: %s\r\n", c.RoomID, nMsg.MsgCommon)
		return false
	})

	for {
		// 传入房间号
		err := bili.ConnectServer(102)
		log.Println("与弹幕服务器连接中断,3秒后重连。原因:", err.Error())
		time.Sleep(time.Second * 3)
	}
}
