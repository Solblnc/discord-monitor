package discord

import (
	"discord-monitor/internal/binance"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

const frequency = 5

func Start(token string) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error in creating a discord session", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening discord connection,", err)
		return
	}

	err = dg.UpdateListeningStatus("Tinkoff p2p prices")
	if err != nil {
		fmt.Printf("Unable to set activity: \n", err)
	} else {
		fmt.Println("Monitor is started...")
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
	}

	changeFrequency := frequency * time.Second
	var nickname string

	// watch p2p price
	for {

		// get gas prices
		priceBuy, err := binance.SendReqBuy()
		if err != nil {
			fmt.Printf("Error getting rates: %s\n", err)
			time.Sleep(changeFrequency)
			continue
		}

		priceSell, err := binance.SendReqSell()
		if err != nil {
			fmt.Printf("Error getting rates: %s\n", err)
			time.Sleep(changeFrequency)
			continue
		}

		nickname = fmt.Sprintf("%s | BUY | %s | SELL  ", priceBuy, priceSell)

		// change nickname

		for _, g := range guilds {

			err = dg.GuildMemberNickname(g.ID, "@me", nickname)
			if err != nil {
				fmt.Printf("Error updating nickname: %s\n", err)
				continue
			} else {
				fmt.Printf("Set price in %s: %s\n", g.Name, nickname)
			}
		}

		time.Sleep(changeFrequency)

	}
}
