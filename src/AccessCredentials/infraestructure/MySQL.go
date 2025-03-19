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

    query := "INSERT INTO access_credentials (id_student, email, password) VALUES (?,?,?)"
	
    _, err = mysql.conn.ExecutePreparedQuery(query, accessCredentials.IdStudent, accessCredentials.Email, accessCredentials.Password)

    if err != nil {
        return fmt.Errorf("Error al ejecutar la consulta: %v", err)
    }

    return nil
}

func (mysql *MySQL) GetAccessCredentialsAllStudents() (studentsAC *[]entities.StudentAccessCredentials, err error) {
	var studentsWithAccessCredentials []entities.StudentAccessCredentials
	var student entities.StudentAccessCredentials

	query := "SELECT s.id_student, s.name, s.age ,s.phone_number, ac.email FROM students s iNNER JOIN access_credentials ac ON ac.id_student = s.id_student"

	rows := mysql.conn.FetchRows(query)

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&student.IdStudent, &student.Name, &student.Age, &student.PhoneNumber, &student.Email); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %w", err)
		}

		studentsWithAccessCredentials = append(studentsWithAccessCredentials, student)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error después de iterar sobre las filas: %w", err)
	}

	return &studentsWithAccessCredentials, nil



}