package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// GameBiasRepository handles game bias operations
type GameBiasRepository struct {
	db *gorm.DB
}

// NewGameBiasRepository creates a new game bias repository
func NewGameBiasRepository(db *gorm.DB) *GameBiasRepository {
	return &GameBiasRepository{db: db}
}

// FindByID finds a game bias by ID
func (r *GameBiasRepository) FindByID(id int) (*models.GameBias, error) {
	var bias models.GameBias
	if err := r.db.First(&bias, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bias, nil
}

// FindByGuildAndUserAndGame finds game bias for a user and game type
func (r *GameBiasRepository) FindByGuildAndUserAndGame(guildID string, userID int, gameType string) (*models.GameBias, error) {
	var bias models.GameBias
	if err := r.db.Where("guild_id = ? AND user_id = ? AND game_type = ?", guildID, userID, gameType).First(&bias).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &bias, nil
}

// Create creates a new game bias record
func (r *GameBiasRepository) Create(bias *models.GameBias) error {
	return r.db.Create(bias).Error
}

// Update updates a game bias record
func (r *GameBiasRepository) Update(bias *models.GameBias) error {
	return r.db.Save(bias).Error
}

// AddLoss increments the loss count for a game
func (r *GameBiasRepository) AddLoss(guildID string, userID int, gameType string) error {
	return r.db.Model(&models.GameBias{}).
		Where("guild_id = ? AND user_id = ? AND game_type = ?", guildID, userID, gameType).
		Update("loss_count", gorm.Expr("loss_count + ?", 1)).Error
}
