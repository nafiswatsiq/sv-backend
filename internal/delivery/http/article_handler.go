package http

import (
	"net/http"
	"strconv"
	"strings"

	articleDomain "sharingvision/internal/domain/article"
	"sharingvision/internal/service"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	service service.ArticleService
}

func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var input articleDomain.CreateArticleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	article, err := h.service.Create(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) ListArticles(c *gin.Context) {
	limit := 10
	offset := 0

	if limitParam := c.Param("limit"); limitParam != "" {
		if parsed, err := strconv.Atoi(limitParam); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if offsetParam := c.Param("offset"); offsetParam != "" {
		if parsed, err := strconv.Atoi(offsetParam); err == nil && parsed >= 0 {
			offset = parsed
		}
	}
	if qLimit := strings.TrimSpace(c.Query("limit")); qLimit != "" {
		if parsed, err := strconv.Atoi(qLimit); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if qOffset := strings.TrimSpace(c.Query("offset")); qOffset != "" {
		if parsed, err := strconv.Atoi(qOffset); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	articles, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid article id"})
		return
	}

	article, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid article id"})
		return
	}

	var input articleDomain.UpdateArticleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	article, err := h.service.Update(c.Request.Context(), uint(id), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid article id"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article deleted successfully"})
}

func (h *ArticleHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Sharing Vision backend is running"})
}

func SetupRoutes(router *gin.Engine, handler *ArticleHandler) {
	router.GET("/health", handler.HealthCheck)
	router.POST("/article", handler.CreateArticle)
	router.GET("/articles", handler.ListArticles)
	router.GET("/article/:id", handler.GetArticleByID)
	router.PUT("/article/:id", handler.UpdateArticle)
	router.PATCH("/article/:id", handler.UpdateArticle)
	router.DELETE("/article/:id", handler.DeleteArticle)

	router.GET("/article", func(c *gin.Context) {
		limitStr := strings.TrimSpace(c.Query("limit"))
		offsetStr := strings.TrimSpace(c.Query("offset"))
		limit := 10
		offset := 0
		if limitStr != "" {
			if parsed, err := strconv.Atoi(limitStr); err == nil && parsed > 0 {
				limit = parsed
			}
		}
		if offsetStr != "" {
			if parsed, err := strconv.Atoi(offsetStr); err == nil && parsed >= 0 {
				offset = parsed
			}
		}
		articles, err := handler.service.List(c.Request.Context(), limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, articles)
	})
}
