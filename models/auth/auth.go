package auth

type RegisterRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	DateOfBirthday string `json:"date_of_birthday" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Gender         string `json:"gender"`
	DateOfBirthday string `json:"date_of_birthday"`
}

type ResponseUser struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	DateOfBirthday string `json:"date_of_birthday"`
}

type ResponseLogin struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	DateOfBirthday string `json:"date_of_birthday"`
	Token          string `json:"token"`
}
