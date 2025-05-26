package service

type modelService struct{

}

func newModelService() *modelService{
	return &modelService{}
}

func (s *modelService) GetLevel() string{
	return "A1"
}