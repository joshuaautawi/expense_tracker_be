package category

type Service interface {
	FindAll() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Category, error) {
	return s.repository.FindAll()
}
