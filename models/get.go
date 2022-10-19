package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Get(id int) (records Record, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id_user, foto,nome_usuario,email,senha FROM records WHERE id_user=$1`, id)

	err = row.Scan(&records.Id_user, &records.Foto, &records.Nome_usuario, &records.Email, &records.Senha)

	return
}
