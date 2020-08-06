# go-telegram
Go tools for Telegram API. Helps to manage your Telegram bots.

go get -u github.com/AlexanderOrex/go-telegram
go get gonum.org/v1/plot/...

## Usage example
```
func main() {
    botName := "BOT_NAME"
    botToken := "BOT_TOKEN"
    chatId := 123456789
    config := telegram.InitConfig(botName, botToken, chatId)

    mode := "html"
    messageForTelegram := "You will get some graph soon!"
    telegram.SendMessage(config, messageForTelegram, mode)

    xInchSize := float64(10)
    yInchSize := float64(5)
    title := "Math functions"

    statsByMarketX := make(map[string][]float64)
    statsByMarketX["sin"] = []float64{}
    statsByMarketX["cos"] = []float64{}

    statsByMarketY := make(map[string][]float64)
    statsByMarketY["sin"] = []float64{}
    statsByMarketY["cos"] = []float64{}

    period := 2*math.Pi

    for i := 0.0; i < period; i+=0.01 {
      statsByMarketX["sin"] = append(statsByMarketX["sin"], i)
      statsByMarketY["sin"] = append(statsByMarketY["sin"], math.Sin(i))

      statsByMarketX["cos"] = append(statsByMarketX["cos"], i)
      statsByMarketY["cos"] = append(statsByMarketY["cos"], math.Cos(i))
    }

    telegram.SendPlot(config, statsByMarketX, statsByMarketY, xInchSize, yInchSize, title)
}
```

## Graph example:
![graph example](https://i.ibb.co/vP3vgWj/sin-cos.jpg)
