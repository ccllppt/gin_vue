package dto

import "Go_Gin_Vue_Project/Model"

// UserDTO 是用户数据传输对象，用于返回用户信息时隐藏敏感字段
type UserDTO struct {
	Name      string `json:"name"`      // 用户名
	Telephone string `json:"telephone"` // 手机号
}

// ToUserDTOs 将 User 模型转换为 UserDTO
func ToUserDTOs(user Model.User) UserDTO {
	return UserDTO{
		Name:      user.Name,      // 设置用户名
		Telephone: user.Telephone, // 设置手机号
	}
}
