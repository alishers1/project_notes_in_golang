package handlers

import (
	"net/http"
	"notes/configs"
	"notes/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CustomHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Custom header", "Alisher")
		c.Next()
	}
}

func InitRoutes() error {
	r := gin.Default()

	utils.LoadNotesFromFile()

	v1 := r.Use(CustomHeaderMiddleware())

	v1 = r.Group("/notes")
	{
		v1.POST("/", CreateNote)
		v1.GET("/", GetNotes)
		v1.GET("/:id", GetNote)
		v1.PUT("/:id", UpdateNote)
		v1.DELETE("/:id", DeleteNote)
	}

	config, err := configs.InitConfigs()
	if err != nil {
		return err
	}

	address := config.Host + config.Port

	srv := http.Server{
		Addr:    address,
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
