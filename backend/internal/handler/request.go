package handler

type userRegisterRequest struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type userLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type placementTestRequest struct {
	Answers []struct {
		ID   uint `json:"id" validate:"required"`
		Know bool `json:"know" validate:"required"`
	} `json:"answers" validate:"required"`
}

type categoryRequest struct {
	Name  string `json:"name" validate:"required"`
	Level string `json:"level" validate:"required,level"`
}
