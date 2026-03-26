package repositories

import (
	"github.com/skyia-jp/shiro-go/internal/models"
	"gorm.io/gorm"
)

// CurrencyRepository handles currency and balance operations
type CurrencyRepository struct {
	db *gorm.DB
}

// NewCurrencyRepository creates a new currency repository
func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

// GetBalance gets the currency balance for a user in a guild
func (r *CurrencyRepository) GetBalance(guildID string, userID int) (*models.CurrencyBalance, error) {
	var balance models.CurrencyBalance
	if err := r.db.Where("guild_id = ? AND user_id = ?", guildID, userID).First(&balance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &balance, nil
}

// CreateBalance creates a new currency balance
func (r *CurrencyRepository) CreateBalance(balance *models.CurrencyBalance) error {
	return r.db.Create(balance).Error
}

// UpdateBalance updates a currency balance
func (r *CurrencyRepository) UpdateBalance(balance *models.CurrencyBalance) error {
	return r.db.Save(balance).Error
}

// AddTransaction adds a currency transaction
func (r *CurrencyRepository) AddTransaction(transaction *models.CurrencyTransaction) error {
	return r.db.Create(transaction).Error
}

// GetTransactions gets transactions for a user
func (r *CurrencyRepository) GetTransactions(guildID string, userID int, limit int) ([]models.CurrencyTransaction, error) {
	var transactions []models.CurrencyTransaction
	query := r.db.Where("guild_id = ? AND user_id = ?", guildID, userID).Order("created_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
