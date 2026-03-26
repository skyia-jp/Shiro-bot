package events

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"

	"github.com/skyia-jp/shiro-go/internal/database/repositories"
)

// GuildMemberAdd handles the guildMemberAdd event
func GuildMemberAdd(log *zap.SugaredLogger, guildRepo *repositories.GuildRepository, userRepo *repositories.UserRepository) func(*discordgo.Session, *discordgo.GuildMemberAdd) {
	return func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		if m == nil || m.User == nil {
			return
		}

		log.Debugf("Member joined: %s#%s (ID: %s) in guild %s", m.User.Username, m.User.Discriminator, m.User.ID, m.GuildID)

		if m.GuildID != "" {
			if _, err := guildRepo.GetOrCreateGuild(m.GuildID, ""); err != nil {
				log.Errorf("Failed to ensure guild record on member join: %v", err)
			}
		}

		// Get or create user in database
		user, err := userRepo.GetOrCreateUser(m.User.ID, &m.User.Username)
		if err != nil {
			log.Errorf("Failed to create user record: %v", err)
			return
		}

		// TODO: Handle member join events
		// - Track member in GuildUser
		// - Apply auto-roles if configured
		// - Send welcome message if required
		if user != nil {
			log.Debugf("User record ensured for %s", m.User.ID)
		}
	}
}
