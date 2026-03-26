package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// RoleMenuRepository handles role menu operations
type RoleMenuRepository struct {
	db *gorm.DB
}

// NewRoleMenuRepository creates a new role menu repository
func NewRoleMenuRepository(db *gorm.DB) *RoleMenuRepository {
	return &RoleMenuRepository{db: db}
}

// FindByID finds a role menu by ID
func (r *RoleMenuRepository) FindByID(id int) (*models.RoleMenu, error) {
	var menu models.RoleMenu
	if err := r.db.Preload("Options").First(&menu, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

// FindByGuild finds role menus in a guild
func (r *RoleMenuRepository) FindByGuild(guildID string) ([]models.RoleMenu, error) {
	var menus []models.RoleMenu
	if err := r.db.Preload("Options").Where("guild_id = ?", guildID).Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

// Create creates a new role menu
func (r *RoleMenuRepository) Create(menu *models.RoleMenu) error {
	return r.db.Create(menu).Error
}

// Update updates a role menu
func (r *RoleMenuRepository) Update(menu *models.RoleMenu) error {
	return r.db.Save(menu).Error
}

// Delete deletes a role menu
func (r *RoleMenuRepository) Delete(id int) error {
	return r.db.Delete(&models.RoleMenu{}, "id = ?", id).Error
}

// CreateOption creates a role menu option
func (r *RoleMenuRepository) CreateOption(option *models.RoleMenuOption) error {
	return r.db.Create(option).Error
}

// DeleteOption deletes a role menu option
func (r *RoleMenuRepository) DeleteOption(id int) error {
	return r.db.Delete(&models.RoleMenuOption{}, "id = ?", id).Error
}
