package company

import "database/sql"

type postgresRepo struct {
	db *sql.DB
}

//NewMySQLRepository create new repository
func NewPostgresRepository(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}