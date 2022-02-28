package syncd

import "database/sql"

// PGMutexKey is an int64 key based distributed lock that provides mutual exclusion across 
// processes and servers. It relies on PostgreSQL's system level locking functions for 
// lock acquisition and release.
type PGMutexKey int64

// Lock implements the Locker interface by using session level locking on your PostgreSQL instance.
func (m PGMutexKey) Lock(ctx context.Context, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT pg_advisory_lock($1)", m)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// Unlock implements the Unlocker interface by releasing session level locks on your PostgreSQL instance.
func (m PGMutexKey) Unlock(ctx context.Context, tx *sql.Tx) error {
	rows, err := tx.Query("SELECT pg_advisory_unlock($1)", m)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
