package commands

import "github.com/bwmarrin/discordgo"

// Reply sends a normal interaction response.
func (c *Context) Reply(message string) error {
	return c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: message},
	})
}

// ReplyEphemeral sends an ephemeral interaction response.
func (c *Context) ReplyEphemeral(message string) error {
	return c.Session.InteractionRespond(c.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
