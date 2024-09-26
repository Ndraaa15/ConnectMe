package domain

type Favourite struct {
	UserID   string `gorm:"type:varchar(36);primaryKey"`
	WorkerID string `gorm:"type:varchar(36);primaryKey"`
}
