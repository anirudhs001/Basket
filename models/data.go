package models

//User struct
type User struct {
	ParentGroup string
	Name        string
}

//Item struct
type Item struct {
	ID, Name, Item string
}

//Seller struct
type Seller struct {
	ID, Name, Addr, OpenTime, CloseTime string
}

type group struct {
	id    string
	users []User
}

func (g group) addUser(u User) {
	g.users = append(g.users, u)
}
