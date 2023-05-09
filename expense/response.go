package expense


type Category struct {
	Name string
}

type ExpenseResponse struct {
	ID          uint64    
	CategoryId  uint64   
	Amount      uint64
	Description string
	Date        uint64
	CreatedAt   uint64 
	UpdatedAt   uint64 
    Category Category
}
