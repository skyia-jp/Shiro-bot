package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// ===== Main Models =====

// Guild represents a Discord guild
type Guild struct {
	ID        string    `gorm:"primaryKey"`
	Name      *string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	Config               *GuildConfig
	Reminders            []Reminder
	RoleMenus            []RoleMenu
	Warnings             []Warning
	ModerationLog        []ModerationAction
	PinnedMessages       []PinnedMessage
	Tickets              []Ticket
	Members              []GuildUser
	Polls                []Poll
	Notes                []Note
	Tasks                []Task
	ActivityRecords      []ActivityRecord
	ScheduledJobs        []ScheduledJob
	GameBiases           []GameBias
	CurrencyBalances     []CurrencyBalance
	CurrencyTransactions []CurrencyTransaction
	ShopItems            []ShopItem
}

// GuildConfig represents guild configuration
type GuildConfig struct {
	ID                     int       `gorm:"primaryKey"`
	GuildID                string    `gorm:"uniqueIndex"`
	LogChannelID           *string
	MuteRoleID             *string
	Timezone               *string   `gorm:"default:Asia/Tokyo"`
	ReminderRoleID         *string
	AutoRoleID             *string
	PasswordAuthRoleID     *string
	PasswordAuthSecretHash *string
	PasswordAuthHint       *string
	PasswordAuthUpdatedAt  *time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
}

// PasswordReveal represents password reveal panel
type PasswordReveal struct {
	ID          int        `gorm:"primaryKey"`
	CustomID    string     `gorm:"uniqueIndex"`
	GuildID     string
	ChannelID   string
	CreatedByID string
	Password    *string
	Title       *string
	Description *string
	ButtonLabel *string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	Mode        string `gorm:"default:DISPLAY"` // DISPLAY or INPUT
}

// User represents a Discord user
type User struct {
	ID       int    `gorm:"primaryKey"`
	DiscordID string `gorm:"uniqueIndex"`
	Username *string

	// Relations
	GuildUsers           []GuildUser
	Reminders            []Reminder
	WarningsReceived     []Warning `gorm:"foreignKey:UserID;references:ID"`
	WarningsIssued       []Warning `gorm:"foreignKey:ModeratorID;references:ID"`
	ActionsReceived      []ModerationAction `gorm:"foreignKey:UserID;references:ID"`
	ActionsIssued        []ModerationAction `gorm:"foreignKey:ModeratorID;references:ID"`
	CreatedRoleMenus     []RoleMenu `gorm:"foreignKey:CreatedByID;references:ID"`
	PinnedMessages       []PinnedMessage `gorm:"foreignKey:PinnedByID;references:ID"`
	Tickets              []Ticket `gorm:"foreignKey:UserID;references:ID"`
	PollsCreated         []Poll `gorm:"foreignKey:CreatedByID;references:ID"`
	PollVotes            []PollVote
	Notes                []Note `gorm:"foreignKey:AuthorID;references:ID"`
	TasksCreated         []Task `gorm:"foreignKey:CreatorID;references:ID"`
	TasksAssigned        []Task `gorm:"foreignKey:AssigneeID;references:ID"`
	GameBiases           []GameBias
	CurrencyBalances     []CurrencyBalance
	CurrencyTransactions []CurrencyTransaction
}

// GuildUser represents a guild member
type GuildUser struct {
	ID       int       `gorm:"primaryKey"`
	GuildID  string    `gorm:"index"`
	UserID   int
	JoinedAt time.Time
	IsMuted  bool      `gorm:"default:false"`

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
	User  User  `gorm:"foreignKey:UserID;references:ID"`
}

