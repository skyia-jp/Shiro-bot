package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/skyia-jp/shiro-go/internal/commands"
	"github.com/skyia-jp/shiro-go/internal/commands/utility"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/skyia-jp/shiro-go/internal/bot/events"
	"github.com/skyia-jp/shiro-go/internal/config"
)

// Client represents the Discord bot client
type Client struct {
	session  *discordgo.Session
	db       *gorm.DB
	log      *zap.SugaredLogger
	config   *config.Config
	router   *commands.Router
}

// NewClient creates a new Discord bot client
func NewClient(cfg *config.Config, db *gorm.DB, log *zap.SugaredLogger) (*Client, error) {
	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Create Discord session
	session, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord session: %w", err)
	}

	router := commands.NewRouter(log, nil)
	if err := utility.RegisterBuiltins(router); err != nil {
		return nil, fmt.Errorf("failed to register built-in commands: %w", err)
	}

	client := &Client{
		session: session,
		db:      db,
		log:     log,
		config:  cfg,
		router:  router,
	}

	// Register event handlers
	client.registerEventHandlers()

	return client, nil
}

// registerEventHandlers registers event handlers for Discord events
func (c *Client) registerEventHandlers() {
	c.session.AddHandler(events.Ready(c.log))
	c.session.AddHandler(events.InteractionCreate(c.log, c.router))
}

// Start starts the Discord bot
func (c *Client) Start() error {
	// Minimal intent set for slash-command focused bot.
	c.session.Identify.Intents = discordgo.IntentsGuilds

	if err := c.session.Open(); err != nil {
		return fmt.Errorf("failed to open Discord session: %w", err)
	}

	return nil
}

// Close closes the Discord session and database connection
func (c *Client) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

// GetSession returns the Discord session
func (c *Client) GetSession() *discordgo.Session {
	return c.session
}

// GetDB returns the database connection
func (c *Client) GetDB() *gorm.DB {
	return c.db
}

// GetLogger returns the logger
func (c *Client) GetLogger() *zap.SugaredLogger {
	return c.log
}

// GetConfig returns the configuration
func (c *Client) GetConfig() *config.Config {
	return c.config
}

