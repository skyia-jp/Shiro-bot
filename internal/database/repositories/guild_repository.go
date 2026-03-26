package repositories

import (
	"errors"

	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// GuildRepository handles guild operations
type GuildRepository struct {
	db *gorm.DB
}

// NewGuildRepository creates a new guild repository
func NewGuildRepository(db *gorm.DB) *GuildRepository {
	return &GuildRepository{db: db}
}

// FindByID finds a guild by ID
func (r *GuildRepository) FindByID(guildID string) (*models.Guild, error) {
	var guild models.Guild
	if err := r.db.Preload("Config").First(&guild, "id = ?", guildID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &guild, nil
}

// Create creates a new guild
func (r *GuildRepository) Create(guild *models.Guild) error {
	return r.db.Create(guild).Error
}

// Update updates a guild
func (r *GuildRepository) Update(guild *models.Guild) error {
	return r.db.Save(guild).Error
}

// Delete deletes a guild
func (r *GuildRepository) Delete(guildID string) error {
	return r.db.Delete(&models.Guild{}, "id = ?", guildID).Error
}

// GetOrCreateGuild gets or creates a guild
func (r *GuildRepository) GetOrCreateGuild(guildID string, guildName string) (*models.Guild, error) {
	guild, err := r.FindByID(guildID)
	if err != nil {
		return nil, err
	}

	if guild != nil {
		return guild, nil
	}

	// Create new guild
	newGuild := &models.Guild{
		ID:   guildID,
		Name: &guildName,
	}

	if err := r.Create(newGuild); err != nil {
		return nil, err
	}

	// Create default config
	config := &models.GuildConfig{
		GuildID:  guildID,
		Timezone: stringPtr("Asia/Tokyo"),
	}

	if err := r.db.Create(config).Error; err != nil {
		return nil, err
	}

	newGuild.Config = config
	return newGuild, nil
}

// stringPtr returns a pointer to a string
func stringPtr(s string) *string {
	return &s
}
