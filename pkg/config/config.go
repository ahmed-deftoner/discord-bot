package config

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

var (
	Token     string
	BotPrefix string
	Config    *configStruct
)
