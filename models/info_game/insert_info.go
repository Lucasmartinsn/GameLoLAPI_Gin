package infogame

import "github.com/Lucasmartinsn/new-api-gin/db"

func Insertinfo(infogame Infogame) (id int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	sql := `INSERT INTO user_game_infor (nick_game, descricao, days_play, position_game, play_time, comunication) 	VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	err = conn.QueryRow(sql, infogame.Nick_game, infogame.Descricao, infogame.Days_play, infogame.Position_game, infogame.Play_time, infogame.Comunication).Scan(&id)

	defer conn.Close()

	return
}
