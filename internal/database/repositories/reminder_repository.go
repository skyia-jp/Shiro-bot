package repositories

import (
	"time"

	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// ReminderRepository handles reminder operations
type ReminderRepository struct {
	db *gorm.DB
}

// NewReminderRepository creates a new reminder repository
func NewReminderRepository(db *gorm.DB) *ReminderRepository {
	return &ReminderRepository{db: db}
}

// FindByID finds a reminder by ID
func (r *ReminderRepository) FindByID(id int) (*models.Reminder, error) {
	var reminder models.Reminder
	if err := r.db.First(&reminder, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &reminder, nil
}

// FindByGuildAndUser finds reminders for a guild user
func (r *ReminderRepository) FindByGuildAndUser(guildID string, userID int) ([]models.Reminder, error) {
	var reminders []models.Reminder
	if err := r.db.Where("guild_id = ? AND user_id = ?", guildID, userID).Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

// FindDueReminders finds reminders that are due
func (r *ReminderRepository) FindDueReminders() ([]models.Reminder, error) {
	var reminders []models.Reminder
	now := time.Now().UTC()
	if err := r.db.Where("next_trigger_at IS NOT NULL AND next_trigger_at <= ?", now).Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

// Create creates a new reminder
func (r *ReminderRepository) Create(reminder *models.Reminder) error {
	return r.db.Create(reminder).Error
}

// Update updates a reminder
func (r *ReminderRepository) Update(reminder *models.Reminder) error {
	return r.db.Save(reminder).Error
}

// Delete deletes a reminder
func (r *ReminderRepository) Delete(id int) error {
	return r.db.Delete(&models.Reminder{}, "id = ?", id).Error
}
