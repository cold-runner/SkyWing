package store

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	Users() UserStore
	Roles() RoleStore
	Policies() PolicyStore
	Close() error
}
