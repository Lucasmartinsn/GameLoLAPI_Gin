package models

import "github.com/Lucasmartinsn/new-api-gin/db"

func GetAll() (redords []Recordg, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT id, foto,nome_usuario,email,nick_game, days_play, position_game, play_time, comunication FROM recordss`)
	if err != nil {
		return
	}

	for rows.Next() {
		var newRecordes Recordg

		err = rows.Scan(&newRecordes.Id, &newRecordes.Foto, &newRecordes.Nome_usuario, &newRecordes.Email, &newRecordes.Nick_game, &newRecordes.Days_play, &newRecordes.Position_game,
			&newRecordes.Play_time, &newRecordes.Comunication)
		if err != nil {
			continue
		}

		redords = append(redords, newRecordes)
	}
	return
}
