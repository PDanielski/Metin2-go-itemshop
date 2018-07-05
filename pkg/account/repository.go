package account

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

//Repository is used for retrieving accounts from any database
type Repository interface {
	ByID(id ID) (*Account, bool)
	ByLogin(login string) (*Account, bool)
}

//SQLRepository is the most common implementation
type SQLRepository struct {
	db *sql.DB
}

//NewSQLRepository creates a new instance of SQLRepository
func NewSQLRepository(db *sql.DB) Repository {
	return SQLRepository{db: db}
}

//ByID is used for searching an account by its id
func (repo SQLRepository) ByID(id ID) (*Account, bool) {
	row := repo.db.QueryRow("SELECT login, password, social_id, email, create_time, gold, warpoints, biscuits FROM account.account WHERE id = ?", id)
	var (
		login        string
		password     string
		socialID     string
		email        string
		creationTime mysql.NullTime
		gold         int
		warpoints    int
		biscuits     int
	)
	err := row.Scan(&login, &password, &socialID, &email, &creationTime, &gold, &warpoints, &biscuits)
	if err != nil {

		return nil, false
	}
	return New(id, login, NewPassword(password, true), socialID, email, gold, warpoints, biscuits), true
}

//ByLogin is used for searching an account by its login
func (repo SQLRepository) ByLogin(login string) (*Account, bool) {
	row := repo.db.QueryRow("SELECT id, password, social_id, email, create_time, gold, warpoints, biscuits FROM account.account WHERE login = ?", login)
	var (
		id           ID
		password     string
		socialID     string
		email        string
		creationTime mysql.NullTime
		gold         int
		warpoints    int
		biscuits     int
	)
	err := row.Scan(&id, &password, &socialID, &email, &creationTime, &gold, &warpoints, &biscuits)
	if err != nil {
		return nil, false
	}
	return New(id, login, NewPassword(password, true), socialID, email, gold, warpoints, biscuits), true
}
