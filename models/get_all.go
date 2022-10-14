package models

import "github.com/Lucasmartinsn/new-api-gin/server/db"

func GetAll() (records []Record) {
	conn := db.Conectar()

	defer conn.Close()

	rows, err := conn.Query(`SELECT id_user, foto,nome_usuario,email,senha FROM records`)
	if err != nil {
		return
	}

	for rows.Next() {
		var newrecods Record

		err = rows.Scan(&newrecods.Id_user, &newrecods.Foto, &newrecods.Nome_ususario, &newrecods.Email, &newrecods.Senha)
		if err != nil {
			continue
		}

		records = append(records, newrecods)
	}
	return
}
