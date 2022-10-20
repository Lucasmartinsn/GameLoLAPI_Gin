package infogame

import "github.com/Lucasmartinsn/new-api-gin/db"

func Updateinfo(id int64, infogame Infogame) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE user_game_infor SET nick_game=$2, days_play=$3, position_game=$4 play_time=$5 comunication=$6 WHERE id=$1`,
		id, infogame.Nick_game, infogame.Days_play, infogame.Position_game, infogame.Play_time, infogame.Comunication)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
