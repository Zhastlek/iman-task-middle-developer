package service

type createrPostService struct {
}

func NewCreaterService() ServiceCreater {
	return &createrPostService{}
}
