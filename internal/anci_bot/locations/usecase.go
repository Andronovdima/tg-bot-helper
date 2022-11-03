package locations

import (
	"math/rand"
	"strconv"

	"github.com/Andronovdima/tg-bot-helper/pkg/telegram"
)

type Usecase struct {
	tgClient           *telegram.Client
	anciBotAccessToken string
}

func NewUsecase(tgClient *telegram.Client, botAccessToken string) *Usecase {
	return &Usecase{
		tgClient:           tgClient,
		anciBotAccessToken: botAccessToken,
	}
}

func (u *Usecase) GetRandomLocation(chatID int64) error {
	num := rand.Int()
	if err := u.tgClient.SendMessageFromBot(chatID, "number is "+strconv.Itoa(num), u.anciBotAccessToken); err != nil {
		return err
	}

	return nil
}
