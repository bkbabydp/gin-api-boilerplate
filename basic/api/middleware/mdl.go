package middleware

import (
	"basic/config"
	"github.com/go-eyas/toolkit/log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

// Common 全局通用的中间件
func Common(r *gin.Engine, conf *config.Config) {
	r.Use(Ginzap(log.Logger, time.RFC3339, false))
	r.Use(RecoveryWithZap(log.Logger, false))
	// r.Use(gin.Recovery())
	r.Use(ErrorMiddleware(log.SugaredLogger))

	// Cors
	corsConf := conf.Cors
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     corsConf.Origin,
		AllowMethods:     corsConf.Methods,
		AllowHeaders:     corsConf.Headers,
		AllowCredentials: corsConf.Credentials,
		AllowOriginFunc: func(origin string) bool {
			for _, host := range corsConf.Origin {
				if origin == host || host == "*" {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	}))
}