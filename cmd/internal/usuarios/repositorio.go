package usuarios

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

func NewRepository() Repository {
	return &repository{}
}
