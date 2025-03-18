package common

import (
	"Go_Gin_Vue_Project/Model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwtKey 是用于签名和验证 JWT 的密钥
var jwtKey = []byte("a_secret_creat")

// Claims 是 JWT 的声明结构体，包含用户 ID 和标准声明
type Claims struct {
	UserId             uint // 用户 ID
	jwt.StandardClaims      // JWT 标准声明（如过期时间、签发者等）
}

// RelaeseToken 生成并返回一个 JWT Token
func RelaeseToken(user Model.User) (string, error) {
	// 设置 Token 的过期时间为 7 天后
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// 创建 JWT 的声明
	claims := &Claims{
		UserId: user.ID, // 用户 ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 签发时间
			Issuer:    "oceanlearn.tech",     // 签发者
			Subject:   "user token",          // 主题
		},
	}

	// 使用 HS256 算法生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对 Token 进行签名
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err // 如果签名失败，返回错误
	}

	// 返回生成的 Token 字符串
	return tokenString, nil
}

// ParseToken 解析并验证 JWT Token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	// 创建一个空的 Claims 结构体
	claims := &Claims{}

	// 解析 Token，并将声明填充到 Claims 结构体中
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil // 返回密钥用于验证
	})

	// 返回解析后的 Token、Claims 和可能的错误
	return token, claims, err
}
