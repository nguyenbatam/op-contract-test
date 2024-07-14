package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nguyenbatam/op-contract-test/backend/services"
	"time"
)

type apiHandler struct {
	s *services.Service
}

func StartServer(s *services.Service, httpPort string) {
	handler := &apiHandler{
		s: s,
	}
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	restApiNoAuthGroup := router.Group("/api")
	{
		restApiNoAuthGroup.POST("/upload", handler.UserUploadSample)
		restApiNoAuthGroup.POST("/:document_id/info", handler.GetDocumentInfo)
	}
	router.Run(":" + httpPort)
}
