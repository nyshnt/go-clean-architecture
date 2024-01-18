package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/api/controller"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/bootstrap"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/domain"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/mongo"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/repository"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/usecase"
)

func AuthRoutes(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	ac := controller.AuthController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	// group.POST("/signup", sc.Signup)
	group.POST("/signup-using-email", ac.SignupUsingEmail)
	group.POST("/signin-using-email", ac.SigninUsingEmail)
	group.GET("/verify", ac.Verify)
	group.GET("/confirm", ac.Confirm)
	group.POST("/social-login", ac.SocialLogin)
	group.POST("/admin-signup", ac.AdminSignup)
	group.POST("/admin-login", ac.AdminLogin)

}
