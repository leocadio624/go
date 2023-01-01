package main

import (
	"fmt"

	"example.com/estructuras2/user"
)

func main() {
	martha := user.Usuario{Id: 1, Name: "Martha", Age: 20}
	pedro := user.Usuario{Id: 2, Name: "Pedro", Age: 21}
	jhon := user.Usuario{Id: 3, Name: "Jhon", Age: 20}
	jane := user.Usuario{Id: 4, Name: "Jane", Age: 21}

	//fmt.Printf("%+v\n", martha)
	martha.SayHello()
	martha.AddFriend(pedro)
	martha.AddFriend(jhon)
	martha.AddFriend(jane)

	martha.ListFriends()
	martha.RemoveFriend(jhon.Id)

	fmt.Printf("\n")
	martha.ListFriends()

}
