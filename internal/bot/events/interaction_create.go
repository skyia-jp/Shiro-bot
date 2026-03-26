package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/skyia-jp/shiro-go/internal/commands"
	"go.uber.org/zap"
)

// InteractionCreate handles the interactionCreate event
func InteractionCreate(log *zap.SugaredLogger, router *commands.Router) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i == nil || s == nil {
			return
		}

		log.Debugf("Interaction received (Type: %d, ID: %s)", i.Type, i.ID)

		// Route interactions to appropriate handlers based on type
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if router == nil {
				log.Warn("Command router is nil")
				return
			}
			if err := router.HandleInteraction(s, i); err != nil {
				log.Errorf("Failed to handle command interaction: %v", err)
			}

		case discordgo.InteractionMessageComponent:
			// Handle button/select menu interaction
			customID := i.MessageComponentData().CustomID
			log.Debugf("Component interaction: %s", customID)
			// TODO: Route to interaction handler

		case discordgo.InteractionModalSubmit:
			// Handle modal submission
			customID := i.ModalSubmitData().CustomID
			log.Debugf("Modal submitted: %s", customID)
			// TODO: Route to modal handler

		default:
			log.Debugf("Unknown interaction type: %d", i.Type)
		}
	}
}
