package middleware

import (
	"net/http"
	"strings"

	"github.com/NatananPh/kiosk-machine-api/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte(config.GetConfig().Auth.Secret)

func RoleBasedMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized: No token provided",
				})
			}
			tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
				}
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized: Invalid token",
				})
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				
				if isAdmin, ok := claims["admin"].(bool); !ok || !isAdmin {
					return c.JSON(http.StatusForbidden, map[string]string{
						"message": "Forbidden: Admin access required",
					})
				}

				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized: Invalid token claims",
			})
		}
	}
}

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized: No token provided",
				})
			}
			tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
				}
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized: Invalid token",
				})
			}

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized: Invalid token claims",
			})
		}
	}
}
