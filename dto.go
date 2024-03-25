package main

type SignUpDTO struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type SignInDTO struct {
	SignUpDTO
}

type AuthToken struct {
	Token       string `json:"token"`
	ExpiredTime int    `json:"expired_time"`
}

type CreateSubjectDTO struct {
	Name string `json:"name"`
}

type CreateCommentDTO struct {
	Score int8 `json:"score" validate:"required,min=1,max=10"`
}
