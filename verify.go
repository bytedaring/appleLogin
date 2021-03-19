package apple_login

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	ljwt "github.com/lestrrat-go/jwx/jwt"
)

func Verify(token string, alg string, publicKey string) error {
	parts := strings.Split(token, ".")

	method := jwt.GetSigningMethod(alg)
	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	return err
}

func VerifyToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		set, err := jwk.Fetch(context.Background(), "https://appleid.apple.com/auth/keys")
		if err != nil {
			return nil, err
		}

		kid := t.Header["kid"]
		if key, ok := set.LookupKeyID(fmt.Sprint(kid)); ok {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
			}

			var rawKey interface{}
			if err = key.Raw(&rawKey); err != nil {
				return nil, err
			}

			return rawKey, nil
		}

		return nil, nil
	})

	return err
}

func VerifyToken2(payload string) error {
	set, err := jwk.Fetch(context.Background(), "https://appleid.apple.com/auth/keys")
	if err != nil {
		return err
	}

	token, err := ljwt.Parse([]byte(payload), ljwt.WithKeySet(set), ljwt.UseDefaultKey(true))
	if err != nil {
		return err
	}

	fmt.Println(token)
	return nil
}
