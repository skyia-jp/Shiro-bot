package repositories

import (
	"errors"

	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// UserRepository handles user operations
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(userID int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByDiscordID finds a user by Discord ID
func (r *UserRepository) FindByDiscordID(discordID string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "discord_id = ?", discordID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// GetOrCreateUser gets or creates a user
func (r *UserRepository) GetOrCreateUser(discordID string, username *string) (*models.User, error) {
	user, err := r.FindByDiscordID(discordID)
	if err != nil {
		return nil, err
	}

	if user != nil {
		// Update username if provided and different
		if username != nil && (user.Username == nil || *user.Username != *username) {
			user.Username = username
			if err := r.Update(user); err != nil {
				return nil, err
			}
		}
		return user, nil
	}

	// Create new user
	newUser := &models.User{
		DiscordID: discordID,
		Username:  username,
	}

	if err := r.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
