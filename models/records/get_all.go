package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func GetAll() (redords []Record, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT id, foto,nome_usuario,email FROM records`)
	if err != nil {
		return
	}

	for rows.Next() {
		var newRecordes Record

		err = rows.Scan(&newRecordes.Id, &newRecordes.Foto, &newRecordes.Nome_usuario, &newRecordes.Email)
		if err != nil {
			continue
		}

		redords = append(redords, newRecordes)
	}
	return
}
