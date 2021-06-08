package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
)

type JwtTokenGen struct {
	Issuer     string
	PemPath    string
	privateKey *rsa.PrivateKey
	nowFunc    func() time.Time
}

type UserClaims struct {
	UserId int64    `json:"user_id"`
	Name   string `json:"username"`
	State  int    `json:"state"`
	Auth   int    `json:"auth"`
	jwt.StandardClaims
}

func NewJwtTokenGen(viper *viper.Viper) (*JwtTokenGen, error) {
	j := &JwtTokenGen{}
	err := viper.UnmarshalKey("jwt", j)
	if err != nil {
		return nil, errors.Wrap(err, "viper unmarshal失败")
	}
	j.nowFunc = time.Now
	file, err := os.Open(j.PemPath)
	if err != nil {
		return nil, errors.Wrap(err, "打开rsa privatekey文件失败")
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "读取文件失败")
	}

	j.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		return nil, errors.Wrap(err, "viper unmarshal失败")
	}
	return j, nil
}

func (j *JwtTokenGen) GenTokenExpire(id int64, name string, state, auth int, expire time.Duration) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, UserClaims{
		UserId:         id,
		Name:           name,
		State:          state,
		Auth:           auth,
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.Issuer, // 签发人唯一标识
			IssuedAt:  j.nowFunc().Unix(),
			ExpiresAt: j.nowFunc().Add(expire).Unix(),
		},
	})
	signedStr, err := tkn.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}
	return signedStr, nil
}
