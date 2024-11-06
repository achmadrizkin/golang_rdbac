package model

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);uniqueIndex" json:"name"`
}