// Reminder represents a reminder
type Reminder struct {
	ID             int       `gorm:"primaryKey"`
	GuildID        string    `gorm:"index"`
	UserID         int
	ChannelID      string
	Message        string
	CronExpression string
	NextTriggerAt  *time.Time
	Timezone       string `gorm:"default:Asia/Tokyo"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
	User  User  `gorm:"foreignKey:UserID;references:ID"`
}

// RoleMenu represents a role menu
type RoleMenu struct {
	ID            int       `gorm:"primaryKey"`
	GuildID       string    `gorm:"index"`
	ChannelID     string
	MessageID     *string
	Title         string
	Description   *string
	MaxSelectable int       `gorm:"default:1"`
	CreatedByID   int
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Relations
	Guild   Guild           `gorm:"foreignKey:GuildID;references:ID"`
	CreatedBy User           `gorm:"foreignKey:CreatedByID;references:ID"`
	Options []RoleMenuOption
}

// RoleMenuOption represents a role menu option
type RoleMenuOption struct {
	ID          int    `gorm:"primaryKey"`
	RoleMenuID  int    `gorm:"index"`
	RoleID      string
	Label       string
	Description *string
	Emoji       *string
	IsDefault   bool `gorm:"default:false"`

	// Relations
	RoleMenu RoleMenu `gorm:"foreignKey:RoleMenuID;references:ID"`
}

// Warning represents a user warning
type Warning struct {
	ID          int       `gorm:"primaryKey"`
	GuildID     string    `gorm:"index"`
	UserID      int
	ModeratorID int
	Reason      string
	PenaltyLevel int      `gorm:"default:1"`
	CreatedAt   time.Time

	// Relations
	Guild     Guild `gorm:"foreignKey:GuildID;references:ID"`
	User      User  `gorm:"foreignKey:UserID;references:ID"`
	Moderator User  `gorm:"foreignKey:ModeratorID;references:ID"`
}

// ModerationAction represents a moderation action
type ModerationAction struct {
	ID          int       `gorm:"primaryKey"`
	GuildID     string    `gorm:"index"`
	UserID      int
	ModeratorID int
	ActionType  string
	Reason      *string
	ExpiresAt   *time.Time
	CreatedAt   time.Time

	// Relations
	Guild     Guild `gorm:"foreignKey:GuildID;references:ID"`
	User      User  `gorm:"foreignKey:UserID;references:ID"`
	Moderator User  `gorm:"foreignKey:ModeratorID;references:ID"`
}

// PinnedMessage represents a pinned message
type PinnedMessage struct {
	ID              int       `gorm:"primaryKey"`
	GuildID         string    `gorm:"index"`
	ChannelID       string
	MessageID       string
	PinnedByID      int
	ExpiresAt       *time.Time
	SourceMessageID *string
	CloneMessageID  *string
	SnapshotContent *string `gorm:"type:longtext"`
	SnapshotEmbeds  json.RawMessage `gorm:"type:json"`
	SnapshotFiles   json.RawMessage `gorm:"type:json"`
	CreatedAt       time.Time

	// Relations
	Guild   Guild `gorm:"foreignKey:GuildID;references:ID"`
	PinnedBy User  `gorm:"foreignKey:PinnedByID;references:ID"`
}

// Poll represents a poll
type Poll struct {
	ID        int       `gorm:"primaryKey"`
	GuildID   string    `gorm:"index"`
	ChannelID string
	MessageID *string
	Question  string
	Status    string `gorm:"default:open"`
	CreatedByID int
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Relations
	Guild     Guild        `gorm:"foreignKey:GuildID;references:ID"`
	CreatedBy User         `gorm:"foreignKey:CreatedByID;references:ID"`
	Options   []PollOption
	Votes     []PollVote
}

// PollOption represents a poll option
type PollOption struct {
	ID       int    `gorm:"primaryKey"`
	PollID   int    `gorm:"index"`
	Label    string
	Emoji    *string
	Position int

	// Relations
	Poll  Poll       `gorm:"foreignKey:PollID;references:ID"`
	Votes []PollVote `gorm:"foreignKey:OptionID;references:ID"`
}

// PollVote represents a poll vote
type PollVote struct {
	ID       int       `gorm:"primaryKey"`
	PollID   int       `gorm:"index"`
	OptionID int
	UserID   int
	CreatedAt time.Time

	// Relations
	Poll   Poll       `gorm:"foreignKey:PollID;references:ID"`
	Option PollOption `gorm:"foreignKey:OptionID;references:ID"`
	User   User       `gorm:"foreignKey:UserID;references:ID"`
}

// Note represents a note
type Note struct {
	ID        int       `gorm:"primaryKey"`
	GuildID   string    `gorm:"index"`
	AuthorID  int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	Guild  Guild `gorm:"foreignKey:GuildID;references:ID"`
	Author User  `gorm:"foreignKey:AuthorID;references:ID"`
}

// Task represents a task
type Task struct {
	ID          int       `gorm:"primaryKey"`
	GuildID     string    `gorm:"index"`
	CreatorID   int
	AssigneeID  *int
	Description string
	Status      string `gorm:"default:open"`
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Relations
	Guild    Guild `gorm:"foreignKey:GuildID;references:ID"`
	Creator  User  `gorm:"foreignKey:CreatorID;references:ID"`
	Assignee *User `gorm:"foreignKey:AssigneeID;references:ID"`
}

// ActivityRecord represents user activity
type ActivityRecord struct {
	ID            int       `gorm:"primaryKey"`
	GuildID       string    `gorm:"index:idx_guild_user_date"`
	DiscordUserID string    `gorm:"index:idx_guild_user_date"`
	Date          time.Time `gorm:"index:idx_guild_user_date"`
	MessageCount  int       `gorm:"default:0"`
	VoiceMinutes  int       `gorm:"default:0"`
	LastUpdated   time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
}

// ScheduledJob represents a scheduled job
type ScheduledJob struct {
	ID        int       `gorm:"primaryKey"`
	GuildID   string    `gorm:"index"`
	Type      string
	Schedule  string
	Data      json.RawMessage `gorm:"type:json"`
	LastRun   *time.Time
	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
}

// Ticket represents a support ticket
type Ticket struct {
	ID        int       `gorm:"primaryKey"`
	GuildID   string    `gorm:"index"`
	UserID    int
	ChannelID string
	Status    string `gorm:"default:open"`
	CreatedAt time.Time
	ClosedAt  *time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
	User  User  `gorm:"foreignKey:UserID;references:ID"`
}

// GameBias represents game history for bias tracking
type GameBias struct {
	ID        int       `gorm:"primaryKey"`
	GuildID   string    `gorm:"index"`
	UserID    int
	GameType  string
	LossCount int       `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
	User  User  `gorm:"foreignKey:UserID;references:ID"`
}

