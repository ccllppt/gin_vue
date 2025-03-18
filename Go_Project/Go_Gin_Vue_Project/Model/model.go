package Model

import "gorm.io/gorm"

// User 是用户模型，对应数据库中的用户表
type User struct {
	gorm.Model        // 内嵌 GORM 的默认模型（包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段）
	Name       string `gorm:"type:varchar(20);not null"`         // 用户名，不能为空
	Telephone  string `gorm:"type:varchar(100);not null;unique"` // 手机号，唯一且不能为空
	Password   string `gorm:"type:varchar(255);not null"`        // 密码，不能为空
}
