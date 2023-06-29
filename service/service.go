package service

import (
	"context"
	"github.com/acom21/pet-proj-dsandbox/config"
	"github.com/acom21/pet-proj-dsandbox/pkg/mongodb"
	tele "gopkg.in/telebot.v3"
)

type Service struct {
	cfg  *config.Config
	Bot  tBot
	Rate currRates
	Repo mongodb.Repo
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		cfg:  cfg,
		Bot:  &TgBot{accessToken: cfg.TgAPI.AccessToken},
		Rate: &Converter{API: cfg.CurrencyRatesAPI},
	}
}

type tBot interface {
	SetupBot(accessToken string) (*tele.Bot, error)
}

type currRates interface {
	ConvertCurrency(ctx context.Context) (*ConverterResp, error)
}

func (s *Service) Run() error {

	bot, err := s.Bot.SetupBot(s.cfg.TgAPI.AccessToken)
	if err != nil {
		return err
	}
	bot.Close()
	return nil
}

func (s *Service) calculate() {

}
