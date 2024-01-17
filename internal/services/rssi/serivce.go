package rssi

import "github.com/ZecretBone/ips-bff/internal/config"

type Service interface {
	AddingText(text string) string
}

type RSSIService struct {
	cfg config.RSSIServiceConfig
}

func ProvideRSSIService(cfg config.RSSIServiceConfig) Service {
	return &RSSIService{
		cfg: cfg,
	}
}

func (s *RSSIService) AddingText(text string) string {
	return text + "--Adding text"
}
