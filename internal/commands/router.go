package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/skyia-jp/shiro-go/internal/database/repositories"
	"go.uber.org/zap"
)

// HandlerFunc handles a slash command.
type HandlerFunc func(*Context) error

// Command defines a slash command and its handler.
type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    HandlerFunc
}

// Context is passed to command handlers.
type Context struct {
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
	Repos       *repositories.ModuleRepository
	Log         *zap.SugaredLogger
}

// Router routes slash commands.
type Router struct {
	log      *zap.SugaredLogger
	repos     *repositories.ModuleRepository
	commands map[string]Command
}

// NewRouter creates a new command router.
func NewRouter(log *zap.SugaredLogger, repos *repositories.ModuleRepository) *Router {
	return &Router{
		log:      log,
		repos:     repos,
		commands: make(map[string]Command),
	}
}

// Register registers one slash command.
func (r *Router) Register(cmd Command) error {
	if cmd.Definition == nil {
		return fmt.Errorf("command definition is nil")
	}
	if strings.TrimSpace(cmd.Definition.Name) == "" {
		return fmt.Errorf("command name is empty")
	}
	if cmd.Handler == nil {
		return fmt.Errorf("handler is nil for command %s", cmd.Definition.Name)
	}

	name := strings.ToLower(cmd.Definition.Name)
	if _, exists := r.commands[name]; exists {
		return fmt.Errorf("command already registered: %s", name)
	}

	r.commands[name] = cmd
	return nil
}

// Definitions returns all command definitions sorted by name.
func (r *Router) Definitions() []*discordgo.ApplicationCommand {
	names := make([]string, 0, len(r.commands))
	for name := range r.commands {
		names = append(names, name)
	}
	sort.Strings(names)

	definitions := make([]*discordgo.ApplicationCommand, 0, len(names))
	for _, name := range names {
		definitions = append(definitions, r.commands[name].Definition)
	}
	return definitions
}

// HandleInteraction routes slash command interactions.
func (r *Router) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	if s == nil || i == nil {
		return nil
	}
	if i.Type != discordgo.InteractionApplicationCommand {
		return nil
	}

	data := i.ApplicationCommandData()
	name := strings.ToLower(data.Name)
	cmd, ok := r.commands[name]
	if !ok {
		return respondEphemeral(s, i, fmt.Sprintf("Unknown command: /%s", data.Name))
	}

	ctx := &Context{
		Session:     s,
		Interaction: i,
		Repos:       r.repos,
		Log:         r.log,
	}

	if err := cmd.Handler(ctx); err != nil {
		r.log.Errorf("Command /%s failed: %v", name, err)
		return respondEphemeral(s, i, "Command failed. Please try again later.")
	}

	return nil
}

func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, message string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
