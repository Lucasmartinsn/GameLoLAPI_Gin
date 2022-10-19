package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func GetAll() (redords []Record, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT id_user, foto,nome_usuario,email,senha FROM records`)
	if err != nil {
		return
	}

	for rows.Next() {
		var newRecordes Record

		err = rows.Scan(&newRecordes.Id_user, &newRecordes.Foto, &newRecordes.Nome_usuario, &newRecordes.Email, &newRecordes.Senha)
		if err != nil {
			continue
		}

		redords = append(redords, newRecordes)
	}
	return
}
