package models

import (
	"github.com/Lucasmartinsn/new-api-gin/db"
)

func Insert(records Record) (id int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	sql := `INSERT INTO records (foto, nome_usuario, email, senha) VALUES ($1,$2,$3,$4) RETURNING id_user`

	err = conn.QueryRow(sql, records.Foto, records.Nome_usuario, records.Email, records.Senha).Scan(&id)

	defer conn.Close()

	return
}
