package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec("DELETE FROM records WHERE id=$1", id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
