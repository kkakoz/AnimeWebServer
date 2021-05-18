package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type JwtTokenGen struct {
	Issuer     string
	Pem        string
	privateKey *rsa.PrivateKey
	nowFunc    func() time.Time
}

func NewJwtTokenGen(viper *viper.Viper) (*JwtTokenGen, error) {
	j := &JwtTokenGen{}
	err := viper.UnmarshalKey("jwt", j)
	if err != nil {
		return nil, err
	}
	j.nowFunc = time.Now
	j.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(j.Pem))
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (j *JwtTokenGen) GenTokenExpire(id string, expire time.Duration) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    j.Issuer, // 签发人唯一标识
		IssuedAt:  j.nowFunc().Unix(),
		ExpiresAt: j.nowFunc().Add(expire).Unix(),
		Subject:   id,
	})
	signedStr, err := tkn.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}
	return signedStr, nil
}
