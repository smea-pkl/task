package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Config menyimpan konfigurasi pembuatan dan validasi token JWT.
type Config struct {
	Secret string        // kunci HMAC untuk menandatangani token
	TTL    time.Duration // masa berlaku token
	Issuer string        // penerbit token
}

// Claims adalah data yang disimpan di dalam token JWT.
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// NewConfig membangun konfigurasi dari environment variable:
//   - JWT_SECRET    (default: "***")
//   - JWT_TTL_HOURS (default: 24)
//   - JWT_ISSUER    (default: "pkl")
func NewConfig() Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "***"
	}

	ttl := 24 * time.Hour
	if v := os.Getenv("JWT_TTL_HOURS"); v != "" {
		if h, err := strconv.Atoi(v); err == nil && h > 0 {
			ttl = time.Duration(h) * time.Hour
		}
	}

	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "pkl"
	}

	return Config{Secret: secret, TTL: ttl, Issuer: issuer}
}

// Generate membuat token JWT (HS256) untuk user tertentu.
func (c Config) Generate(userID uint, username string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    c.Issuer,
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(c.TTL)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(c.Secret))
}

// Parse memvalidasi token JWT dan mengembalikan claims-nya.
func (c Config) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tak terduga: %v", t.Header["alg"])
		}
		return []byte(c.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}
	return claims, nil
}

// ExtractToken mengambil token dari header "Authorization: Bearer <token>".
func ExtractToken(authHeader string) string {
	const prefix = "bearer "
	if len(authHeader) >= len(prefix) && strings.EqualFold(authHeader[:len(prefix)], prefix) {
		return strings.TrimSpace(authHeader[len(prefix):])
	}
	return ""
}
