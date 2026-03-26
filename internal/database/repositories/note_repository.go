package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// NoteRepository handles note operations
type NoteRepository struct {
	db *gorm.DB
}

// NewNoteRepository creates a new note repository
func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

// FindByID finds a note by ID
func (r *NoteRepository) FindByID(id int) (*models.Note, error) {
	var note models.Note
	if err := r.db.Preload("Guild").Preload("Author").First(&note, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

// FindByGuild finds notes in a guild
func (r *NoteRepository) FindByGuild(guildID string) ([]models.Note, error) {
	var notes []models.Note
	if err := r.db.Preload("Author").Where("guild_id = ?", guildID).Order("created_at DESC").Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

// Create creates a new note
func (r *NoteRepository) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

// Update updates a note
func (r *NoteRepository) Update(note *models.Note) error {
	return r.db.Save(note).Error
}

// Delete deletes a note
func (r *NoteRepository) Delete(id int) error {
	return r.db.Delete(&models.Note{}, "id = ?", id).Error
}
