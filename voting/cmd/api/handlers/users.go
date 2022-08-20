package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/galifornia/go-polls-voting/cmd/api/auth"
	"github.com/galifornia/go-polls-voting/models"
)

func LogIn(ctx echo.Context) error {
	usr := new(models.User)
	if err := ctx.Bind(usr); err != nil {
		log.Fatalf("error in bind: %s", err)
		return err
	}

	accessToken, _, err := auth.GenerateAccessToken(usr)
	if err != nil {
		return err
	}

	refreshToken, exp, err := auth.GenerateRefreshToken(usr)
	if err != nil {
		return err
	}
	auth.SetTokenCookie(auth.REFRESH_TOKEN_COOKIE_NAME, refreshToken, exp, ctx)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":      "login successful",
		"access_token": accessToken,
	})
}

func RefreshToken(ctx echo.Context) error {
	cookie, err := ctx.Cookie(auth.REFRESH_TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			return ctx.NoContent(http.StatusNoContent)
		}
		log.Fatal(err)
		return err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &models.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.GetRefreshJWTSecret()), nil
		})
	if err != nil {
		log.Fatal(err)
		return err
	}

	if !token.Valid {
		return errors.New("token is not valid")
	}

	claims := token.Claims.(*models.JwtCustomClaims)
	log.Printf("Claims email is %s", claims.Email)

	usr := &models.User{Name: claims.Email}

	accessToken, _, err := auth.GenerateAccessToken(usr)
	if err != nil {
		return err
	}

	refreshToken, exp, err := auth.GenerateRefreshToken(usr)
	if err != nil {
		return err
	}

	auth.SetTokenCookie(auth.REFRESH_TOKEN_COOKIE_NAME, refreshToken, exp, ctx)

	return ctx.JSON(http.StatusOK, echo.Map{"access_token": accessToken})
}

func GetUser(ctx echo.Context) error {
	accessToken := ctx.Request().Header.Get("Authorization")

	token, err := jwt.ParseWithClaims(accessToken, &models.JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.JWT_SECRET_KEY), nil
		})
	if err != nil {
		log.Println(err)
		return err
	}

	claims := token.Claims.(*models.JwtCustomClaims)
	if err != nil {
		log.Println(err)
		return err
	}

	if !token.Valid {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unathorized",
		})
	}

	// TODO: fetch user info from DB
	return ctx.JSON(http.StatusOK, echo.Map{"email": claims.Email})
}

func UpdateUserInfo(ctx echo.Context) error {
	return nil
}

func RemoveUser(ctx echo.Context) error {
	return nil
}

func LogOut(ctx echo.Context) error {
	auth.SetTokenCookie(auth.REFRESH_TOKEN_COOKIE_NAME, "", time.Now().Add(-time.Hour), ctx)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "User has been logged out",
	})
}

func Register(ctx echo.Context) error {
	usr := new(models.User)
	if err := ctx.Bind(usr); err != nil {
		return err
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 14)
	if err != nil {
		return err
	}

	// !FIXME: throw error if user is already registered
	usr.Password = pwd

	// database.DB.Create(&user)

	return ctx.JSON(http.StatusCreated, usr)
}
