package health

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}
