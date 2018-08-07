package UserPlayer

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Grettings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func newPlayer(id int, name, location string, gameId int) *Player {
	return &Player{
		User:   &User{id, name, location},
		GameId: gameId,
	}
}

func Run() {
	player := newPlayer(1, "Mason", "Chicago", 1570584)
	fmt.Println(player.Grettings())
}
