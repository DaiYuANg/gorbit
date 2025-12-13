package jwt

import "crypto/rsa"

type Option func(*jwtOptions)

type jwtOptions struct {
	PrivateKey *rsa.PrivateKey
	SigningAlg string
	PathPrefix string // 你可以只对特定路径启用
}

// 默认配置
func defaultJwtOptions() *jwtOptions {
	return &jwtOptions{
		SigningAlg: "RS256",
		PathPrefix: "/", // 默认全局
	}
}

// Option helpers
func WithPrivateKey(key *rsa.PrivateKey) Option {
	return func(o *jwtOptions) { o.PrivateKey = key }
}

func WithSigningAlg(alg string) Option {
	return func(o *jwtOptions) { o.SigningAlg = alg }
}

func WithPathPrefix(prefix string) Option {
	return func(o *jwtOptions) { o.PathPrefix = prefix }
}
