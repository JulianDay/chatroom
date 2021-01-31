package logic

import (
	"context"
	"errors"
	"fmt"
	"io"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

// 用户
type Player struct {
	Name    string    `json:"name"`
	NewTime time.Time `json:"newtime"`
	Addr    string    `json:"addr"`

	conn  *websocket.Conn
	isNew bool
}

func NewPlayer(conn *websocket.Conn, name, addr string) *Player {
	p := &Player{
		Name:    name,
		Addr:    addr,
		NewTime: time.Now(),
		conn:    conn,
	}
	return p
}

func (p *Player) Read(ctx context.Context) error {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {
		err = wsjson.Read(ctx, p.conn, &receiveMsg)
		if err != nil {
			// 判定连接是否关闭了，正常关闭，不认为是错误
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			} else if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		// 内容发送到聊天室
		content := receiveMsg["content"]
		dispatherMessage(p, content)
	}
}

func (p *Player) SendMessage(msg *Message) {
	err := wsjson.Write(context.Background(), p.conn, msg)
	if err != nil {
		fmt.Println("send msg:", err)
	}
}

func (p *Player) KickOff() {
	p.conn.Close(websocket.StatusUnsupportedData, "server kickoff")
}
