package message

type messageRepo interface{}

type service struct {
	messageRepo messageRepo
}

func NewService(messageRepo messageRepo) *service {
	return &service{
		messageRepo: messageRepo,
	}
}
