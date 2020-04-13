package adding

type Repository interface {
	AddBook(*Book) error
}

type Service interface {
	AddBook(*Book) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddBook(b *Book) error {
	return s.r.AddBook(b)
}