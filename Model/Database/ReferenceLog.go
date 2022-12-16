package Database

type ReferenceLog struct {
	ID          uint   `gorm:"primarykey"`
	Action      string `gorm:"not null"`
	ReferenceId string `gorm:"unique;not null"`
}

func (ReferenceLog) TableName() string {
	return "reference_log"
}
