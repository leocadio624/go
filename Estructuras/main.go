package main

import (
	"fmt"
	"time"

	us "Estructuras/user"
)

type pepe struct {
	us.Usuario
}

func main() {

	u := new(pepe)
	u.AltaUsuario(1, "Victor", time.Now(), true)
	fmt.Println(u.Usuario)

}
