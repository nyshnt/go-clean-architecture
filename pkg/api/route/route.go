package route

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/api/middleware"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/bootstrap"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/mongo"
)

var started time.Time

// handler function of the route
func hello(c *gin.Context) {
	started = time.Now()
	// c.JSON(http.StatusOK, gin.H{"Message": "Hello World"})
	uptime := time.Since(started).Round(time.Second).Seconds()
	c.JSON(http.StatusOK, gin.H{
		"started": fmt.Sprintf("Started at: %s", started),
		"uptime":  uptime,
	})
}

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRoute := gin.Group("")
	// All Public APIs
	publicRoute.GET("/", hello)

	// NewSignupRouter(env, timeout, db, publicRoute)
	AuthRoutes(env, timeout, db, publicRoute)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	UserProfileRoutes(env, timeout, db, protectedRouter)
	// NewProfileRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)

}
