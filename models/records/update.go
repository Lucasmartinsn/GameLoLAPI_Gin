package models

import (
	"fmt"

	"github.com/Lucasmartinsn/new-api-gin/db"
)

func (s StructUpdateInformacoes) UpdateInformacoes(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	comandoSQL := fmt.Sprintf(`UPDATE recordss SET foto=%s, nome_usuario=%s email=%s senha=%d nick_game=%s descricao=%s days_play=%s position_game=%s play_time=%d comunication=%s where id=%d`, s.Foto,	s.Nome_usuario, s.Email, s.Senha, s.Nick_game, s.Descricao, s.Days_play, s.Position_game, s.Play_time, s.Comunication, id)

	fmt.Println(comandoSQL)

	response, err := conn.Exec(comandoSQL)
	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
