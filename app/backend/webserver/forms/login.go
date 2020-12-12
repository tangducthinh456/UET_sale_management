package forms

type Token struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}

type LoginForm struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}
