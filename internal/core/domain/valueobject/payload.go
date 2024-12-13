package valueobject

import (
	"errors"
	"time"
)

const (
	AccessToken  = "access"
	RefreshToken = "refresh"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

type(
	Payload struct{
		ID string `json:"id"`
		IssuedAt time.Time `json:"issued_at"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		ID:        userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if !time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

