package config

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	// Set secret key from .env file.
	secret = Config("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ = strconv.Atoi(Config("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
}

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(id uint) (string, error) {
	// Create a new claims.
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func CheckToken(c *fiber.Ctx) (uint, error) {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := extractClaims(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return 0, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "msg": err.Error()})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.ExpiresAt

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return 0, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	issuer, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return 0, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "msg": err.Error()})
	}
	return uint(issuer), nil
}

// ExtractClaims func to extract claims from JWT.
func extractClaims(c *fiber.Ctx) (*jwt.StandardClaims, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if ok && token.Valid {
		// Expires time.
		return claims, nil
	}

	return nil, err
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(Config("JWT_SECRET_KEY")), nil
}

// GenerateNewCookie func for generate a new fiber cookie.
func GenerateNewCookie(token string, login bool) fiber.Cookie {
	var t time.Duration
	if login {
		t = time.Minute * time.Duration(minutesCount)
	} else {
		t = -time.Hour
	}

	return fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(t),
		HTTPOnly: true,
	}
}
