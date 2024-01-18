package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/bootstrap"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/mongo"
)

func UserProfileRoutes(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(db, domain.CollectionUser)
	// ac := controller.AuthController{
	// 	SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	// 	Env:           env,
	// }
	// group.POST("/signup", sc.Signup)

}
