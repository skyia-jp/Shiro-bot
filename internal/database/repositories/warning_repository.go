package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// WarningRepository handles warning operations
type WarningRepository struct {
	db *gorm.DB
}

// NewWarningRepository creates a new warning repository
func NewWarningRepository(db *gorm.DB) *WarningRepository {
	return &WarningRepository{db: db}
}

// FindByID finds a warning by ID
func (r *WarningRepository) FindByID(id int) (*models.Warning, error) {
	var warning models.Warning
	if err := r.db.Preload("User").Preload("Moderator").First(&warning, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &warning, nil
}

// FindByGuildAndUser finds warnings for a user in a guild
func (r *WarningRepository) FindByGuildAndUser(guildID string, userID int) ([]models.Warning, error) {
	var warnings []models.Warning
	if err := r.db.Where("guild_id = ? AND user_id = ?", guildID, userID).Order("created_at DESC").Find(&warnings).Error; err != nil {
		return nil, err
	}
	return warnings, nil
}

// CountByGuildAndUser counts warnings for a user in a guild
func (r *WarningRepository) CountByGuildAndUser(guildID string, userID int) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Warning{}).Where("guild_id = ? AND user_id = ?", guildID, userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Create creates a new warning
func (r *WarningRepository) Create(warning *models.Warning) error {
	return r.db.Create(warning).Error
}

// Delete deletes a warning
func (r *WarningRepository) Delete(id int) error {
	return r.db.Delete(&models.Warning{}, "id = ?", id).Error
}
