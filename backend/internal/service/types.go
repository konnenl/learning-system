package service

type WordAnswer struct {
	WordID uint
	Known  bool
}

type PlacementTestInput struct {
	UserID  uint
	Answers []WordAnswer
}
