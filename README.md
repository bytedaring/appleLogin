# 概述
这里记录了，接入AppleID登录，服务端的实现。

针对后端验证苹果提供了两种验证方式：

- 一种是 基于JWT的算法验证
- 一种是 基于授权码的验证

identityToken是一个经过签名的JSON Web Token的方式实现，更多见[Sign in with Apple REST API](https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api)。
另一种方式是校验授权码，其主要分为两步：1. 用户授权后获取code；2. 通过code换取token。

服务端验证identityToken时，App客户端需要将苹果接口返回的identityToken, userID这两个参数传递给服务器，用于验证本次登录的有效性。

服务端，获取到Apple公钥，来验证identityToken的有效性。

公钥获取接口：https://appleid.apple.com/auth/keys ， 文档见[Fetch Apple's Public Key for Verifying Token Signature](https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature)。

# 实现
实现1，主要使用github.com/dgrijalva/jwt-go 和  github.com/lestrrat-go/jwx/jwk

```golang
func VerfiyToken(token string) error {
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
```

实现2 github.com/lestrrat-go/jwx/jwk
```golang
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
```





