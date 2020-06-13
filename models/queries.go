package models

func updateRow(s string, name string) error {

	_, err := db.Exec(`
	UPDATE customers 
	set items = $1 
	where name = $2;`, s, name)

	return err

}
