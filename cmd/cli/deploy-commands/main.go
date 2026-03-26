package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/skyia-jp/shiro-go/internal/commands"
	"github.com/skyia-jp/shiro-go/internal/commands/utility"
	"github.com/skyia-jp/shiro-go/internal/logger"
)

func getenvAny(keys ...string) string {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}

func main() {
	_ = godotenv.Load()

	log, err := logger.Initialize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	token := getenvAny("DISCORD_TOKEN", "DISCORD_BOT_TOKEN")
	appID := getenvAny("DISCORD_APP_ID", "DISCORD_CLIENT_ID")
	guildID := os.Getenv("DISCORD_GUILD_ID")

	if token == "" {
		log.Fatal("DISCORD_TOKEN is required")
	}
	if appID == "" {
		log.Fatal("DISCORD_APP_ID is required")
	}

	router := commands.NewRouter(log, nil)
	if err := utility.RegisterBuiltins(router); err != nil {
		log.Fatalf("failed to register commands: %v", err)
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("failed to create discord session: %v", err)
	}
	defer s.Close()

	defs := router.Definitions()

	if guildID != "" {
		_, err = s.ApplicationCommandBulkOverwrite(appID, guildID, defs)
		if err != nil {
			log.Fatalf("failed to deploy guild commands: %v", err)
		}
		log.Infof("Deployed %d guild commands to guild %s", len(defs), guildID)
		return
	}

	_, err = s.ApplicationCommandBulkOverwrite(appID, "", defs)
	if err != nil {
		log.Fatalf("failed to deploy global commands: %v", err)
	}
	log.Infof("Deployed %d global commands", len(defs))
}
