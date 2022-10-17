package models

type Record struct {
	Id_user       int    `json:"id_user"`
	Foto          string `json:"foto"`
	Nome_ususario string `json:"nome_usuario"`
	Email         string `json:"email"`
	Senha         int    `json:"senha"`
}
