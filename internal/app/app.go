package app

import (
	"log"
	"sharingvision/internal/config"
	"sharingvision/internal/database"
	httpDelivery "sharingvision/internal/delivery/http"
	articleDomain "sharingvision/internal/domain/article"
	"sharingvision/internal/repository"
	"sharingvision/internal/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New(cfg config.Config) (*App, error) {
	db, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&articleDomain.Article{}); err != nil {
		return nil, err
	}

	repo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(repo)
	handler := httpDelivery.NewArticleHandler(articleService)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	httpDelivery.SetupRoutes(router, handler)

	return &App{router: router}, nil
}

func (a *App) Run(addr string) error {
	log.Printf("server is running on %s", addr)
	return a.router.Run(addr)
}
