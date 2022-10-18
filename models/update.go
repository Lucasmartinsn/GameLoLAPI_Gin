package models

import "github.com/Lucasmartinsn/new-api-gin/server/db"

func Update(id int64, records Record) (int64, error) {
	conn, err := db.Conectar()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE records SET foto=$2, nome_usuario=$3, email=$4, senha=$5 WHERE id_user=$1`,
		id, records.Foto, records.Nome_ususario, records.Email, records.Senha)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
