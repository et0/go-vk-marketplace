package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CheckToken(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return next(c)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return next(c)
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return next(c)
			}

			if userID, ok := claims["sub"].(float64); ok {
				c.Set("userID", uint(userID))
				c.Set("isAuthenticated", true)
			}

			return next(c)
		}
	}
}

func IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if isAuth, ok := c.Get("isAuthenticated").(bool); !ok || !isAuth {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authentication required")
		}
		return next(c)
	}
}
