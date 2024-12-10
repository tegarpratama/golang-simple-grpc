package message

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}
