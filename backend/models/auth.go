package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	SessionToken string `json:"session_token"`
	CRSFToken    string `json:"crsf_token"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
}
