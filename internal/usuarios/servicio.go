package usuarios

import "time"

type Service interface {
	GetAll() ([]User, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo bool) (User, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool) (User, error)
	Delete(id int) error
	UpdateName(id int, nombre string, edad int) (User, error)
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
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}
	lastID++
	fechaCreado := time.Now().GoString()

	user, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fechaCreado)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool) (User, error) {

	fechaCreado := time.Now().GoString()
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaCreado)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateName(id int, nombre string, edad int) (User, error) {
	return s.repository.UpdateName(id, nombre, edad)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
