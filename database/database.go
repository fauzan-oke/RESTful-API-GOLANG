package database

var conn string

func init() {
	conn = "mysql"
}

func GetDatabase() string {
	return conn
}
