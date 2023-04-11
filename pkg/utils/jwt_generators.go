package utils

import (
	"NewProUser/configs"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var config = configs.Load()

type Token struct {
	Access  string
	Refresh string
}

func GenerateNewTokenForUser(id string, credentials map[string]string) (*Token, error) {
	accessToken, err := generateAccessTokenForUser(id, credentials)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateNewRefreshToken(id, credentials)
	if err != nil {
		return nil, err
	}
	return &Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateAccessTokenForUser(id string, credentials map[string]string) (string, error) {
	//create a new claims
	claims := jwt.MapClaims{}
	// set id
	claims["id"] = id
	// set role
	claims["role"] = credentials["role"]

	if config.Environment == "development" {
		//in local server  access token 10 days
		claims["expires"] = time.Now().Add(time.Minute * time.Duration(10*config.JWTSecretKeyExpireMinutes)).Unix()
	} else {
		//in local server acces token 1day
		claims["expires"] = time.Now().Add(time.Minute * time.Duration(config.JWTSecretKeyExpireMinutes)).Unix()
	}
	// create new jwt token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Generate token

	t, err := token.SignedString([]byte(config.JWTSecretKey))
	if err != nil {
		return "", nil
	}
	return t, nil
}

func generateNewRefreshToken()(string, error){
	sha256:=sha256.New()

	refresh:=config.JWTSecretKey+time.Now().String()

	_, err:=sha256.Write([]byte(refresh))
if err!=nil {
	return "",nil
}
expireTime:=fmt.Sprint(time.Now().Add(time.Hour*time.Duration(config.JWTSecretKeyExpireMinutes)).Unix())

//create a new refresh token
t:=hex.EncodeToString(sha256.Sum(nil))+"."+expireTime

return t, nil
}

func ParseRefreshToken(refreshToekn string)(int64, error){
	return strconv.ParseInt(strings.Split(refreshToekn, ".")[1],0,64)
}
