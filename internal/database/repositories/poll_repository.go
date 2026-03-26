package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// PollRepository handles poll operations
type PollRepository struct {
	db *gorm.DB
}

// NewPollRepository creates a new poll repository
func NewPollRepository(db *gorm.DB) *PollRepository {
	return &PollRepository{db: db}
}

// FindByID finds a poll by ID
func (r *PollRepository) FindByID(id int) (*models.Poll, error) {
	var poll models.Poll
	if err := r.db.Preload("Options").Preload("Votes").First(&poll, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &poll, nil
}

// Create creates a new poll
func (r *PollRepository) Create(poll *models.Poll) error {
	return r.db.Create(poll).Error
}

// Update updates a poll
func (r *PollRepository) Update(poll *models.Poll) error {
	return r.db.Save(poll).Error
}

// Delete deletes a poll
func (r *PollRepository) Delete(id int) error {
	return r.db.Delete(&models.Poll{}, "id = ?", id).Error
}

// CreateOption creates a poll option
func (r *PollRepository) CreateOption(option *models.PollOption) error {
	return r.db.Create(option).Error
}

// CreateVote creates a poll vote
func (r *PollRepository) CreateVote(vote *models.PollVote) error {
	return r.db.Create(vote).Error
}
