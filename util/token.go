package util

import (
	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"time"
	"fmt"
)

var (
	JWT_KEY               string = beego.AppConfig.String("jwt-key")
	JWT_EXPIRE_SECONDS, _        = beego.AppConfig.Int("jwt-expire-second")
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"json"`
}

type CustomClaims struct {
	User
	jwt.StandardClaims
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", err
	}
	SigningKey := []byte(JWT_KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(JWT_EXPIRE_SECONDS)).Unix()
	newClaims := CustomClaims{
		claims.User,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    claims.User.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString(SigningKey)
	if err != nil {
		log.Error("generate new fresh json web token failed !! error :", err)
		return "", err
	}
	return tokenStr, err
}

func ValidateToken(tokenString string) (u User, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		u = claims.User
	} else {
		log.Info("validate tokenString failed !!!", err)
		return u, err
	}
	return u, nil
}

func GenerateToken(UserId int64, UserName string) (tokenString string) {
	SigningKey := []byte(JWT_KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(JWT_EXPIRE_SECONDS)).Unix()

	user := User{fmt.Sprintf("%d", UserId), UserName}
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(SigningKey)
	if err != nil {
		log.Error("generate json web token failed !! error :", err)
	}
	return tokenStr

}

func getHeaderTokenValue(tokenString string) string {
	return fmt.Sprintf("Bearer %s", tokenString)
}
