package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func Get(id int) (records Record, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id, foto,nome_usuario,email,nick_game, days_play, position_game, play_time, comunication FROM recordss WHERE id=$1`, id)

	err = row.Scan(&records.Id, &records.Foto, &records.Nome_usuario, &records.Email, &records.Nick_game, &records.Days_play, &records.Position_game,
		&records.Play_time, &records.Comunication)

	return
}
