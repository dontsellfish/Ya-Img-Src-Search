package main

import (
	"encoding/json"
	"gopkg.in/telebot.v3"
	"log"
	"main/tgyaimg"
	"os"
	"time"
)

const ConfigPath = "cfg.txt"

type Cfg struct {
	Whitelist  []string `json:"whitelist"`
	ImgbbToken string   `json:"imgbb_token"`
	TgToken    string   `json:"tg_token"`
}

func main() {
	buffer, err := os.ReadFile(ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	var cfg Cfg
	err = json.Unmarshal(buffer, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	tgyaimg.Start(telebot.Settings{
		Poller:      &telebot.LongPoller{Timeout: time.Second * 60},
		Token:       cfg.TgToken,
		Synchronous: true,
	},
		cfg.ImgbbToken,
		cfg.Whitelist...)

	return
}
