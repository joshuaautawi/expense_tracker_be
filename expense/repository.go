package expense

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Expense, error)
    FindById(ID uint64) (Expense, error)
	Create(expense Expense) (Expense, error)
    Update( expense Expense) (Expense, error)
    Delete(expense Expense) (Expense, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Expense, error) {
	var expenses []Expense
	err := r.db.Table("expenses").Select("*").Joins("JOIN categories ON expenses.category_id = categories.id").Find(&expenses).Error
	return expenses, err
}

func (r *repository) FindById(ID uint64)(Expense,error){
    var expense Expense
    err := r.db.Find(&expense,ID).Error
    return expense ,err
}

func (r *repository) Create(expense Expense) (Expense, error) {
	err := r.db.Create(&expense).Error
	return expense, err
}

func (r *repository) Update(expense Expense) (Expense,error){
    err := r.db.Save(&expense).Error
    return expense,err
}

func (r *repository) Delete( expense Expense) (Expense,error){
    err := r.db.Delete(&expense).Error
    return  expense,err 
}
