package infraestructure

import (
	"api_accesscredentials/src/AccessCredentials/domain/entities"
	"api_accesscredentials/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	fmt.Printf("Conexión a la base de datos")

	return &MySQL{conn: conn}
}


func (mysql *MySQL) AssociateAccessCredentialsToStudent(accessCredentials *entities.AccessCredentials) (err error) {

    query := "INSERT INTO access_credentials (id_student, user, password) VALUES (?,?,?)"
	
    _, err = mysql.conn.ExecutePreparedQuery(query, accessCredentials.IdStudent, accessCredentials.User, accessCredentials.Password)

    if err != nil {
        return fmt.Errorf("Error al ejecutar la consulta: %v", err)
    }

    return nil
}