package models

/// Tabela com os typos e os campos da tabela
type Record struct {
	Id           int    `json:"id"`
	Foto         string `json:"foto"`
	Nome_usuario string `json:"nome_usuario"`
	Email        string `json:"Email"`
}

type Testerecord struct {
	Foto string `json:"foto"`
}
