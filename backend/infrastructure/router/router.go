package router

import (
	"grease_trace/interfaces/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ginのハンドラ関数のシグネチャに合わせる
func greeting_hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func NewRouter(sessionCtrl *controller.SessionController) *gin.Engine {
	r := gin.Default()

	r.Use(CORS())

	api := r.Group("/api")

	// ルート定義をスコープ外に出す
	test := api.Group("/test")
	session := api.Group("/session")

	// APIグループの定義
	{
		session.POST("/create", sessionCtrl.Create)
		session.GET("/validate/:sessionId", sessionCtrl.Validate)
	}

	// TESTグループの定義
	{
		test.GET("/greeting", greeting_hello)
	}

	return r
}
