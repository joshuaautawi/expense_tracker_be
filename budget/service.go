package budget

type Service interface {
	FindAll() ([]Budget, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Budget, error) {
	return s.repository.FindAll()
}
