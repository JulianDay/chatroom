package config

import (
	"bufio"
	"chatroom/dfa"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ServerCfg struct {
	Addr string `json:"addr"`
}

type SensitiveCfg struct {
	Words []string `json:"word"`
}

var SensitiveWords SensitiveCfg
var Server ServerCfg

func loadSensitive() {
	var words []string
	r, _ := os.Open(RootDir + "/config/sensitivelist.txt")
	defer r.Close()

	s := bufio.NewScanner(r)
	for s.Scan() {
		words = append(words, s.Text())
	}

	dfa.Init(SensitiveWords.Words)
	log.Println("all sensitive words:", len(words))
}

func loadServer() {
	err := loadJson(RootDir+"/config/server.json", &Server)
	if err != nil {
		panic(err)
	}
}

func loadJson(filename string, conf interface{}) error {
	data, err := ioutil.ReadFile(filename) //read config file
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return err
	}
	return nil
}
