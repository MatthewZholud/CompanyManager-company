package company

type Reader interface {
	Get(id ID) (*Company, error)
	Search(query string) ([]*Company, error)
	List() ([]*Company, error)
}

//Writer book writer
type Writer interface {
	Create(e *Company) (ID, error)
	Update(e *Company) error
	Delete(id ID) error
}

//repository interface
type repository interface {
	Reader
	Writer
}

type companyRepository struct {
	repo repository
}

func NewCompanyRepository() *companyRepository {
	return &companyRepository{}
}

