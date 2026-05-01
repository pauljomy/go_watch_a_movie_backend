package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	accesTokenClaims := AccessTokenClaims{
		Name: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			Subject:   fmt.Sprintf("%d", user.ID),
			Audience:  jwt.ClaimStrings{j.Audience},
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(j.TokenExpiry)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accesTokenClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	refreshTokenClaims := RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", user.ID),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(j.RefreshExpiry)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	return TokenPairs{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil

}

func (j *Auth) GetRefreshCookie(refershToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Value:    refershToken,
		Path:     j.CookiePath,
		Domain:   j.CookieDomain,
		Expires:  time.Now().UTC().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}

func (j *Auth) GetExpiredRefreshCookie(refershToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Value:    "",
		Path:     j.CookiePath,
		Domain:   j.CookieDomain,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}
