package expense

import (
	"time"
)

type Service interface {
	FindAll() ([]Expense, error)
	Create(request ExpenseRequest) (Expense, error)
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

func (s *service) Create(request ExpenseRequest) (Expense, error) {

	expense := Expense{
		Amount:      request.Amount,
		Description: request.Description,
		CategoryId:  request.CategoryId,
		Date:        uint64(request.Date),
		CreatedAt:   uint64(time.Now().Unix()),
	}
	return s.repository.Create(expense)
}
