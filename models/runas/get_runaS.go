package runas

import "github.com/Lucasmartinsn/new-api-gin/db"

func Getrunes() (runas []Runasec, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, id_rune, img, nome FROM runas_secundarias`)
	if err != nil {
		return
	}

	for rows.Next() {
		var newrunas Runasec
		err = rows.Scan(&newrunas.Id, &newrunas.Id_rune, &newrunas.Img, &newrunas.Nome)

		if err != nil {
			continue
		}

		runas = append(runas, newrunas)
	}
	return
}
