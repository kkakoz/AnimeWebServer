package auth

import (
	"github.com/spf13/viper"
	"strconv"
	"testing"
	"time"
)

func TestJwtTokenGen(t *testing.T) {
	//pem, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	//if err != nil {
	//	t.Fatal("parse private err:", err)
	//}
	var id int64 = 123456
	viper.SetDefault("jwt.pemPath", "configs/privatekey")
	viper.SetDefault("jwt.publicKeyPath", "configs/publickey")
	//jwtTokenGen.nowFunc = func() time.Time {
	//	return time.Unix(1516239022, 0)
	//}
	tokenGen, err := NewJwtTokenGen(viper.GetViper())
	if err != nil {
		t.Fatal("new jwt token gen err:", err)
	}
	token, err := tokenGen.GenTokenExpire(strconv.FormatInt(id, 10), time.Second * 10)
	if err != nil {
		t.Error("gen token err:", err)
	}
	verifier, err := NewJwtTokenVerifier(viper.GetViper())
	if err != nil {
		t.Fatal("new jwt token verifier err:", err)
	}
	userId, err := verifier.Verifier(token)
	if err != nil {
		t.Fatal("verifier err:", err)
	}
	if userId != strconv.FormatInt(id, 10) {
		t.Fatal("token valid mistake")
	}
}

const privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDJJwr4RwYtD6tg
oCTC9yr5bZAsJY+ntnw9yz2BedTtzujx/R4dMzUB4Kaix5MYIDSjmk+xCNd3pTu7
g2zGS3pykgGIQ1ycx1Q2Bo6s2iLWMU4x/2FTnPp2xRssZ1eoBvW5tYr456IGL0HH
talOkiPn++IgPX6hsOjkiqSE9l7dvcsc6eg8KzRsVPEngKGK3QfZZfi+uoIoLD0J
xc6R1oWtraeZRrnnxUPqgpVTqJ9GhnUl7FMT5b7MfYdzkP3z7GCq0GlETWIbANOt
Y/QuTwNlndakfdQDDdEisASEVXqkX76oaZVDza0Ydav1LQDkJE/L8S492VsyXiad
U4S8G9utAgMBAAECggEBAMarYRJGU7s9tq4AfU5ygGdJ5xMzJecnPR5rFZxDkCIw
MbKPpKaCZOAt2Wb8ZjN124eaaQsZCHaLI6vX6h8PfSibPBgxL89Ir8uMPm5KJA4Z
NHn9GUtTx6x2kJgSmWjDNv7sZw3e+Q/SrM1qhoWroPsGtBfTpLZd3CedJ1CLZLbr
c8LQgUS2aPsvjgFMRStB/ZINTSqL1IKZSWNoc47i6dY8RXtBlv0Nc1eMpoE4XMZ2
6YJMsXUGsuGntA7PMGj60SvKRvY18yy/xCxeUQMIM2AmpQqenGzIm5B0NNv2zxLT
vueXf8Hk+4uwjNrQ5yiuftH3l0Jv8Jat/qvfzbws/qECgYEA87YkfgCzRWS1lvUe
Ewp3oHuQa/Dh0jpTYver+oMN+vDdOwird0BfYsFs6DHky14SI5lCkePMIvJ76dTt
mPTlp3Bparv1voru09Y9eqJLEeTqYOr4vNV4ucxj5jr+VAjSfLGlfLv83i591qdK
C2b2jP03U9BwBKawEgBsOHb3ARcCgYEA00uLNnLO6+QewyTFT7q+Z3IKdJGBPNw7
dU9hZv849NJ+vw1uKhMtnzPn1mYkx+vLK7P/tLcDRXDrWaGJS2YQtpuGFeml2QYb
nUPjL8nd8y2jRW/Y6K9x8Akw+1LRZWggH8eInXHvAMasVno5I4TgZyTzl2/yI0tM
Sp16o7Ujm9sCgYEA38aPG2NpOH6QflvzkWg7D5Blu7ciovYLOwRPVWagn5oqiNod
FxJ0gyk35rxpaJKn9Sf0iCCygCRGwx5QS/ISLPx6zxZnPt4zDS/ao5ABfhzDWNpo
KnuYroGN+QiSHnc7TmOPoEi8lwX5Ze+VfYK9QBgBhWQOdzbW1LCureoOQ2ECgYEA
vZPtLwgpceq2Ux59zkBeL9BZYydeDm4HBwUW/mOGBduLDv4M1sFoUIwwuePhomKE
YwzYI5uE2twqvbu6xKSp4D2AO87sF+FsC1lq0GjNtC9Ba76jnnozv0tv4D75U4Pu
NrU/dQxRhZ+75sc41w6UjNbRnBE77sDTjarn61RNw7sCgYBN5S46n5kJQDb3+LP+
rp0D24nYJz4psRlGND30NSrXiO9heQZAGm9eHGggSlFcWcA3wZQS8tW+HFFyPAlz
E3kuKiamOiAkC/v2GCEXTw6IX4ffmdgCj1TAWOuO3i03a5QwYHCABv9B5nVx1M7o
WtR8Mbhxmh/OcibHqjujI35qmA==
-----END PRIVATE KEY-----
`

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyScK+EcGLQ+rYKAkwvcq
+W2QLCWPp7Z8Pcs9gXnU7c7o8f0eHTM1AeCmoseTGCA0o5pPsQjXd6U7u4Nsxkt6
cpIBiENcnMdUNgaOrNoi1jFOMf9hU5z6dsUbLGdXqAb1ubWK+OeiBi9Bx7WpTpIj
5/viID1+obDo5IqkhPZe3b3LHOnoPCs0bFTxJ4Chit0H2WX4vrqCKCw9CcXOkdaF
ra2nmUa558VD6oKVU6ifRoZ1JexTE+W+zH2Hc5D98+xgqtBpRE1iGwDTrWP0Lk8D
ZZ3WpH3UAw3RIrAEhFV6pF++qGmVQ82tGHWr9S0A5CRPy/EuPdlbMl4mnVOEvBvb
rQIDAQAB
-----END PUBLIC KEY-----
`