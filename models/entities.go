package models

/// Tabela com os typos e os campos da tabela
type Record struct {
	Id_user      int    `json:"id"`
	Foto         string `json:"foto"`
	Nome_usuario string `json:"nome_usuario"`
	Email        string `json:"Email"`
	Senha        int    `json:"Senha"`
}
