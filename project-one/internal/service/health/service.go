package health

type healthRepo interface{}

type service struct {
	healthRepo healthRepo
}

func NewService(healthRepo healthRepo) *service {
	return &service{
		healthRepo: healthRepo,
	}
}
