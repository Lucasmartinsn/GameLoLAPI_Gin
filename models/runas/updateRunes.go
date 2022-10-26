package runas

import "github.com/Lucasmartinsn/new-api-gin/db"

func UpdateRuna(id int64, runas UpfotorunaS) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	response, err := conn.Exec(`UPDATE runas_secundarias SET img=$2 WHERE id_rune=$1`, id, runas.Img)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
