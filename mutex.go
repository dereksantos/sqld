package sqld

import "database/sql"

// Mutex
type Mutex int32

func (m Mutex) Lock(tx sql.Tx) error {
	return nil
}

func (m Mutex) Unlock(tx sql.Tx) error {
	return nil
}
