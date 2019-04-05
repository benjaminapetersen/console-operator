package errors

// a sync progressing error captures the idea that the operand is in an incomplete state
// and needs to be reconciled.  It should be used when the sync loop should abort, but
// it implies the idea of "progressing", not "failure".  A built-in error type can be
// used and returned when a true failure is encountered.
type SyncProgressingError struct {
	// The summary error message
	message string
}

// implementing the error interface
func (e *SyncProgressingError) Error() string {
	return e.message
	// return fmt.Sprintf("%s from %s (%s)", e.message, e.resource.name, e.resource.groupVersionKind)
}

// NewSyncError("Sync failed on xyz")
func NewSyncError(msg string) *SyncProgressingError {
	err := &SyncProgressingError{
		message: msg,
	}
	return err
}

func IsSyncError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*SyncProgressingError)
	return ok
}
