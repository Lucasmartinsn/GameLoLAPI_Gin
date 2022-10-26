package models

/// Tabela com os typos e os campos da tabela
/*
type Record struct {
	Id           int    `json:"id"`
	Foto         string `json:"foto"`
	Nome_usuario string `json:"nome_usuario"`
	Email        string `json:"Email"`
}
*/
type Putfoto struct {
	Foto string `json:"foto"`
}
type Putnome struct {
	Nome_usuario string `json:"nome_usuario"`
}

type Putsenha struct {
	Senha int `json:"senha"`
}

type Recordg struct {
	Id            int    `json:"id"`
	Foto          string `json:"foto"`
	Nome_usuario  string `json:"nome_usuario"`
	Email         string `json:"Email"`
	Nick_game     string `json:"nick_game"`
	Descricao     string `json:"descricao"`
	Days_play     string `json:"days_play"`
	Position_game string `json:"position_game"`
	Play_time     int    `json:"play_time"`
	Comunication  string `json:"comunication"`
}

type Record struct {
	Id            int    `json:"id"`
	Foto          string `json:"foto"`
	Nome_usuario  string `json:"nome_usuario"`
	Email         string `json:"Email"`
	Senha         int    `json:"senha"`
	Nick_game     string `json:"nick_game"`
	Descricao     string `json:"descricao"`
	Days_play     string `json:"days_play"`
	Position_game string `json:"position_game"`
	Play_time     int    `json:"play_time"`
	Comunication  string `json:"comunication"`
}
