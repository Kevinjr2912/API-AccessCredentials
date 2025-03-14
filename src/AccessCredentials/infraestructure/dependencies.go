package infraestructure

import "api_accesscredentials/src/AccessCredentials/infraestructure/adapters"

var (
	mysql  *MySQL
	rabbit *adapters.Rabbit
)

func InitDependencies() {
	mysql = NewMySQL()
	rabbit = adapters.NewRabbitMq()
}

func GetMySQL() *MySQL {
	return mysql
}

func GetRabbit() *adapters.Rabbit {
	return rabbit
}
