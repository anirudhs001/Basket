package models

import "database/sql"

func updateCustomerRow(s string, name string) error {

	_, err := db.Exec(`
		UPDATE customers 
		set items = $1 
		where name = $2;`,
		s,
		name)

	return err

}

func insertSellerRow(s Seller) error {

	_, err := db.Exec(`
		INSERT into sellers
		(name, address, opentime, closetime)
		VALUES ($1, $2, $3, $4);`,
		s.Name,
		s.Addr,
		s.OpenTime,
		s.CloseTime)

	return err
}

//reads DB and sends items in row if found
// if row not found, creates one too
// if an error occurs, returns an empty string and the err
func getCustomerRow(familyName string) (string, error) {

	//query the family, create if does not exist
	row := db.QueryRow("SELECT items from customers where name = $1;", familyName)
	var err error
	var s string
	if err := row.Scan(&s); err != nil {

		if err == sql.ErrNoRows { // if row not found
			//create a new row
			err = nil
			_, err = db.Exec("INSERT into customers (name,items) VALUES ($1,$2);", familyName, "")
			//sanity check
			if err != nil {
				return "", err
			}
		}
	}
	return s, err
}

//getSellerRows returns the matching rows and an error
//if no arguments passed: returns all the sellers
//if a single argument passed (shop name), returns the corresponding row
func getSellerRows(args ...string) (*sql.Rows, error) {

	var rows *sql.Rows
	var err error

	if len(args) == 0 { //no arguments
		rows, err = db.Query("Select * from sellers;")
	}
	if len(args) == 1 { //seller name sent
		rows, err = db.Query("select * from sellers where name=$1;", args[0])
	}
	return rows, err
}
