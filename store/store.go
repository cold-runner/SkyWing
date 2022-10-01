package store

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	Users() UserStore
	Roles() RoleStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}
