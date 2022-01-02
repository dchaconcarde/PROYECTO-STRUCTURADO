package usuarios

import (
	"fmt"

	"github.com/dchaconcarde/proyecto-structurado/pkg/store"
)

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
	LastID() (int, error)
}

type repository struct {
	db store.Store
}

func (r *repository) LastID() (int, error) {
	var users []User
	if err := r.db.Read(&users); err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, nil
	}

	return users[len(users)-1].ID, nil

}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	r.db.Read(&users)
	return users, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreado string) (User, error) {
	var users []User
	r.db.Read(&users)
	u := User{id, nombre, apellido, email, edad, altura, activo, fechaCreado}
	users = append(users, u)
	if err := r.db.Write(users); err != nil {
		return User{}, err
	}
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
	var users []User
	r.db.Read(&users)

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
	if err := r.db.Write(users); err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var users []User

	r.db.Read(&users)

	for i := range users {
		if users[i].ID == id {
			deleted = true
			users = append(users[:i], users[i+1:]...)
		}
	}
	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", id)
	}
	if err := r.db.Write(&users); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateName(id int, apellido string, edad int) (User, error) {
	var u User
	updated := false
	var users []User
	r.db.Read(&users)
	for i := range users {
		if users[i].ID == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			updated = true
			u = users[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	if err := r.db.Write(users); err != nil {
		return User{}, err
	}
	return u, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
