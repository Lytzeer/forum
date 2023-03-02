package forum

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

func EditMail(id int, mail string, mailconf string) {
	if mail == mailconf {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_edit_mail, err := db.Prepare("UPDATE User SET email=? WHERE id=?")
		CheckErr(err)
		request_edit_mail.Exec(mail, id)
		db.Close()
	}
}

func EditPassword(id int, currentpassword string, password string, passwordconf string) {
	if password == passwordconf {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_current_password, err := db.Prepare("SELECT password FROM User WHERE id=?")
		CheckErr(err)
		var currentpasshash []byte
		err = request_current_password.QueryRow(id).Scan(&currentpasshash)
		CheckErr(err)
		err = bcrypt.CompareHashAndPassword(currentpasshash, []byte(currentpassword))
		CheckErr(err)
		request_edit_password, err := db.Prepare("UPDATE User SET password=? WHERE id=?")
		CheckErr(err)
		passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		CheckErr(err)
		request_edit_password.Exec(passhash, id)
		db.Close()
	}
}

func EditUsername(id int, username string) {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_edit_username, err := db.Prepare("UPDATE User SET username=? WHERE id=?")
	CheckErr(err)
	request_edit_username.Exec(username, id)
	db.Close()
}

func DeleteAccount(id int) {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_delete_account, err := db.Prepare("DELETE From User WHERE id= ?")
	CheckErr(err)
	request_delete_account.Exec(id)
	db.Close()
}
