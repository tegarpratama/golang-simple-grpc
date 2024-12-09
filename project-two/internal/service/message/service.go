package message

type messageRepo interface{}

type service struct {
	MessageRepo messageRepo
}

func NewService(messageRepo messageRepo) *service {
	return &service{
		MessageRepo: messageRepo,
	}
}
