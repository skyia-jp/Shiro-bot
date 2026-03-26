package repositories

import (
	"gorm.io/gorm"
)

// ModuleRepository centralizes access to all repositories
type ModuleRepository struct {
	Guild      *GuildRepository
	User       *UserRepository
	Currency   *CurrencyRepository
	Warning    *WarningRepository
	Poll       *PollRepository
	Task       *TaskRepository
	Note       *NoteRepository
	RoleMenu   *RoleMenuRepository
	Activity   *ActivityRepository
	GameBias   *GameBiasRepository
}

// NewModuleRepository creates a new module repository
func NewModuleRepository(db *gorm.DB) *ModuleRepository {
	return &ModuleRepository{
		Guild:      NewGuildRepository(db),
		User:       NewUserRepository(db),
		Currency:   NewCurrencyRepository(db),
		Warning:    NewWarningRepository(db),
		Poll:       NewPollRepository(db),
		Task:       NewTaskRepository(db),
		Note:       NewNoteRepository(db),
		RoleMenu:   NewRoleMenuRepository(db),
		Activity:   NewActivityRepository(db),
		GameBias:   NewGameBiasRepository(db),
	}
}
