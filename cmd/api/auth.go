package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Auth holds all JWT configuration for the application.
// Create one instance of this at startup and pass it around.
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

// jwtUser holds the minimal user data embedded inside a token.
// We don't put the full user row — just enough to identify them.
type jwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// TokenPairs holds both the access token and the refresh token
// returned to the client on a successful login.
type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Claims embeds jwt.RegisteredClaims so we can add custom fields later
// while still satisfying the jwt.Claims interface.
type Claims struct {
	jwt.RegisteredClaims
}

// GenerateTokenPair creates a signed access token and a signed refresh token
// for the given user, and returns them together as a TokenPairs.
func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {

	// ── ACCESS TOKEN ────────────────────────────────────────────────────────

	// Step 1: create a blank access token object using HS256 signing
	accessToken := jwt.New(jwt.SigningMethodHS256)

	// Step 2: get the claims map and fill in the payload
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	accessTokenClaims["sub"] = fmt.Sprint(user.ID) // subject — who this token is about
	accessTokenClaims["aud"] = j.Audience
	accessTokenClaims["iss"] = j.Issuer
	accessTokenClaims["iat"] = time.Now().UTC().Unix() // issued-at timestamp
	accessTokenClaims["typ"] = "JWT"

	// Step 3: set the expiry (e.g. 15 minutes from now)
	accessTokenClaims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	// Step 4: sign the access token with our secret key
	signedAccessToken, err := accessToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// ── REFRESH TOKEN ───────────────────────────────────────────────────────

	// Step 5: create a blank refresh token object using HS256 signing
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	// Step 6: get the claims map — refresh token only needs sub, iat, exp
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	// Step 7: set a longer expiry for the refresh token (e.g. 7 days)
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	// Step 8: sign the refresh token with the same secret key
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// ── RETURN BOTH ─────────────────────────────────────────────────────────

	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokenPairs, nil
}
