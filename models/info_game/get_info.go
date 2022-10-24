package infogame

import "github.com/Lucasmartinsn/new-api-gin/db"

func Getinfo(id int) (infogame Infogame, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id, nick_game, days_play, position_game, play_time, comunication FROM recordss WHERE id=$1`, id)

	err = row.Scan(&infogame.Id, &infogame.Nick_game, &infogame.Days_play, &infogame.Position_game, &infogame.Play_time, &infogame.Comunication)

	return
}
