package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Get(id int) (records Record, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id, foto,nome_usuario,email FROM records WHERE id=$1`, id)

	err = row.Scan(&records.Id, &records.Foto, &records.Nome_usuario, &records.Email)

	return
}
