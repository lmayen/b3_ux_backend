package auth

import (
	"b3_ux_backend/internal/entities"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olekukonko/errors"
)

const (
	UserAuthContextKey  = "userAuthData"
	SessionCookieName   = "b3_session"
	SecretLength        = 32
	sessionCookieMaxAge = 30 * 24 * 60 * 60
)

var errInvalidCredentials = errors.New("invalid credentials")
var errInvalidSession = errors.New("invalid session")

type UserAuthData struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IsAdmin  bool      `json:"isAdmin"`
}

func (a *UserAuthData) AddToContext(ctx *gin.Context) {
	ctx.Set(UserAuthContextKey, a)
}

func NewUserAuthDataFromDB(artist *entities.User) *UserAuthData {
	return &UserAuthData{
		ID:       artist.ID,
		Username: artist.Username,
		IsAdmin:  artist.IsAdmin,
	}
}

// GenerateSessionToken creates a new credential suitable for a session cookie.
func GenerateSessionToken() (string, error) {
	randomBytes := make([]byte, SecretLength)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("generate secure secret: %w", err)
	}

	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)
	return encoded, nil
}

// HashSecret creates the value stored in the database. The original token or API key cannot be recovered from this hash.
func HashSecret(secret string) string {
	hash := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(hash[:])
}
