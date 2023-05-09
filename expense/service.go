package expense

import (
	"time"
)

type Service interface {
	FindAll() ([]Expense, error)
	FindById(ID uint64) (Expense, error)
	Create(request ExpenseRequest) (Expense, error)
	Update(ID uint64, expense ExpenseRequest) (Expense, error)
	Delete(ID uint64) (Expense, error)
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

func (s *service) FindById(ID uint64) (Expense, error) {
	return s.repository.FindById(ID)
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

func (s *service) Update(ID uint64, request ExpenseRequest) (Expense, error) {

	expense, _ := s.repository.FindById(ID)

	expense.Amount = uint64(request.Amount)
	expense.CategoryId = uint64(request.CategoryId)
	expense.Date = uint64(request.Date)
	expense.Description = request.Description

	return s.repository.Update(expense)
}

func (s *service) Delete(ID uint64) (Expense, error) {

	expense, _ := s.repository.FindById(ID)

	return s.repository.Delete(expense)
}
