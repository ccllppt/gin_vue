package dto

import "Go_Gin_Vue_Project/Model"

type UserDTO struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDTOs(user Model.User) UserDTO {
	return UserDTO{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
