package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq" //need the init function
)

var db *sql.DB //not exported

func init() {

	var err error
	db, err = sql.Open("postgres", "postgres://temp_user:password@localhost/basket?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("database connected")
}

//ReadItemsDB reads database and returns all the items
func ReadItemsDB(w http.ResponseWriter, req *http.Request, familyName string) ([]map[string]string, error) {

	//query the family
	s, err := getRow(w, familyName)
	var items []map[string]string
	if err != nil {
		return items, err
	}

	//read s if not empty
	if s != "" {
		li := strings.Split(s, "\n")
		for _, v := range li {
			user := strings.Split(v, "\t")[0]
			item := strings.Split(v, "\t")[1]
			items = append(items, map[string]string{
				user: item,
			})
		}
	}
	return items, err
}

//InsertItem appends the item to a familyrow in database
//also takes care of error handling
func InsertItem(w http.ResponseWriter, req *http.Request, familyName string, item string, user string) error {

	//query the row
	s, err := getRow(w, familyName)
	if err != nil {
		return err
	}

	//add new item
	s = s + "\n" + user + "\t" + item

	_, err = db.Exec(`UPDATE customers set items = $1 where name = $2;`, s, familyName)

	if err != nil {
		return err
	}
	fmt.Println("item added to familygroup:" + familyName)
	return err
}

func getRow(w http.ResponseWriter, familyName string) (string, error) {

	//query the family, create if does not exist
	row := db.QueryRow("SELECT items from customers where name = '$1';", familyName)
	var err error
	var s string
	if err := row.Scan(&s); err != nil {

		if err == sql.ErrNoRows { // if row not found
			//create a new row
			err = nil
			_, err = db.Exec("INSERT into customers (name items) VALUES ('$1' '$2');", familyName, "")
			//sanity check
			if err != nil {
				return "", err
			}
			fmt.Println("row created!")
		}
	}
	return s, err
}
