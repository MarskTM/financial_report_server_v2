package model

// Sử dụng dể gửi code trong email
type EmailForgotPassword struct {
	FogortCode string `json:"forgot_code"`
}

