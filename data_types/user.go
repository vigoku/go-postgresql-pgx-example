package data_types

type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

func UserInputSQLQuery() string {
	return `SELECT id, name, email, age FROM users`
}
func UserOutputSQLQuery() string {
	return `INSERT INTO users (id, name, email, age) VALUES ($1, $2, $3, $4)`
}

func (*User) Process() interface{} {
	return ("10", "ABC","abc@bcd.com", "21");
}
