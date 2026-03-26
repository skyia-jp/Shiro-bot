package events

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"

	"github.com/skyia-jp/shiro-go/internal/database/repositories"
)

// MessageCreate handles the messageCreate event
func MessageCreate(log *zap.SugaredLogger, repos *repositories.ModuleRepository) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if s == nil || m == nil || m.Author == nil {
			return
		}

		if s.State == nil || s.State.User == nil {
			return
		}

		// Ignore bot messages
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Ignore DMs (no guild ID)
		if m.GuildID == "" {
			return
		}

		log.Debugf("Message from %s#%s in guild %s: %s", m.Author.Username, m.Author.Discriminator, m.GuildID, m.Content)

		// TODO: Handle message events
		// - Process commands if prefix-based
		// - Track message activity
		// - Handle filters/moderation
	}
}
