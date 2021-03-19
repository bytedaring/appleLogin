package apple_login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JWK struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"Use"`
}

type JWKSet struct {
	Keys []*JWK `json:"keys"`
}

func getJWKList() (*JWKSet, error) {
	res, err := http.Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var keys JWKSet
	err = json.Unmarshal(body, &keys)
	if err != nil {
		return nil, err
	}
	return &keys, nil
}
