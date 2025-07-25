/*
 * Copyright © 2025 Berryl Radian Hamesha
 * All rights reserved. Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 * Created by Berryl Radian Hamesha <berrylhamesha@gmail.com> on June 25, 2025
 */

package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/config"
	"github.com/berrylradianh/makerble-golang-coding-assesment/library/helper"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware
var signingKey []byte
var myrole map[string][]string

var (
	blacklist      = make(map[string]time.Time)
	blacklistMutex sync.RWMutex
)

type TokenStructure struct {
	UserID int
	Role   string
	Phone  string
	Email  string
}

type TokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiredIn   float64 `json:"expired_in"`
	ExpiredAt   int64   `json:"expired_at"`
}

type clinicPortalAuth struct {
	signature []byte
}

type ClinicPortalAuth interface {
	GenerateToken(data TokenStructure) (response *TokenResponse, err error)
}

func NewClinicPortalAuth(signature []byte) ClinicPortalAuth {
	return &clinicPortalAuth{signature}
}

const (
	EXPIRED_IN = time.Hour * (24 * 90) // 90 days
)

func GetSigningKey() []byte {
	return signingKey
}

func AddToBlacklist(token string, expiry time.Time) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	blacklist[token] = expiry
}

func IsTokenBlacklisted(token string) bool {
	blacklistMutex.RLock()
	defer blacklistMutex.RUnlock()

	if expiry, exists := blacklist[token]; exists {
		if time.Now().Before(expiry) {
			return true
		}
		blacklistMutex.Lock()
		defer blacklistMutex.Unlock()
		delete(blacklist, token)
	}
	return false
}

func CleanBlacklist() {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	for token, expiry := range blacklist {
		if time.Now().After(expiry) {
			delete(blacklist, token)
		}
	}
}

func NewMiddlewareConfig(conf config.Config) error {

	superadmin := strings.Split(conf.GetString("permission.superadmin"), ",")
	admin := strings.Split(conf.GetString("permission.admin"), ",")
	receptionist := strings.Split(conf.GetString("permission.receptionist"), ",")
	doctor := strings.Split(conf.GetString("permission.doctor"), ",")

	signature := conf.GetString("app.signature")

	role := make(map[string][]string)
	role["superadmin"] = superadmin
	role["admin"] = admin
	role["receptionist"] = receptionist
	role["doctor"] = doctor

	InitRole(role)
	InitJWTMiddlewareCustom([]byte(signature), jwt.SigningMethodHS512)

	go func() {
		for {
			time.Sleep(time.Hour)
			CleanBlacklist()
		}
	}()

	return nil
}

func InitRole(roles map[string][]string) {
	myrole = roles
}

func MyAuth(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := checkJWTToken(ctx.Request); err != nil {
			abortMission(ctx, http.StatusUnauthorized, err)
			return
		}

		for _, role := range roles {
			if err := checkRole(ctx.Request, role); err != nil {
				abortMission(ctx, http.StatusForbidden, err)
				return
			}
		}
	}
}

func abortMission(ctx *gin.Context, statusCode int, err error) {
	response := new(helper.ResponseError)
	response.Status = "failed"
	response.Message = err.Error()

	ctx.AbortWithStatusJSON(statusCode, &response)
}

func checkJWTToken(r *http.Request) error {
	if !jwtMiddleware.Options.EnableAuthOnOptions {
		if r.Method == "OPTIONS" {
			return nil
		}
	}

	token, err := jwtMiddleware.Options.Extractor(r)
	if err != nil {
		eExtractor := errors.New("400")
		return eExtractor
	}

	if token == "" {
		if jwtMiddleware.Options.CredentialsOptional {
			return nil
		}
		eReqiredToken := errors.New("required authorization token not found")
		return eReqiredToken
	}

	if IsTokenBlacklisted(token) {
		return errors.New("token has been expired")
	}

	parsedToken, err := jwt.Parse(token, jwtMiddleware.Options.ValidationKeyGetter)
	if err != nil {
		ePassingToken := errors.New("Error parsing token: " + err.Error())
		return ePassingToken
	}

	if jwtMiddleware.Options.SigningMethod != nil && jwtMiddleware.Options.SigningMethod.Alg() != parsedToken.Header["alg"] {
		errorMsg := fmt.Sprintf("Expected %s signing method but token specified %s",
			jwtMiddleware.Options.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		eTokenSpecified := errors.New(errorMsg)
		return eTokenSpecified
	}

	if !parsedToken.Valid {
		eInvalidToken := errors.New("token invalid")
		return eInvalidToken
	}

	newRequest := r.WithContext(context.WithValue(r.Context(), jwtMiddleware.Options.UserProperty, parsedToken))
	*r = *newRequest
	return nil
}

func InitJWTMiddlewareCustom(secret []byte, signingMethod jwt.SigningMethod) {
	signingKey = secret
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
		SigningMethod: signingMethod,
	})
}

func (cAuth *clinicPortalAuth) GenerateToken(data TokenStructure) (response *TokenResponse, err error) {
	conf := config.NewConfig()

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	expiredIn := EXPIRED_IN
	expiredAt := time.Now().Add(EXPIRED_IN)

	myCrypt, err := bcrypt.GenerateFromPassword([]byte(conf.GetString("app.signature")), 8)
	if err != nil {
		return nil, fmt.Errorf("failed generating password: %v", err)
	}

	claims["user_id"] = data.UserID
	claims["phone"] = data.Phone
	claims["email"] = data.Email
	claims["role"] = data.Role
	claims["hash"] = string(myCrypt)
	claims["exp"] = expiredAt.Unix()

	tokenString, err := token.SignedString(cAuth.signature)
	if err != nil {
		return nil, fmt.Errorf("failed signing string: %v", err)
	}

	response = new(TokenResponse)
	response.AccessToken = tokenString
	response.TokenType = "Bearer"
	response.ExpiredAt = expiredAt.Unix()
	response.ExpiredIn = expiredIn.Seconds()

	return response, nil
}

func checkRole(r *http.Request, roles string) (err error) {
	tokenRole, err := ExtractToken(r, "role")
	if err != nil || tokenRole == nil {
		err = errors.New("you don't have permission to access this route")
		return err
	}

	if roles == "*" {
		return nil
	}

	for k, r := range myrole {
		if k == roles {
			for _, c := range r {
				if c == tokenRole {
					return nil
				}
			}
			break
		}
	}

	err = errors.New("access denied")
	return err
}

func ExtractToken(r *http.Request, key string) (interface{}, error) {
	tokenStr, err := jwtMiddleware.Options.Extractor(r)
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims[key], nil
	} else {
		return "", nil
	}
}

func GetAuthenticatedUser(r *http.Request) (int, error) {
	userID, err := ExtractToken(r, "user_id")
	if err != nil {
		return 0, err
	}
	return int(userID.(float64)), nil
}
