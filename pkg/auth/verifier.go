package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type JwtTokenVerifier struct {
	PublicKey string
	publicKey *rsa.PublicKey
}

func NewJwtTokenVerifier(viper *viper.Viper) (*JwtTokenVerifier, error) {
	verifier := &JwtTokenVerifier{}
	err := viper.UnmarshalKey("jwt", verifier)
	if err != nil {
		return nil, err
	}
	pem, err := jwt.ParseRSAPublicKeyFromPEM([]byte(verifier.PublicKey))
	if err != nil {
		return nil, err
	}
	return &JwtTokenVerifier{publicKey: pem}, nil
}

func (v *JwtTokenVerifier) Verifier(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return v.publicKey, nil
		})
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errors.New("token not valid")
	}
	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("token not standardclaims")
	}

	if err = clm.Valid(); err != nil {
		return "", errors.New("claim not valid:" + err.Error())
	}
	return clm.Subject, nil
}
