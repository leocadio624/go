package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status

}
