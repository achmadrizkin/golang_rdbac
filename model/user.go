package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(100);uniqueIndex" json:"username"` // Set a length for Username
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
}