// CurrencyBalance represents user currency balance
type CurrencyBalance struct {
	ID        int       `gorm:"primaryKey"`
	GuildID   string    `gorm:"index"`
	UserID    int       `gorm:"index"`
	Balance   int       `gorm:"default:0"`
	UpdatedAt time.Time

	// Relations
	Guild        Guild                  `gorm:"foreignKey:GuildID;references:ID"`
	User         User                   `gorm:"foreignKey:UserID;references:ID"`
	Transactions []CurrencyTransaction
}

// CurrencyTransaction represents a currency transaction
type CurrencyTransaction struct {
	ID           int            `gorm:"primaryKey"`
	GuildID      string         `gorm:"index"`
	UserID       int            `gorm:"index"`
	BalanceID    *int
	Type         string         // EARN, SPEND, TRANSFER_IN, TRANSFER_OUT, ADJUST, GAME_BET, GAME_WIN, DAILY_BONUS
	Amount       int
	BalanceAfter int
	Reason       *string
	Metadata     json.RawMessage `gorm:"type:json"`
	CreatedAt    time.Time

	// Relations
	Guild   Guild            `gorm:"foreignKey:GuildID;references:ID"`
	User    User             `gorm:"foreignKey:UserID;references:ID"`
	Balance *CurrencyBalance `gorm:"foreignKey:BalanceID;references:ID"`
}

// ShopItem represents a shop item
type ShopItem struct {
	ID          string `gorm:"primaryKey"`
	GuildID     string `gorm:"index"`
	Name        string
	Description *string
	Price       int
	RoleID      *string
	CreatedBy   *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Relations
	Guild Guild `gorm:"foreignKey:GuildID;references:ID"`
}

// ===== Query Helper Functions =====

// AfterFind is a GORM hook for post-fetch processing
func (u *User) AfterFind(tx *gorm.DB) error {
	// Additional processing if needed
	return nil
}
