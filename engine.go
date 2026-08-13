package server

import "github.com/gin-gonic/gin"

func NewEngine(opts ...EngineOption) (*gin.Engine, error) {
	config := &engineConfig{}

	for _, opt := range opts {
		opt(config)
	}

	engine := gin.New()
	engine.RedirectTrailingSlash = false

	if len(config.trustedProxies) > 0 {
		if err := engine.SetTrustedProxies(config.trustedProxies); err != nil {
			return nil, err
		}
	}

	for _, handler := range config.middleware {
		engine.Use(handler)
	}

	return engine, nil
}
