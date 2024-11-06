package model

type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"type:varchar(100);uniqueIndex" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}
