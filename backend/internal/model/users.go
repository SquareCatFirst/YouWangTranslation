package model

import "time"

type User struct {
	ID        int64   `gorm:"column:id;primaryKey;autoIncrement"`
	AvatarURL *string `gorm:"column:avatar_url;type:varchar(255)"`
	Nickname  string  `gorm:"column:nickname;type:varchar(64);not null"`

	Role     int     `gorm:"column:role;not null;default:1"`
	QQOpenID *string `gorm:"column:qq_openid;type:varchar(64);unique"`
	Status   int     `gorm:"column:status;not null;default:1"`

	Username     *string `gorm:"column:username;type:varchar(64);unique"`
	PasswordHash *string `gorm:"column:password_hash;type:text"`

	RegisteredAt time.Time  `gorm:"column:registered_at;autoCreateTime"`
	LastLoginAt  *time.Time `gorm:"column:last_login_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
