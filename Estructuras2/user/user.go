package user

import "fmt"

type Usuario struct {
	Id      int
	Name    string
	Age     int
	friends []Usuario
}

func (u Usuario) SayHello() {
	fmt.Println("Hola me llamo", u.Name)
}

func (u *Usuario) AddFriend(friend Usuario) {
	u.friends = append(u.friends, friend)
}

func (u *Usuario) RemoveFriend(Id int) {
	index := u.findFriend(Id)
	if index < 0 {
		return
	}
	u.friends = append(u.friends[:index], u.friends[index+1:]...)
}

func (u Usuario) findFriend(Id int) int {
	for i, f := range u.friends {
		if f.Id == Id {
			return i
		}
	}
	return -1
}

func (u Usuario) ListFriends() {
	for i, f := range u.friends {
		fmt.Printf("%d. %s\n", i+1, f.Name)
	}
}
