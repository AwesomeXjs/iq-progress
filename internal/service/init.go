package service

type IService interface {
}

type Service struct {
}

func New() IService {
	return &Service{}
}
