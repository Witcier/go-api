package utils

import (
	"errors"
	"time"
	"witcier/go-api/global"
	"witcier/go-api/model/request"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	Signingkey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		// 缓冲时间1天 缓冲时间内会获得新的token刷新令牌
		BufferTime: global.Config.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			// 签名生效时间
			NotBefore: time.Now().Unix() - 1000,
			// 过期时间 7天
			ExpiresAt: time.Now().Unix() + global.Config.JWT.ExpiresTime,
			// 签名的发行者
			Issuer: global.Config.JWT.Issuer,
		},
	}

	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.Signingkey)
}

func (j *JWT) RefreshToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})

	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.Signingkey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}

		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
