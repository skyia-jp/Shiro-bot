package utility

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/skyia-jp/shiro-go/internal/commands"
)

// RegisterBuiltins registers low-risk baseline commands.
func RegisterBuiltins(router *commands.Router) error {
	if err := router.Register(pingCommand()); err != nil {
		return err
	}
	if err := router.Register(helpCommand(router)); err != nil {
		return err
	}
	if err := router.Register(uptimeCommand()); err != nil {
		return err
	}
	return nil
}

func pingCommand() commands.Command {
	return commands.Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Check bot latency",
		},
		Handler: func(ctx *commands.Context) error {
			handlerStartedAt := time.Now()
			now := handlerStartedAt
			latency := ctx.Session.HeartbeatLatency().Milliseconds()

			receiveMs := int64(0)
			interactionID := ""
			if ctx.Interaction != nil && ctx.Interaction.Interaction != nil {
				interactionID = ctx.Interaction.Interaction.ID
			}

			if interactionID != "" {
				if ts, err := discordgo.SnowflakeTimestamp(interactionID); err == nil {
					receiveMs = now.Sub(ts).Milliseconds()
					if receiveMs < 0 {
						receiveMs = 0
					}
				}
			}

			// Fallback for environments where interaction snowflake parsing is unavailable.
			if receiveMs == 0 {
				receiveMs = time.Since(handlerStartedAt).Milliseconds()
				if receiveMs == 0 {
					receiveMs = 1
				}
			}

			embed := &discordgo.MessageEmbed{
				Title: "Pong!",
				Color: 0x22C55E,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "WebSocket",
						Value:  fmt.Sprintf("`%d ms`", latency),
						Inline: true,
					},
					{
						Name:   "コマンド受信",
						Value:  fmt.Sprintf("`%d ms`", receiveMs),
						Inline: true,
					},
				},
			}

			return ctx.Session.InteractionRespond(ctx.Interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}},
			})
		},
	}
}

func helpCommand(router *commands.Router) commands.Command {
	return commands.Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "help",
			Description: "Show available commands",
		},
		Handler: func(ctx *commands.Context) error {
			defs := router.Definitions()
			lines := make([]string, 0, len(defs))
			for _, def := range defs {
				lines = append(lines, fmt.Sprintf("/%s - %s", def.Name, def.Description))
			}
			content := "Available commands:\n" + strings.Join(lines, "\n")
			return ctx.Reply(content)
		},
	}
}

func uptimeCommand() commands.Command {
	startedAt := time.Now()
	return commands.Command{
		Definition: &discordgo.ApplicationCommand{
			Name:        "uptime",
			Description: "Show bot uptime",
		},
		Handler: func(ctx *commands.Context) error {
			d := time.Since(startedAt).Round(time.Second)
			return ctx.Reply(fmt.Sprintf("Uptime: %s", d))
		},
	}
}
