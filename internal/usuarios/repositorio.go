package usuarios

import (
	"fmt"
)

var users []User
var TokenToPass string = "NormalicemosEcharleMayonesaALosFideosConTuco"

type User struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre" binding:"required"`
	Apellido    string  `json:"apellido" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	Edad        int     `json:"edad" binding:"required"`
	Altura      float64 `json:"altura" binding:"required"`
	Activo      bool    `json:"activo" binding:"required"`
	FechaCreado string  `json:"fechaCreado"`
}

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreado string) (User, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreado string) (User, error)
	UpdateName(id int, nombre string, edad int) (User, error)
	Delete(id int) error
}

type repository struct{}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreado string) (User, error) {

	u := User{id, nombre, apellido, email, edad, altura, activo, fechaCreado}
	users = append(users, u)
	return u, nil

}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreado string) (User, error) {
	u := User{
		Nombre:      nombre,
		Apellido:    apellido,
		Email:       email,
		Edad:        edad,
		Altura:      altura,
		Activo:      activo,
		FechaCreado: fechaCreado,
	}
	updated := false

	for i := range users {
		if users[i].ID == id {
			u.ID = id
			users[i] = u
			updated = true

		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range users {
		if users[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}

func (r *repository) UpdateName(id int, nombre string, edad int) (User, error) {
	var u User
	updated := false
	for i := range users {
		if users[i].ID == id {
			users[i].Nombre = nombre
			users[i].Edad = edad
			updated = true
			u = users[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return u, nil
}

func NewRepository() Repository {
	return &repository{}
}
