package main

import (
	"discord-monitor/internal/config"
	"discord-monitor/internal/discord"
)

func main() {
	token := config.FromEnv("token")
	discord.Start(token)
}
