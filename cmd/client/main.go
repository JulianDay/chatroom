package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"os"
	"time"
)

type Player struct {
	Id      uint64    `json:"uid"`
	Name    string    `json:"name"`
	NewTime time.Time `json:"newtime"`
	Addr    string    `json:"addr"`
}
type Message struct {
	Player         *Player   `json:"player"`
	Content        string    `json:"content"`
	MsgTime        time.Time `json:"msgtime"`
	ClientSendTime time.Time `json:"client_send_time"`
}

func main() {
	var name string
	fmt.Println("input your name:")
	fmt.Scanln(&name)
	c, _, err := websocket.Dial(context.Background(), "ws://localhost:8088/login?name="+name, nil)
	if err != nil {
		panic(err)
	}

	defer c.Close(websocket.StatusInternalError, "内部错误！")
	go func() {
		for {
			var v Message
			err = wsjson.Read(context.Background(), c, &v)
			if err != nil {
				panic(err)
			}
			str := "->"
			if v.Player != nil {
				str = (v.Player.Name + ":")
			}
			str += v.Content
			fmt.Println(str)
		}
	}()
	fmt.Println("input 'q' quit chat")
	for {
		reader := bufio.NewReader(os.Stdin)
		bytes, _, _ := reader.ReadLine()
		content := string(bytes)
		if content == "q" {
			break
		}
		chatConnent := map[string]string{}
		chatConnent["content"] = content
		err = wsjson.Write(context.Background(), c, chatConnent)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
		}
	}
	c.Close(websocket.StatusNormalClosure, "")
}
