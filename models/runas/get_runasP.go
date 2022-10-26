package runas

import "github.com/Lucasmartinsn/new-api-gin/db"

func GetrRuna() (runes []Runaprinc, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, img, nome FROM runas_principais`)
	if err != nil {
		return
	}

	for rows.Next() {
		var newrunes Runaprinc

		err = rows.Scan(&newrunes.Id, &newrunes.Img, &newrunes.Nome)
		if err != nil {
			continue
		}
		runes = append(runes, newrunes)
	}
	return
}
