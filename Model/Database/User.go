package Database

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primarykey"`
}

func (User) TableName() string {
	return "user"
}
