package infraestructure

var (
	mysql *MySQL
)

func InitMySQL() {
	mysql = NewMySQL()
}

func GetMySQL() *MySQL {
	return mysql
}