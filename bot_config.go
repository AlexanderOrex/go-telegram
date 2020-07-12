package telegram

type BotConfig struct {
  Token string
  BotName string
}

func InitConfig(botName string, token string) BotConfig {
  var config BotConfig
  config.Token = token
  config.BotName = botName

  return config
}
