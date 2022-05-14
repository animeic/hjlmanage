package ws

import (
	"animeic-gin/app/response"
	"animeic-gin/global"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func (manager *ClientManager) Start() {
	for {
		log.Println("<---监听管道通信--->")
		select {
		case conn := <-Manager.Register: // 建立连接
			log.Printf("建立新连接: %v", conn.ID)
			Manager.Clients[conn.ID] = conn
			replyMsg := &ReplyMsg{
				Code:    response.WebSocketCode,
				Content: "已连接至服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister: // 断开连接
			log.Printf("连接失败:%v", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    response.WebSocketCode,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		//广播信息
		case broadcast := <-Manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false // 默认对方不在线
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID
			if flag {
				log.Println("对方在线应答")
				replyMsg := &ReplyMsg{
					Code:    response.WebSocketCode,
					Content: "对方在线应答",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(global.App.Config.Mongo.DbName, id, string(message), 1, int64(3*month))
				if err != nil {
					fmt.Println("InsertOneMsg Err", err)
				}
			} else {
				log.Println("对方不在线")
				replyMsg := ReplyMsg{
					Code:    response.WebSocketCode,
					Content: "对方不在线应答",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(global.App.Config.Mongo.DbName, id, string(message), 0, int64(3*month))
				if err != nil {
					fmt.Println("InsertOneMsg Err", err)
				}
			}
		}
	}
}
