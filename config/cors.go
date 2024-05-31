package config

import "github.com/gin-contrib/cors"

var corsConfig = cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowHeaders:    []string{"Authorization", "Content-Type", "Timezone"},
	AllowMethods:    []string{"PUT", "PATCH", "DELETE", "POST", "OPTIONS", "GET"},
})
