package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Upfoto(id int64, records Testerecord) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE records SET foto=$2 WHERE id=$1`,
		id, records.Foto)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
