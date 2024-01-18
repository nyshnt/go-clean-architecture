package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/bootstrap"
	"gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/domain"
)

type AuthController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

// func (sc *AuthController) Signup(c *gin.Context) {
// 	var request domain.SignupRequest

// 	err := c.ShouldBind(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
// 	}

// 	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
// 	if err == nil {
// 		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
// 		return
// 	}

// }

func (sc *AuthController) SignupUsingEmail(c *gin.Context) {
	var request domain.SignupRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

}

func (sc *AuthController) SigninUsingEmail(c *gin.Context) {

}

func (sc *AuthController) Verify(c *gin.Context) {

}

func (sc *AuthController) Confirm(c *gin.Context) {

}

func (sc *AuthController) SocialLogin(c *gin.Context) {

}

func (sc *AuthController) AdminSignup(c *gin.Context) {

}

func (sc *AuthController) AdminLogin(c *gin.Context) {

}
