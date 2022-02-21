package syncd

// Locker interface abstracts the implementation of lock acquisition mechanisms.
// Implementors shoud handle context cancelling and use tx to rely on undlerying database
// technologies for lock acqusition.
type Locker interface {
	Lock(context.Context, tx sql.Tx) error
}

// LockerFunc provides an option for building functional implementations of distributed locks.
type LockerFunc func(context.Context, tx sql.Tx) error

// Lock implements Locker interface by delegating to the receiver func.
func (lf LockerFunc) Lock(context.Context, tx sql.Tx) error {
	return lf(ctx, tx)
}

// Unlocker interface abstracts the implementation of lock release mechanisms.
// Implementors shoud handle context cancelling and use tx to rely on undlerying database
// technologies for lock releasing.
type Unlocker interface {
	Unlock(tx sql.Tx) error
}

// UnlockerFunc provides an option for building functional implementations of distributed locks.
type UnlockerFunc func(context.Context, tx sql.Tx) error

// Lock implements Locker interface by delegating to the receiver func.
func (uf UnlockerFunc) Unlock(context.Context, tx sql.Tx) error {
	return uf(ctx, tx)
}

// LockUnlocker composes acquisition and release behaviours into a single interface for
// more convenient handling of lock and unlock implementation.
type LockUnlocker interface {
	Locker
	Unlocker
}