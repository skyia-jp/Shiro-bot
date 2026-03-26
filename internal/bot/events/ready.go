package events

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// Ready handles the ready event
func Ready(log *zap.SugaredLogger) func(*discordgo.Session, *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		if s == nil || s.State == nil || s.State.User == nil {
			log.Warn("Ready event received with incomplete session state")
			return
		}

		log.Infof("Bot logged in as %v", s.State.User)

		// Set bot status
		status := "v1.5.1 | /help for help"
		err := s.UpdateGameStatus(0, status)
		if err != nil {
			log.Errorf("Failed to set bot status: %v", err)
		}
	}
}
