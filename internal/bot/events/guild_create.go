package events

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"

	"github.com/skyia-jp/shiro-go/internal/database/repositories"
)

// GuildCreate handles the guildCreate event
func GuildCreate(log *zap.SugaredLogger, guildRepo *repositories.GuildRepository) func(*discordgo.Session, *discordgo.GuildCreate) {
	return func(s *discordgo.Session, g *discordgo.GuildCreate) {
		log.Infof("Guild joined: %s (ID: %s) with %d members", g.Guild.Name, g.Guild.ID, g.MemberCount)

		// Get or create guild in database
		_, err := guildRepo.GetOrCreateGuild(g.Guild.ID, g.Guild.Name)
		if err != nil {
			log.Errorf("Failed to create guild record: %v", err)
			return
		}

		log.Debugf("Guild record created/updated for %s", g.Guild.ID)
	}
}
