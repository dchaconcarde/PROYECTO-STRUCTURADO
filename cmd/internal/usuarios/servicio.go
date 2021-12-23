package usuarios

import "time"

type Service interface {
	GetAll() ([]User, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo bool) (User, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(nombre, apellido, email string, edad int, altura float64, activo bool) (User, error) {
	id := len(users)
	fechaCreado := time.Now().GoString()

	user, err := s.repository.Store(id, nombre, apellido, email, edad, altura, activo, fechaCreado)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
