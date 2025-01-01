package main

import (
	"net/http"

	"github.com/MortalSC/GoForIM/internal/auth"
	"github.com/MortalSC/GoForIM/internal/user"
	"github.com/MortalSC/GoForIM/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db.InitDB()
	defer db.DB.Close()

	r := gin.Default()

	r.POST("/register", user.Register)
	r.POST("/login", user.Login)

	protected := r.Group("/api")
	protected.Use(auth.JWTMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
	})

	r.Run(":8080")
	// log.Println("Server starting...")
}
