package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"

	config "github.com/Andronovdima/tg-bot-helper/configs/anci_bot"
)

func Start() error {
	cfg := config.NewConfig()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8777"
	} else {
		port = ":" + port
	}

	tgBotAccessToken := os.Getenv("ANCI_BOT_ACCESS_TOKEN")
	if tgBotAccessToken == "" {
		return errors.New("no ANCI_BOT_ACCESS_TOKEN in ENV")
	}

	cfg.BindAddr = port
	cfg.TelegramBotAccessToken = tgBotAccessToken

	zapLogger, _ := zap.NewProduction()
	defer func() {
		if err := zapLogger.Sync(); err != nil {
			log.Println("zap logger error", err)
		}
	}()
	sugaredLogger := zapLogger.Sugar()

	srv, err := NewServer(cfg, sugaredLogger)
	if err != nil {
		return err
	}

	srv.Config = cfg
	srv.ConfigureServer()

	fmt.Println("Start server on port " + cfg.BindAddr)

	return http.ListenAndServe(cfg.BindAddr, srv)
}
