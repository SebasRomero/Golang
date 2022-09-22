package users

import "fmt"

type User struct {
	ID      int
	Name    string
	Age     int
	friends []User
}

func (u User) SayHello() {
	fmt.Println("Hola me llamo ", u.Name)
}

func (u *User) AddFriend(friend User) {
	u.friends = append(u.friends, friend)
}

func (u User) ListFriends() {
	for i, f := range u.friends {
		fmt.Printf("%d. %s\n", i+1, f.Name)
	}
}

func (u *User) RemoveFriend(id int) {
	index := u.findFriend(id)
	if index < 0 {
		return
	}
	u.friends = append(u.friends[:index], u.friends[index+1:]...)
}

func (u User) findFriend(id int) int {
	for i, f := range u.friends {
		if id == f.ID {
			return i
		}
	}
	return -1
}
