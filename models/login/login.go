package login

import "github.com/Lucasmartinsn/new-api-gin/db"

func Verificalogin(senha int) (login Login, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := conn.QueryRow(`SELECT senha FROM recordss WHERE senha=$1'`, senha)

	err = sql.Scan(&login.Senha)

	return
}
