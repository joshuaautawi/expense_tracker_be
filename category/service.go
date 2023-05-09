package category

type Service interface {
	FindAll() ([]Category, error)
	Create(request CategoryRequest) (Category, error)
	Update(ID uint64, request CategoryRequest) (Category, error)
	Delete(ID uint64) (Category, error)
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

func (s *service) Create(request CategoryRequest) (Category, error) {
	category := Category{
		Name: request.Name,
	}
	return s.repository.Create(category)
}

func (s *service) Update(ID uint64, request CategoryRequest) (Category, error) {
	category, _ := s.repository.FindById(ID)
	category.Name = request.Name
	return s.repository.Update(category)
}

func (s *service) Delete(ID uint64) (Category, error) {
	category, _ := s.repository.FindById(ID)
	return s.repository.Delete(category)
}
