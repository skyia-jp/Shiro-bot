package repositories

import (
	"time"

	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// ActivityRepository handles activity record operations
type ActivityRepository struct {
	db *gorm.DB
}

// NewActivityRepository creates a new activity repository
func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

// FindByGuildAndUserAndDate finds activity record for a user on a specific date
func (r *ActivityRepository) FindByGuildAndUserAndDate(guildID, discordUserID string, date time.Time) (*models.ActivityRecord, error) {
	var record models.ActivityRecord
	// Truncate date to midnight UTC
	date = date.UTC().Truncate(24 * time.Hour)
	
	if err := r.db.Where("guild_id = ? AND discord_user_id = ? AND date = ?", guildID, discordUserID, date).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

// Create creates a new activity record
func (r *ActivityRepository) Create(record *models.ActivityRecord) error {
	return r.db.Create(record).Error
}

// Update updates an activity record
func (r *ActivityRepository) Update(record *models.ActivityRecord) error {
	return r.db.Save(record).Error
}

// GetUserActivity gets activity records for a user
func (r *ActivityRepository) GetUserActivity(guildID, discordUserID string, limit int) ([]models.ActivityRecord, error) {
	var records []models.ActivityRecord
	query := r.db.Where("guild_id = ? AND discord_user_id = ?", guildID, discordUserID).Order("date DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// GetGuildActivity gets activity records for a guild on a specific date
func (r *ActivityRepository) GetGuildActivity(guildID string, date time.Time) ([]models.ActivityRecord, error) {
	var records []models.ActivityRecord
	date = date.UTC().Truncate(24 * time.Hour)
	
	if err := r.db.Where("guild_id = ? AND date = ?", guildID, date).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
