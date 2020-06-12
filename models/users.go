package models

//User struct
type User struct {
	ParentGroup string
	Name        string
}

type group struct {
	id    string
	users []User
}

func (g group) addUser(u User) {
	g.users = append(g.users, u)
}
