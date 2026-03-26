package events

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"

	"github.com/skyia-jp/shiro-go/internal/database/repositories"
)

// VoiceStateUpdate handles the voiceStateUpdate event
func VoiceStateUpdate(log *zap.SugaredLogger, repos *repositories.ModuleRepository) func(*discordgo.Session, *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
		if v.ChannelID != "" {
			log.Debugf("User %s joined voice channel %s in guild %s", v.UserID, v.ChannelID, v.GuildID)
		} else {
			log.Debugf("User %s left voice channel in guild %s", v.UserID, v.GuildID)
		}

		// TODO: Handle voice state changes
		// - Track user voice status
		// - Update activity records
		// - Handle room-related features
	}
}
