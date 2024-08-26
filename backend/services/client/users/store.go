package users

import (
	"backend/services/client"
	"database/sql"
	"fmt"
)

func UserCreate(db *sql.DB, u *client.UserPayload) error {
	err := db.QueryRow("SELECT DISTINCT Email FROM Users WHERE Email = ?", u.Email)
	if err != nil {
		fmt.Errorf("Email já cadastrado!", err)
	}

	_, errinho := db.Exec("INSERT INTO Users (Name, Email, Password) VALUES (?,?,?)", u.Name, u.Email, u.Password)
	if errinho != nil {
		return errinho
	}

	return nil
}
