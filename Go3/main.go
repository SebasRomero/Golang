package main

import (
	"fmt"

	"github.com/example/users"
)

func main() {
	pedro := users.User{ID: 2, Name: "Pedro", Age: 33}
	juan := users.User{ID: 3, Name: "Juan", Age: 21}
	martha := users.User{ID: 1, Name: "Martha", Age: 28}
	camila := users.User{ID: 4, Name: "Camila", Age: 12}
	martha.SayHello()

	martha.AddFriend(pedro)
	martha.AddFriend(juan)
	martha.AddFriend(camila)

	martha.ListFriends()

	martha.RemoveFriend(camila.ID)
	fmt.Println("--------")
	martha.ListFriends()

}
