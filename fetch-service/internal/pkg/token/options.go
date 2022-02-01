package token

type Options func(*Token)

func WithSecretKey(secretKey string) Options {
	return func(token *Token) {
		token.hmac = []byte(secretKey)
	}
}

func WithAccessTokenTTL(ttl uint64) Options {
	return func(token *Token) {
		token.ttl.accessToken = ttl
	}
}

func WithRefreshTokenTTL(ttl uint64) Options {
	return func(token *Token) {
		token.ttl.refreshToken = ttl
	}
}

func WithIssuer(issue string) Options {
	return func(token *Token) {
		token.issuer = issue
	}
}
