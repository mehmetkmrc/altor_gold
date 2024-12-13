package paseto

import (
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/valueobject"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/auth"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/wire"
)

const (
	SymmetricKeySize = 128
)

var (
	_         auth.TokenMaker = (*PasetoToken)(nil)
	PasetoSet                 = wire.NewSet(NewPaseto)
)

type PasetoToken struct {
	tokenTTL   time.Duration
	refreshTTL time.Duration
}

func NewPaseto(cfg *config.Container) (auth.TokenMaker, error) {
	tokenDuration := cfg.Token.TokenTTL
	refreshDuration := cfg.Token.RefreshTTL

	return &PasetoToken{
		tokenTTL:   tokenDuration,
		refreshTTL: refreshDuration,
	}, nil
}

func (pt *PasetoToken) CreateToken(userID string) (string, string, *valueobject.Payload, error) {
	duration := pt.tokenTTL
	payload, err := valueobject.NewPayload(userID, duration)
	if err != nil {
		return "", "", nil, err
	}

	tokenPaseto := paseto.NewToken()
	tokenPaseto.SetExpiration(payload.ExpiredAt)
	tokenPaseto.SetIssuedAt(payload.IssuedAt)
	tokenPaseto.SetString("id", payload.ID)
	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public().ExportHex()
	encrypted := tokenPaseto.V4Sign(secretKey, nil)

	return encrypted, publicKey, payload, nil
}

func (pt *PasetoToken) CreateRefreshToken(payload *valueobject.Payload) (string, string, error) {
	duration := pt.refreshTTL
	tokenPaseto := paseto.NewToken()
	payload.ExpiredAt = payload.ExpiredAt.Add(duration)
	tokenPaseto.SetExpiration(payload.ExpiredAt)
	tokenPaseto.SetIssuedAt(payload.IssuedAt)
	tokenPaseto.SetString("id", payload.ID)
	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public().ExportHex()
	encrypted := tokenPaseto.V4Sign(secretKey, nil)
	return encrypted, publicKey, nil
}

func (pt *PasetoToken) DecodeToken(pasetoToken, publicKeyHex string) (*valueobject.Payload, error) {
	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parsedToken, err := parser.ParseV4Public(publicKey, pasetoToken, nil)
	if err != nil {
		return nil, err
	}

	payload := new(valueobject.Payload)
	expiredAt, err := parsedToken.GetExpiration()
	if err != nil {
		return nil, err
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	issuedAt, err := parsedToken.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	id, err := parsedToken.GetString("id")
	if err != nil {
		return nil, err
	}

	payload = &valueobject.Payload{
		ID:        id,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}

	return payload, nil

}
