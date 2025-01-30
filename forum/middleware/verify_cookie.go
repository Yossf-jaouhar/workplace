package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	"time"
)

func VerifyCookie(r *http.Request, db *sql.DB) (*http.Cookie, error) {

	token, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	exist := ""
	var expart time.Time

	err = db.QueryRow("SELECT Session, Expared FROM users WHERE Session = ?", token.Value).Scan(&exist, &expart)
	if err != nil {
		return nil, err
	}

	if exist == "" {
		return nil, errors.New("invalid token")
	}

	if time.Now().UTC().After(expart.UTC()) {
		_, err = db.Exec("UPDATE users SET Session = ? WHERE Session = ?", "", exist)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("expiration session")
	}

	return token, nil

}
