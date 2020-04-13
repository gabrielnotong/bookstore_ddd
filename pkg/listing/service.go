package listing

type Repository interface {
	AllBooks() ([]*Book, error)
	OneBook(string) (*Book, error)
}

type Service interface {
	AllBooks() ([]*Book, error)
	OneBook(string) (*Book, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AllBooks() ([]*Book, error) {
	return s.r.AllBooks()
}

func (s *service) OneBook(isbn string) (*Book, error) {
	return s.r.OneBook(isbn)
}