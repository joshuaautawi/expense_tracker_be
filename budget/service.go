package budget

type Service interface {
	FindAll() ([]Budget, error)
	Create(request BudgetRequest) (Budget, error)
	Update(ID uint64, request BudgetRequest) (Budget, error)
	Delete(ID uint64) (Budget, error)
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

func (s *service) FindById(ID uint64) (Budget, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(request BudgetRequest) (Budget, error) {

	budget := Budget{
		Amount:     uint64(request.Amount),
		Year:       uint(request.Year),
		Month:      uint(request.Year),
		CategoryId: uint64(request.CategoryId),
	}

	return s.repository.Create(budget)

}

func (s *service) Update(ID uint64, request BudgetRequest) (Budget, error) {
	budget, _ := s.repository.FindById(ID)

	budget.CategoryId = request.CategoryId
	budget.Year = request.Year
	budget.Month = request.Month
	budget.Amount = request.Amount

	return s.repository.Update(budget)

}

func (s *service) Delete(ID uint64) (Budget, error) {
	budget, _ := s.repository.FindById(ID)

	return s.repository.Delete(budget)

}
