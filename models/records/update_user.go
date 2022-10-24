package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Update(id int64, records Record) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE records SET foto=$2, nome_usuario=$3, email=$4 WHERE id=$1`, id, records.Foto, records.Nome_usuario, records.Email)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
