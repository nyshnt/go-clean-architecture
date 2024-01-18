package domain

import "context"

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupUsingEmailRequest struct {
	FirstName        string `form:"firstName" binding:"required"`
	LastName         string `form:"lastName" binding:"required"`
	Email            string `form:"email" binding:"required"`
	Country          string `form:"country" binding:"required"`
	DevideId         string `form:"deviceId" binding:"required"`
	DeviceType       string `form:"deviceType" binding:"required"`
	UserReferralCode string `form:"userReferralCode" binding:"required"`
	SignupType       string `form:"signupType"  binding:"required"` // google,facebook,email,apple
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsingEmailResponse struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}

type SignupUsingEmail interface {
	GetUserByEmailId()
}
