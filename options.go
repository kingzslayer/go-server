package server

import "github.com/gin-gonic/gin"

type EngineOption func(*engineConfig)

type engineConfig struct {
	middleware     []gin.HandlerFunc
	trustedProxies []string
}

func WithMiddleware(handlers ...gin.HandlerFunc) EngineOption {
	return func(ec *engineConfig) {
		ec.middleware = append(ec.middleware, handlers...)
	}
}

func WithTrustedProxies(proxies []string) EngineOption {
	return func(ec *engineConfig) {
		ec.trustedProxies = proxies
	}
}
