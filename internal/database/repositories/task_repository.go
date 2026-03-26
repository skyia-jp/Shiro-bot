package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// TaskRepository handles task operations
type TaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new task repository
func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// FindByID finds a task by ID
func (r *TaskRepository) FindByID(id int) (*models.Task, error) {
	var task models.Task
	if err := r.db.Preload("Creator").Preload("Assignee").First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// FindByGuild finds tasks in a guild
func (r *TaskRepository) FindByGuild(guildID string) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("guild_id = ?", guildID).Order("created_at DESC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Create creates a new task
func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// Update updates a task
func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete deletes a task
func (r *TaskRepository) Delete(id int) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}
