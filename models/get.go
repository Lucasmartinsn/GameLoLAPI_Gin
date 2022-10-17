package models

import "github.com/Lucasmartinsn/new-api-gin/server/db"

func GetOne(id int) (redords Record, err error) {
	conn, err := db.Conectar()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id_user, foto,nome_usuario,email,senha FROM records where id_user=$1`, id)

	err = row.Scan(&redords.Id_user, &redords.Foto, &redords.Nome_ususario, &redords.Email, &redords.Senha)

	return
}
