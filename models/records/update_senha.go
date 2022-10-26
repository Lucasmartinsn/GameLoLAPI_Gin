package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Upsenha(id int64, records Putsenha) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE recordss SET senha=$2 WHERE id=$1`,
		id, records.Senha)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
