package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	config "github.com/Andronovdima/tg-bot-helper/configs/anci_bot"
	"github.com/Andronovdima/tg-bot-helper/internal/anci_bot/locations"
	"github.com/Andronovdima/tg-bot-helper/internal/anci_bot/middleware"
	"github.com/Andronovdima/tg-bot-helper/internal/anci_bot/notify"
	"github.com/Andronovdima/tg-bot-helper/pkg/telegram"
)

const tgURL = "https://api.telegram.org"

type Server struct {
	Mux    *mux.Router
	Config *config.Config
	Logger *zap.SugaredLogger
}

func NewServer(config *config.Config, logger *zap.SugaredLogger) (*Server, error) {
	s := &Server{
		Mux:    mux.NewRouter(),
		Logger: logger,
		Config: config,
	}
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func (s *Server) ConfigureServer() {
	mdlwr := middleware.NewMiddleware(s.Logger)

	tgClient := telegram.NewClient(tgURL)
	locationsUsecase := locations.NewUsecase(tgClient, s.Config.TelegramBotAccessToken)
	s.Mux.Use(mdlwr.CORSMiddleware)
	notify.NewNotifyHandler(s.Mux, s.Logger, locationsUsecase)

	return
}
