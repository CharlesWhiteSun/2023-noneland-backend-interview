package pkg

import (
	"net/http"
	"nonelandBackendInterview/internal/api"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// InitHttpHandler 為了分成測試用與正式用，所以把 gin 的初始化抽出來
func InitHttpHandler() (h http.Handler) {
	return h2c.NewHandler(setupGin(), &http2.Server{})
}

func setupGin() http.Handler {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	apiGroup := r.Group("/api")

	// TODO: api router
	apiGroup.GET("hello", api.HelloHandler)

	return r
}
