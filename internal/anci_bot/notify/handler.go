package notify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/Andronovdima/tg-bot-helper/internal/anci_bot/locations"
)

type Handler struct {
	locationUsecase *locations.Usecase
	logger          *zap.SugaredLogger
}

func NewNotifyHandler(m *mux.Router, logger *zap.SugaredLogger, locationUsecase *locations.Usecase) {
	handler := &Handler{
		locationUsecase: locationUsecase,
		logger:          logger,
	}

	m.HandleFunc("/notify", handler.HandleNotify).Methods(http.MethodPost, http.MethodOptions)
}

func (h *Handler) HandleNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request", time.Now())
	defer func() {
		if err := r.Body.Close(); err != nil {
			return
		}
	}()

	var req notifyRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.locationUsecase.GetRandomLocation(req.Message.Chat.ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
