package config

type Config struct {
	BindAddr               string
	LogLevel               string
	ClientUrl              string
	TelegramBotAccessToken string
}

func NewConfig() *Config {
	return &Config{
		BindAddr:               ":9000",
		LogLevel:               "debug",
		ClientUrl:              "",
		TelegramBotAccessToken: "",
	}
}
