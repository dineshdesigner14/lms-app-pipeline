package tokenutil

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/tokendef"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaim struct {
	UserID int
	Email  string
	Role   string
	jwt.StandardClaims
}

func GenToken(UserID int, Email string, Role string, Issuer string, TokenSecret string, TokenExpUnit string, TokenExp int) (int, string) {
	var expiresAt int64
	if strings.EqualFold(TokenExpUnit, tokendef.TOKEN_EXP_SECS) {
		expiresAt = time.Now().Local().Add(time.Second * time.Duration(TokenExp)).Unix()
	} else if strings.EqualFold(TokenExpUnit, tokendef.TOKEN_EXP_MINS) {
		expiresAt = time.Now().Local().Add(time.Minute * time.Duration(TokenExp)).Unix()
	} else if strings.EqualFold(TokenExpUnit, tokendef.TOKEN_EXP_HRS) {
		expiresAt = time.Now().Local().Add(time.Hour * time.Duration(TokenExp)).Unix()
	} else {
		expiresAt = time.Now().Local().Add(time.Minute * time.Duration(TokenExp)).Unix()
	}

	claims := &JwtClaim{
		UserID: UserID,
		Email:  Email,
		Role:   Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		return -1, ""
	}
	return 1, signedToken
}

func VerifyToken(signedToken string, SecretKey string, UserID *int, Email *string, Role *string, Issuer *string, RejectDesc *string) int {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		*RejectDesc = "Token Parse Error"
		return -1
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		*RejectDesc = "Token Parse Error"
		return -1
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		*RejectDesc = "Token Got Expired"
		return -1
	}
	*UserID = claims.UserID
	*Email = claims.Email
	*Role = claims.Role
	*Issuer = claims.Issuer
	return 1
}

func DecodeToken(tokenString string, payload *map[string]interface{}, RejectDesc *string) int {
	parts := strings.Split(tokenString, ".")
	if len(parts) < 2 {
		*RejectDesc = "invalid token format"
		return -1
	}

	payloadSegment := parts[1]

	padding := len(payloadSegment) % 4
	if padding > 0 {
		payloadSegment += strings.Repeat("=", 4-padding)
	}

	payloadBytes, err := base64.URLEncoding.DecodeString(payloadSegment)
	if err != nil {
		*RejectDesc = fmt.Sprintf("error decoding payload: %v", err)
		return -1

	}

	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		*RejectDesc = fmt.Sprintf("error unmarshalling payload: %v", err)
		return -1
	}
	return 1
}
