package infogame

/// Tabela com os typos e os campos da tabela
type Infogame struct {
	Id            int    `json:"id"`
	Nick_game     string `json:"nick_game"`
	Descricao     string `json:"descricao"`
	Days_play     string `json:"days_play"`
	Position_game string `json:"position_game"`
	Play_time     int    `json:"play_time"`
	Comunication  string `json:"comunication"`
}
