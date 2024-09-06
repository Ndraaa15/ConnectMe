package dto

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required,phone"`
}

type VerifyAccountRequest struct {
	ID   string `json:"id" validate:"required"`
	Code string `json:"code" validate:"required,len=4"`
}
