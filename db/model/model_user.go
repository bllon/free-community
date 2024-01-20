package model

type User struct {
	ID            int64  `gorm:"size:32;primarykey"`
	Name          string `gorm:"size:32;index;"`
	LastLoginDate int64  `gorm:"size:32;not null;default:0;comment:最近登录日期"`
}
