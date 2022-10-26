package models

import (
	"github.com/Lucasmartinsn/new-api-gin/db"
)

func Insert(records Record) (id int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	sql := `INSERT INTO recordss (foto, nome_usuario, email, senha, nick_game, descricao, days_play, position_game, play_time, comunication) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err = conn.QueryRow(sql, records.Foto, records.Nome_usuario, records.Email, records.Senha, records.Nick_game, records.Descricao, records.Days_play, records.Position_game,
		records.Play_time, records.Comunication).Scan(&id)

	defer conn.Close()

	return
}
