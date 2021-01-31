package logic

import (
	"chatroom/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

type handler func(*Player, string)

var messageHandler map[string]handler

func registerHandler(prefix string, h handler) {
	if messageHandler == nil {
		messageHandler = make(map[string]handler)
	}
	messageHandler[prefix] = h
}

func dispatherMessage(p *Player, content string) {
	for prefix, handler := range messageHandler {
		if strings.HasPrefix(content, prefix) {
			handler(p, content)
			return
		}
	}
	chatHandler(p, content)
}
func init() {
	log.Println("message handler init")
	registerHandler("/popular", popularHandler)
	registerHandler("/stats", statsHandler)
}

func popularHandler(p *Player, content string) {
	words := Popular.GetPopularWords()
	p.SendMessage(NewSystemMessage(words.String()))
}

func statsHandler(p *Player, content string) {
	params := strings.Split(content, " ")
	if len(params) != 2 {
		p.SendMessage(NewErrorMessage("please input /stats [username]"))
		return
	}
	playerName := params[1]
	player := Room.GetPlayer(playerName)
	if player == nil {
		p.SendMessage(NewErrorMessage("no player [username]"))
		return
	}
	d := time.Now().Sub(player.NewTime)
	p.SendMessage(NewSystemMessage(d.String()))
}

func chatHandler(p *Player, content string) {
	sendMsg := NewMessage(p, content)
	sendMsg.Content = FilterSensitive(content)
	fmt.Println(p.Name + " send:" + sendMsg.Content)

	Room.Broadcast(sendMsg)
	LastMessage.Save(sendMsg)
	words := utils.SplitWords(sendMsg.Content)
	Popular.AddWords(words)
}
