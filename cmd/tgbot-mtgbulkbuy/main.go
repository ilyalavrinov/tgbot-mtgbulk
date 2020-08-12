package main

import (
	"flag"

	"github.com/admirallarimda/tgbotbase"
	"github.com/ilyalavrinov/tgbot-mtgbulkbuy/internal/bot"
	"github.com/ilyalavrinov/tgbot-mtgbulkbuy/internal/log"
	"gopkg.in/gcfg.v1"
)

var argCfg = flag.String("cfg", "./bot.cfg", "path to config")

type config struct {
	tgbotbase.Config
}

func main() {
	flag.Parse()

	var cfg config

	if err := gcfg.ReadFileInto(&cfg, *argCfg); err != nil {
		log.Fatalw("Cannot read config file",
			"filename", *argCfg)
	}

	tgbot := tgbotbase.NewBot(tgbotbase.Config{TGBot: cfg.TGBot, Proxy_SOCKS5: cfg.Proxy_SOCKS5})
	tgbot.AddHandler(tgbotbase.NewIncomingMessageDealer(bot.NewSearchHandler()))

	log.Info("Starting bot")
	tgbot.Start()
	log.Info("Stopping bot")
}
