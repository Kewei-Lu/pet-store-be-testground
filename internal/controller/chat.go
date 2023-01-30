package controller

import (
	"encoding/json"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Chat = ChatRoom{UserMap: gmap.NewStrAnyMap(true)}
)

type ChatRoom struct {
	UserMap *gmap.StrAnyMap // Map structure: [userName:string]*ChatUser
}

type ChatUser struct {
	Name      string
	WebSocket *ghttp.WebSocket
}

type InputMsg struct {
	SenderName   string
	Msg          string
	SendTime     int
	ReceiverName string
	Type         int // 1: Authentication msg 2: common msg
}

type OutputMsg struct {
	SenderName   string
	Msg          string
	SendTime     int
	ReceiverName string
	Type         int // 1: Authentication msg 2: common msg
	UUID         uint64
}

func (c *ChatRoom) RequestWebSocket(r *ghttp.Request) {
	var ctx = r.Context()
	g.Log().Print(ctx, "WebSocket connected")
	// c.UserMap.Set(key string, val interface{})
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(ctx, "WebSocket error:", err)
		r.Exit()
	}
	// Validation stage
	var validationMsg InputMsg
	_, msg, err := ws.ReadMessage()
	if err != nil {
		g.Log().Error(ctx, "Error in receiving validation msg", err)
	}
	if err := json.Unmarshal(msg, &validationMsg); err != nil {
		g.Log().Error(ctx, "Unmarshal error", err)
	}
	g.Log().Debug(ctx, "Validation Msg", validationMsg)

	if validationMsg.SenderName == "" || validationMsg.Msg != "I am an valid user" {
		g.Log().Warning(ctx, "Invalid chat user")
		ws.Close()
		r.Exit()
		return
	}

	// Add socket instance into map
	g.Log().Debugf(ctx, "Validation success, Add websocket to the map, current list: %+v", c.UserMap)
	c.UserMap.Set(validationMsg.SenderName, &ChatUser{Name: validationMsg.SenderName, WebSocket: ws})
	// Listening
	for {
		g.Log().Debug(ctx, "Waiting for message from client")
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			g.Log().Infof(ctx, "User %+v disconnect, remove that user from UserMap", ws)
			c.UserMap.Remove(validationMsg.SenderName)

			return
		}
		var newMsg InputMsg
		var uuidMsg OutputMsg
		g.Log().Print(ctx, "Receive Message through Websocket", ws, "msg: ", msg)
		// add a hashcode for msg
		uuid := ghash.AP64(msg)
		g.Log().Debug(ctx, "uuid:", uuid)
		if err := json.Unmarshal(msg, &newMsg); err != nil {
			g.Log().Print(ctx, "Unmarshal error", err)
		}
		uuidMsg = OutputMsg{
			SenderName:   newMsg.SenderName,
			Msg:          newMsg.Msg,
			SendTime:     newMsg.SendTime,
			ReceiverName: newMsg.ReceiverName,
			Type:         newMsg.Type,
			UUID:         uuid,
		}
		msgToSend, err := json.Marshal(uuidMsg)
		if err != nil {
			g.Log().Error(ctx, "Error in marshal new constructed msg")
		}
		// broadcast message
		if newMsg.ReceiverName == "@ALL" {
			c.UserMap.Iterator(func(k string, v interface{}) bool {
				// assert v is *ChatUser
				user, ok := v.(*ChatUser)
				if !ok {
					g.Log().Print(ctx, "error in asserting *ChatUser")
					return true
				}
				if err = user.WebSocket.WriteMessage(msgType, msgToSend); err != nil {
					g.Log().Print(ctx, "error in Sending Message to", user.Name, "err:", err)
					return true
				}
				return true
			})
		}
	}
}
