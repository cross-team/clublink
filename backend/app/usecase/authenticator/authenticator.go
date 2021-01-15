package authenticator

import (
	"errors"
	"time"

	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
)

// Authenticator securely authenticates an user's identity.
type Authenticator struct {
	tokenizer          crypto.Tokenizer
	timer              timer.Timer
	tokenValidDuration time.Duration
}

func (a Authenticator) isTokenValid(payload Payload, validDuring time.Duration) bool {
	now := a.timer.Now()
	if payload.id == "" {
		return false
	}
	tokenExpireAt := payload.issuedAt.Add(validDuring)
	return !tokenExpireAt.Before(now)
}

func (a Authenticator) getPayload(token string) (Payload, error) {
	tokenPayload, err := a.tokenizer.Decode(token)
	if err != nil {
		return Payload{}, err
	}

	payload, err := fromTokenPayload(tokenPayload)
	if err != nil {
		return Payload{}, err
	}
	return payload, nil
}

// IsSignedIn checks whether user successfully signed in
func (a Authenticator) IsSignedIn(token string) bool {
	payload, err := a.getPayload(token)
	if err != nil {
		return false
	}

	return a.isTokenValid(payload, a.tokenValidDuration)
}

// GetUser decodes authentication token to user data
func (a Authenticator) GetUser(token string) (entity.User, error) {
	payload, err := a.getPayload(token)
	if err != nil {
		return entity.User{}, err
	}

	if !a.isTokenValid(payload, a.tokenValidDuration) {
		return entity.User{}, errors.New("token expired")
	}

	if len(payload.id) < 1 {
		return entity.User{}, errors.New("id can't be empty")
	}
	return entity.User{
		ID: payload.id,
	}, nil
}

// GenerateToken encodes part of user data into authentication token
func (a Authenticator) GenerateToken(user entity.User) (string, error) {
	issuedAt := a.timer.Now()
	payload := newPayload(user.ID, issuedAt)
	tokenPayload := payload.TokenPayload()
	return a.tokenizer.Encode(tokenPayload)
}

// NewAuthenticator initializes authenticator with custom token valid duration
func NewAuthenticator(
	tokenizer crypto.Tokenizer,
	timer timer.Timer,
	tokenValidDuration time.Duration,
) Authenticator {
	return Authenticator{
		tokenizer:          tokenizer,
		timer:              timer,
		tokenValidDuration: tokenValidDuration,
	}
}
