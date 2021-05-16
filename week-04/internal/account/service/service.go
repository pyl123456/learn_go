package service

type Service struct {
	m MailSender
	u  UserRepository
}

func (s *Service)UserSignUp () {
	s.u.AddUser()
	s.m.Send()
}

func NewService(m MailSender,u  UserRepository) *Service{
	return &Service{m:m, u:u}
}

