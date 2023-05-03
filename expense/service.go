package expense

type Service interface {
	FindAll() ([]Expense, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Expense, error) {
	return s.repository.FindAll()
}
