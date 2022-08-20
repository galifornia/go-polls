package auth

import (
	"net/http"
	"time"

	"github.com/galifornia/go-polls-voting/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	ACCESS_TOKEN_COOKIE_NAME  = "access_token"
	REFRESH_TOKEN_COOKIE_NAME = "refresh_token"
	JWT_SECRET_KEY            = "some-secret-key"
	JWT_REFRESH_SECRET_KEY    = "some-refresh-secret-key"
)

func GetRefreshJWTSecret() string {
	return JWT_REFRESH_SECRET_KEY
}

func GenerateRefreshToken(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetRefreshJWTSecret()))
}

func GenerateAccessToken(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(JWT_SECRET_KEY))
}

func generateToken(user *models.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &models.JwtCustomClaims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func SetTokenCookie(name, token string, expiration time.Time, ctx echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.HttpOnly = true

	ctx.SetCookie(cookie)
}
