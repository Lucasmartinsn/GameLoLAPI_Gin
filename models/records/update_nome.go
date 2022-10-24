package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Upname(id int64, records Putnome) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE recordss SET nome_usuario=$2 WHERE id=$1`,
		id, records.Nome_usuario)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
