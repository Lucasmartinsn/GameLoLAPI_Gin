package runas

type Runaprinc struct {
	Id   int    `json:"id"`
	Img  string `json:"img"`
	Nome string `json:"nome"`
}

type Runasec struct {
	Id      int    `json:"id"`
	Id_rune int    `json:"id_rune"`
	Img     string `json:"img"`
	Nome    string `json:"nome"`
}

type UpfotorunaS struct {
	Img string `json:"img"`
}
