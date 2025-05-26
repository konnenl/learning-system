package service

type Service struct{
	Model ModelService
	User UserService
	Lesson LessonService
}

func NewService() *Service{
	return &Service{
		Model: newModelService(),
		User: newUserService(),
		Lesson: newLessonService(),
	}
}

type ModelService interface{
	GetLevel() string
}

type UserService interface{
}

type LessonService interface{
}