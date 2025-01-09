package auth

import (
	"os"
	"time"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/golang-jwt/jwt/v5"
)

var itsAsecret = "SECRET_TOKEN"

func GenerateToken(user types.User) string { /// this function generates the token
	key := os.Getenv(itsAsecret)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Name,
		"user_role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := t.SignedString([]byte(key))
	if err != nil {
		return ""
	}
	return token
}

func ParseToken(token string) (*types.TokenContent, error) { /// this function parses the token
	key := os.Getenv(itsAsecret)
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
	content := types.TokenContent{
		ID:   claims["user_id"].(int),
		Name: claims["user_name"].(string),
		Role: claims["user_role"].(string),
	}
	return &content, nil
}

func RoleUser(token string) string { /// this function returns the role of the user
	content, _ := ParseToken(token)
	return content.Role
}

func ExpToken(token string) time.Time { /// this function returns the expiration time of the token
	content, _ := ParseToken(token)
	return content.Exp
}