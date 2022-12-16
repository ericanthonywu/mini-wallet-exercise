package Database

import (
	"github.com/google/uuid"
	"time"
)

type Wallet struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	User      User
	UserId    uuid.UUID
	Balance   uint64
	Status    bool
	EnabledAt time.Time
}

func (Wallet) TableName() string {
	return "user_wallet"
}
