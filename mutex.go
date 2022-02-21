package sqld

import "database/sql"

// Mutex
type Mutex int64

func (m Mutex) Lock(tx sql.Tx) error {
	rows, err := tx.Query("SELECT pg_advisory_lock($1)", m)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (m Mutex) Unlock(tx sql.Tx) error {
	rows, err := tx.Query("SELECT pg_advisory_unlock($1)", m)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
