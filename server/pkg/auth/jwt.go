package auth

import (
	"os"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user types.User) string {
	key := os.Getenv("SECRET_TOKEN")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Name,
		"user_role": user.Role,
	})
	token, err := t.SignedString(key)
	if err != nil {
		return ""
	}
	return token
}

func ParseToken(token string) (*types.User, error) {
	key := os.Getenv("SECRET_TOKEN")
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	user := types.User{
		ID:   claims["user_id"].(int),
		Name: claims["user_name"].(string),
		Role: claims["user_role"].(string),
	}
	return &user, nil
}

func RoleUser(token string) string {
	user, _ := ParseToken(token)
	return user.Role
}