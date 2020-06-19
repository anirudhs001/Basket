package models

//User struct
type User struct {
	ParentGroup string
	Name        string
}

//Item struct
//ID: each new item inserted has a random(unique) string id
//Name: name of user who put in the item
//Item: name of the item inserted (eg lays)
type Item struct {
	ID, Name, Item string
}

//Order struct
type Order struct {
	FamilyName, Shop, Date string
}

//ShoppingList struct
type ShoppingList struct {
	Date       string
	FamilyName string
	Items      []string
}

//Seller struct
type Seller struct {
	ID, Name, Addr, OpenTime, CloseTime string
}
