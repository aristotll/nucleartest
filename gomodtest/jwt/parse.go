package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSampleSecret []byte

func init() {
	log.SetFlags(log.Lshortfile)
}

var (
	Issuer    = "zz"               // Issuer 签发者
	ExpiresAt = time.Hour * 24     // ExpiresAt 24 hours
	Secret    = []byte("SGlsb3g=") // Secret 加密秘钥
)

func CreateJwt(id string) (string, error) {
	// 指定信息
	claims := jwt.StandardClaims{
		Audience:  "",                               // 受众
		ExpiresAt: time.Now().Add(ExpiresAt).Unix(), // 过期时间
		Id:        id,                               // 编号
		IssuedAt:  time.Now().Unix(),                // 签发时间
		Issuer:    Issuer,                           // 签发人
		NotBefore: time.Now().Unix(),                // 生效时间
		Subject:   "login",                          // 主题
	}
	// 创建 token
	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(Secret)
	if err != nil {
		panic(err)
	}
	return token, nil
}

func ParseToken(token string) *jwt.StandardClaims {
	jwtToken, err := jwt.ParseWithClaims(token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
	if err != nil {
		panic(err)
	}

	claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if ok && jwtToken.Valid {
		return claims
	}
	return nil
}

func main() {
	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	tokenString, err := CreateJwt("1")
	if err != nil {
		panic(err)
	}
	tokenString = "Bearer " + tokenString
	log.Println(tokenString)

	sc := ParseToken(tokenString)
	fmt.Printf("sc.Id: %v\n", sc.Id)
}
