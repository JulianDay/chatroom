package server

import (
	"chatroom/logic"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

var System = &logic.Player{
	Name:    "[System]",
	NewTime: time.Time{},
	Addr:    "",
}

func RegisterHandle() {
	http.HandleFunc("/popular", userListHandleFunc)
	http.HandleFunc("/login", LoginHandleFunc)
}

func userListHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func LoginHandleFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}

	// 1. 新用户进来，构建该用户的实例
	name := req.FormValue("name")
	if p := logic.Room.GetPlayer(name); p != nil { // 顶号
		p.KickOff()
	}

	player := logic.NewPlayer(conn, name, req.RemoteAddr)

	// 2.发送欢迎消息
	player.SendMessage(logic.NewWelcomeMessage(player))

	// 3.给所有用户告知新用户到来
	logic.Room.Broadcast(logic.NewPlayerEnterMessage(System))

	// 4. 将该用户加入广播器的用列表中
	logic.Room.PlayerEnter(player)
	log.Println("player:", name, "enter")

	// 5. 接收用户消息
	err = player.Read(req.Context())

	// 6. 用户离开
	logic.Room.PlayerLeaving(player)
	logic.Room.Broadcast(logic.NewPlayerLeaveMessage(System))
	log.Println("player:", name, "leave")

	// 根据读取时的错误执行不同的 Close
	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		log.Println("read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}
