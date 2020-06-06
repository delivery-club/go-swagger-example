package app

type Service struct {
}

// New инициализирует сервис
func New() *Service {
	return &Service{}
}

func (srv *Service) OnShutdown() {
	// do smth on shutdown...
}
