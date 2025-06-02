package service

type modelService struct {
}

func newModelService() *modelService {
	return &modelService{}
}

func (s *modelService) GetLevel(input PlacementTestInput) (string, error) {
	return "B1", nil
}
